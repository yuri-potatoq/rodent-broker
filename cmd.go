package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	ServerPort int    = 5000
	ServerHost string = "0.0.0.0"
	ServerAddr string
	KafkaEnv   string
)

var serverCmd = &cobra.Command{
	Use:   "Rodent Broker",
	Short: "Command an broker with web requests",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LOGO DO APP")
	},
}

var serverRun = &cobra.Command{
	Use:   "run [OPTIONS]",
	Short: "Init web server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Connection server at %s...\n", ServerAddr)
		InitServer()
	},
}

func init() {
	serverCmd.AddCommand(serverRun)

	serverRun.PersistentFlags().StringVar(
		&ServerAddr, "addr", fmt.Sprintf("%s:%d", ServerHost, ServerPort),
		`full hosting address`)

	serverRun.PersistentFlags().StringVar(
		&KafkaEnv, "kafka-env", "local",
		`kafka config environment. ['local', 'dev', 'prod']`)
}

func Execute() {
	if err := serverCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
