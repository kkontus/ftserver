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
	"fmt"
	"github.com/spf13/cobra"
	"time"
	"encoding/json"
	"net/http"
	"context"
	cproto "ftserver/proto"
)

const (
	URL = "url" // don't change this, it's just so that we don't have string flags on multiple places
)

const (
	NETWORK = iota
	FILE
	DB
)

var url string
var questionsAndAnswers []Questions
var rankings map[string]UserRanking

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Gets list of all questions",
	Long:  `Gets the list of all available questions.`,
	Run: func(cmd *cobra.Command, args []string) {
		url = cmd.Flag(URL).Value.String()

		rankings = make(map[string]UserRanking)

		initGRPC()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)

	defaultUrl := "https://gist.githubusercontent.com/kkontus/323b05ed729e53c7dd5307bf6231693a/raw/2ca073e5dbfd10a7ded4883a565584db71aff85c/questions"
	url = defaultUrl // this will be overwritten if -u flag is used
	questionsCmd.PersistentFlags().StringP(URL, "u", defaultUrl, "Url for the quiz questions")
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

type UserRanking struct {
	Name           string
	Percentage     string
	CorrectAnswers float64
}

type Questions struct {
	Question string   `json:"question"`
	Correct  string   `json:"correct"`
	Answers  []string `json:"answers"`
}

func createClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
}

func (s *QuestionsServiceServer) GetAllQuestions(ctx context.Context, request *cproto.LoadQuestionsList) (*cproto.ReturnQuestionsList, error) {
	result := []*cproto.Questions{}

	// in case we need different options for different things
	if request.Network == NETWORK {
		client := createClient()
		resp, err := client.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		questions := []Questions{}
		err = json.NewDecoder(resp.Body).Decode(&questions)
		if err != nil {
			panic(err)
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
	} else if request.File == FILE {

	} else if request.Db == DB {

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
