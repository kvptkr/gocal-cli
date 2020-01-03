/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/kvptkr/gocal-cli/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add events to your calendar",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		calendarClient := auth.GetClient()
		startTime, err := time.Parse(time.RFC3339, "2020-01-15T13:00:00.00Z")
		if err != nil {
			fmt.Println("There was a problem parsing the start time")
		}

		endTime, err := time.Parse(time.RFC3339, "2020-01-15T14:00:00.00Z")
		if err != nil {
			fmt.Println("There was a problem parsing the end time")
		}

		start := calendar.EventDateTime{
			DateTime: startTime.String(),
		}
		end := calendar.EventDateTime{
			DateTime: endTime.String(),
		}
		event := calendar.Event{
			Summary:     "test",
			Description: "test",
			Start:       &start,
			End:         &end,
		}
		_, err = calendarClient.Events.Insert("primary", &event).Do()
		if err == nil {
			fmt.Println("shit worked")
		} else {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
