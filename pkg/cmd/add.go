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
	"google.golang.org/api/calendar/v3"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add events to your calendar",
	Long: `Add single-day events to your calendar
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addNewEvent,
}

func addNewEvent(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
	calendarClient := auth.GetClient()
	//taking input from flags
	//TODO: input validation on year
	year, err := cmd.Flags().GetString("year")
	if err != nil {
		fmt.Println("There was a problem with the year specified")
		return
	}
	//TODO: input validation on month
	month, err := cmd.Flags().GetString("month")
	if err != nil {
		fmt.Println("There was a problem with the month specified")
		return
	}
	//TODO: input validation on day
	day, err := cmd.Flags().GetString("day")
	if err != nil {
		fmt.Println("There was a problem with the date specified")
		return
	}
	//TODO: input validation on  time
	time, err := cmd.Flags().GetString("time")
	if err != nil {
		fmt.Println("There was a problem with the time specified")
		return
	}
	startDateTime := fmt.Sprintf("%v-%v-%vT-%v:00.000Z", year, month, day, time)

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		fmt.Println("There was an issue with your description")
	}

	summary, err := cmd.Flags().GetString("time")
	if err != nil {
		fmt.Println("There was an issue with your summary")
	}
	start := calendar.EventDateTime{
		DateTime: startDateTime,
		TimeZone: "America/New_York",
	}
	end := calendar.EventDateTime{
		DateTime: startDateTime,
		TimeZone: "America/New_York",
	}
	event := calendar.Event{
		Summary:     summary,
		Description: description,
		Start:       &start,
		End:         &end,
	}
	_, err = calendarClient.Events.Insert("primary", &event).Do()
	if err == nil {
		fmt.Println("shit worked")
	} else {
		fmt.Println(err)
	}

}
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("year", "y", "2020", "year for new event")
	addCmd.Flags().StringP("month", "m", "0", "month for new event, please put the index of the month in")
	addCmd.Flags().Int("day", 0, "day for new event")
	addCmd.Flags().StringP("time", "t", "", "time formatted as HH:MM. if unset will be a full day event")
	addCmd.Flags().StringP("description", "desc", "testFlag", "[optional] longer description of new event")
	addCmd.Flags().String("summary", "sum", "short, high-level overview of the event")
}
