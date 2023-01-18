/*
Copyright © 2022 Amit Karni amitkarni3293@gmail.com
*/
package cmd

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
		storeReminder()
	},
}

var slackMessage string
var reminderText string

func init() {
	rootCmd.AddCommand(reminderCmd)
	// Here you will define your flags and configuration settings.
	reminderCmd.Flags().StringVarP(&reminderText, "set", "s", "", "set a reminder")
	// reminderCmd.Flags().StringP("time", "t", "", "set time to the reminder-Minutes for now")
	reminderCmd.Flags().StringVarP(&slackMessage, "message", "m", "", "The message to send to Slack")
	// and all subcommands, e.g.:
	// reminderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reminderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type config struct {
	SLACK_WEBHOOK_URL string `yaml:"SLACK_WEBHOOK_URL"`
}

// Read the YAML file and parse its contents
func (c *config) readConfig() {
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %v", err)
		return
	}
	if err := yaml.Unmarshal(file, &c); err != nil {
		fmt.Printf("Error parsing config file: %v", err)
		return
	}
}

func runSlackCommand() {

	c := &config{}
	c.readConfig()

	// Use the value from the YAML file
	client := &http.Client{}
	body := fmt.Sprintf(`{"text": "%s"}`, slackMessage)
	req, err := http.NewRequest("POST", c.SLACK_WEBHOOK_URL, strings.NewReader(body))
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

func storeReminder() error {
	// Open a connection to the SQLite database file
	db, err := sql.Open("sqlite3", "reminders.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create the table to store the reminder notes
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS reminders (
                id INTEGER PRIMARY KEY,
                text TEXT NOT NULL,
				datetime TIMESTAMP NOT NULL
        )
    `)
	if err != nil {
		fmt.Printf("There was an error in creating DB table: %v", err)
		return err
	}

	// Get the current time
	now := time.Now().Format("2006-01-02 15:04:05")
	// now := time.Now().Unix()
	fmt.Println(now)

	// Insert the reminder note w text and current timestamp into the table
	rowString := fmt.Sprintf(`INSERT INTO reminders (text, datetime) VALUES (%q, %q)`, reminderText, now)
	fmt.Println(rowString)
	_, err = db.Exec(rowString)
	if err != nil {
		fmt.Printf("There was an error in insert values into DB table: \n %v", err)
		return err
	}
	return nil
}
