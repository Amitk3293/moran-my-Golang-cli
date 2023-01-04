/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// reminderCmd represents the reminder command
var reminderCmd = &cobra.Command{
	Use:   "reminder",
	Short: "reminder is a sub command when you can use to set up and manage reminders",
	Long: `A longer description that spans multiple lines and likely contains examples
		   and usage of using your command. For example:
		   Cobra is a CLI library for Go that empowers applications.
		   This application is a tool to generate the needed files
		   to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runSlackCommand()
	},
}
var slackMessage string

func init() {
	rootCmd.AddCommand(reminderCmd)
	// Here you will define your flags and configuration settings.
	reminderCmd.Flags().StringP("set", "s", "", "set a reminder")
	reminderCmd.Flags().StringP("time", "t", "", "set time to the reminder-Minutes for now")
	reminderCmd.Flags().StringVarP(&slackMessage, "message", "m", "", "The message to send to Slack")
	// and all subcommands, e.g.:
	// reminderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reminderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runSlackCommand() {
	client := &http.Client{}
	body := fmt.Sprintf(`{"text": "%s"}`, slackMessage)
	req, err := http.NewRequest("POST", "https://hooks.slack.com/services/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", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("There was en error: %v", err)
		return
	}
	fmt.Println("Slack message sent successfully!")
}
