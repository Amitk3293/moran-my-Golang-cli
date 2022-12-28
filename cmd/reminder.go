/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Println("reminder called")

		// 1. we need a runner to check the time,
		// 2. create Slack integration.
		// 3. 

		// setReminder, _ := cmd.Flags().GetString("set")
		// setTime, _ := cmd.Flags().GetString("time")

		// // now := time.Now()
		// // duration := time.Since(now)
		// // select {
		// // 	case
		// // }

	},
}

func init() {
	rootCmd.AddCommand(reminderCmd)
	// Here you will define your flags and configuration settings.
	reminderCmd.Flags().StringP("set", "s", "", "set a reminder")
	reminderCmd.Flags().StringP("time", "t", "", "set time to the reminder-Minutes for now")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reminderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reminderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
