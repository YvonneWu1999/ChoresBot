package main

import (
	"ChoresBot/helper"
	"fmt"
)

// only support assigning chores to same amount of people
func AssignChores(weekNum int, year int, people []string, chores []string) map[string]string {
	assignments := make(map[string]string)
	if len(people) != len(chores) {
		fmt.Println("Assign Error: Chores amount must equal to the amount of people")
		return assignments
	}

	turn := weekNum % len(people)
	// handle cross year
	// startYear := 2023
	// yearGap := 0
	// if currentTime.Year() != startYear {
	// 	yearGap = currentTime.Year() - startYear
	// }

	for i, person := range people {
		assignments[person] = chores[(turn+i)%len(people)]
	}
	helper.LogInfo(fmt.Sprintf("week: %v", weekNum))
	helper.LogInfo(fmt.Sprintf("turn: %v", turn))
	helper.LogInfo(fmt.Sprintf("assignments: %v", assignments))

	return assignments
}
