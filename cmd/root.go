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
	cproto "ftserver/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var srv *grpc.Server

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ftserver",
	Short: "CLI app for short test",
	Long:  `CLI app for long long test.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//type AddServiceServer struct{}
type QuestionsServiceServer struct{}

func initGRPC() {
	fmt.Println("initGRPC called")

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	//srv := grpc.NewServer()
	srv = grpc.NewServer()
	cproto.RegisterQuestionsServiceServer(srv, &QuestionsServiceServer{})
	//cproto.RegisterAddServiceServer(srv, &AddServiceServer{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
