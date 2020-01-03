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

	"github.com/kvptkr/gocal-cli/pkg/auth"
	"github.com/spf13/cobra"
)

//TODO: Add a flag to get coloured output consistent with a respective calendar's colour
// listCmd represents the list command
//`This gets a list of all available calendars from Google Calendar`,
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of all available calendars",
	Long:  `This gets a list of all available calendars from Google Calendar`,
	Run: func(cmd *cobra.Command, args []string) {
		counter := 1
		fmt.Println("Available Calendars:")
		calendarClient := auth.GetClient()
		calendars, err := calendarClient.CalendarList.List().Do()
		if err != nil {
			fmt.Println("There was an error getting the available calendars")
		}
		if calendars.Items == nil {
			fmt.Println("Sorry, no calendars were found")
		} else {
			for _, calendarName := range calendars.Items {
				fmt.Println(fmt.Sprintf("[%d]", counter) + calendarName.Summary)
				counter++
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
