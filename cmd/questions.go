// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	cproto "ftserver/proto"
	cutil "ftserver/util"
	cdb "ftserver/databases"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	. "ftserver/entity" // using dot is not a best practice, I only use it on entity/struct that are reused for cleaner code,
	// but in general everything else should use qualifiers
)

// don't change this, it's just so that we don't have string flags on multiple places
const (
	URL      = "url"
	FILEPATH = "filepath"
	MODE     = "mode"
)

const (
	NETWORK = "network"
	FILE    = "file"
	DB      = "db"
)

var url string
var filePath string
var mode string
var questionsAndAnswers []Questions
var rankings map[string]UserRanking

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:       "questions",
	Short:     "Gets list of all questions",
	Long:      `Gets the list of all available questions from different repositories (network, file, db).`,
	ValidArgs: []string{MODE, URL, FILEPATH},
	Run: func(cmd *cobra.Command, args []string) {
		mode = cmd.Flag(MODE).Value.String()
		url = cmd.Flag(URL).Value.String()
		filePath = cmd.Flag(FILEPATH).Value.String()

		optionsForMode := []string{NETWORK, FILE, DB}
		if !cutil.Contains(mode, optionsForMode) {
			fmt.Println("Mode flag only allows following flags: network, file, db")
			os.Exit(1)
		}

		rankings = make(map[string]UserRanking)

		initGRPC()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)

	defaultUrl := "https://gist.githubusercontent.com/kkontus/323b05ed729e53c7dd5307bf6231693a/raw/2ca073e5dbfd10a7ded4883a565584db71aff85c/questions"
	defaultFile := "questions.json"
	url = defaultUrl       // this will be overwritten if -u flag is used
	filePath = defaultFile // this will be overwritten if -f flag is used
	questionsCmd.Flags().StringP(URL, "u", defaultUrl, "Url for the network quiz questions")
	questionsCmd.Flags().StringP(FILEPATH, "f", defaultFile, "File for the local quiz questions")
	questionsCmd.Flags().StringP(MODE, "m", NETWORK, "Mode for the quiz questions (network, file, db)")
	questionsCmd.MarkFlagRequired(MODE)
}

//type QuestionsServiceServer struct{}

//func initGRPC() {
//	fmt.Println("initGRPC called")
//
//	listener, err := net.Listen("tcp", ":4040")
//	if err != nil {
//		panic(err)
//	}
//
//	srv := grpc.NewServer()
//	cproto.RegisterQuestionsServiceServer(srv, &QuestionsServiceServer{})
//	reflection.Register(srv)
//
//	if err = srv.Serve(listener); err != nil {
//		panic(err)
//	}
//}

func createClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
}

func (s *QuestionsServiceServer) GetAllQuestions(ctx context.Context, request *cproto.LoadQuestionsList) (*cproto.ReturnQuestionsList, error) {
	result := []*cproto.Questions{}
	questions := []Questions{}

	// in case we need different options for different things
	if mode == NETWORK {
		client := createClient()
		resp, err := client.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&questions)
		if err != nil {
			panic(err)
		}
	} else if mode == FILE {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()

		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = json.Unmarshal(byteValue, &questions)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if mode == DB {
		db, err := cdb.Connect(cdb.MYSQL_HOST, cdb.MYSQL_PORT, cdb.MYSQL_USERNAME, cdb.MYSQL_PASSWORD, cdb.MYSQL_DBNAME)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		questions = cdb.FindAll(db, err)
	}

	questionsAndAnswers = questions // we will set that in a var so that we can use it to check the user answers
	for _, elem := range questions {
		q := cproto.Questions{
			Question: elem.Question,
			Correct:  elem.Correct,
			Answers:  elem.Answers,
		}
		result = append(result, &q)
	}

	return &cproto.ReturnQuestionsList{Result: result}, nil
}

func (s *QuestionsServiceServer) CheckUserAnswers(ctx context.Context, request *cproto.SendUserAnswers) (*cproto.UserResults, error) {
	var correctAnswers []string
	var userAnswers []string
	var userScore string
	if len(request.Answers) == len(questionsAndAnswers) {
		fmt.Println("Checking user answers...")

		correct := 0.0
		for i, elem := range questionsAndAnswers {
			correctAnswers = append(correctAnswers, elem.Correct)
			if elem.Correct == request.Answers[i] {
				correct++
				userAnswers = append(userAnswers, "correct")
			} else {
				userAnswers = append(userAnswers, "wrong")
			}
		}

		res := (correct / float64(len(questionsAndAnswers))) * 100
		userScore = fmt.Sprintf("%.2f%s", res, "%")

		fmt.Printf("SCORE: %s\n", userScore)

		ur := UserRanking{
			Name:           request.User,
			Percentage:     userScore,
			CorrectAnswers: correct,
		}

		rankings[request.User] = ur

	} else {
		fmt.Println("Number or answers doesn't match number of questions")
	}

	return &cproto.UserResults{Result: userAnswers, Answers: correctAnswers, Percentage: userScore}, nil
}

func (s *QuestionsServiceServer) CheckUserRanking(ctx context.Context, request *cproto.LoadUserRanking) (*cproto.ReturnUserRanking, error) {
	percentageForUser := "No players set"
	scoreOverall := "No ranking set"

	numOfPlayers := len(rankings)
	if numOfPlayers > 0 {

		if searchedUser, ok := rankings[request.User]; ok {
			percentageForUser = searchedUser.Percentage
			correctAnswersForUser := searchedUser.CorrectAnswers
			usersLessThanMe := 0

			for _, elem := range rankings {
				if elem.CorrectAnswers < correctAnswersForUser {
					usersLessThanMe++
				}
			}

			x := float64(usersLessThanMe) / float64(numOfPlayers) * 100
			if numOfPlayers == 1 {
				scoreOverall = fmt.Sprintf("You are the only player so far, your scored %s on the quiz!!!", percentageForUser)
			} else {
				if x > 75.0 {
					scoreOverall = fmt.Sprintf("Great job, you did better than %.2f%s of quizers!!!", x, "%")
				} else if x > 50.0 {
					scoreOverall = fmt.Sprintf("Solid job, you did better than %.2f%s of quizers!!!", x, "%")
				} else {
					scoreOverall = fmt.Sprintf("You can do it better, you did better than %.2f%s of quizers!!!", x, "%")
				}
			}
		} else {
			percentageForUser = "We couldn't find your rank, did you take the quiz?"
			scoreOverall = "You weren't found on ranking board"
		}
	}

	return &cproto.ReturnUserRanking{Score: percentageForUser, ScoreOverall: scoreOverall}, nil
}
