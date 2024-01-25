package main

import (
	"ChoresBot/helper"
	"fmt"
)

// only support assigning chores to same amount of people
func AssignChores(weekNum int, year int, people []string, chores []string) map[string]string {
	assignments := make(map[string]string)
	subChores := divideChores(chores, len(people))
	turn := weekNum % len(people)
	// TODO: handle cross year
	// startYear := 2023
	// yearGap := 0
	// if currentTime.Year() != startYear {
	// 	yearGap = currentTime.Year() - startYear
	// }

	for i, person := range people {
		for _, subChore := range subChores {
			assignments[person] = assignments[person] + subChore[(turn+i)%len(people)]
		}

	}
	helper.LogInfo(fmt.Sprintf("week: %v", weekNum))
	helper.LogInfo(fmt.Sprintf("turn: %v", turn))
	helper.LogInfo(fmt.Sprintf("assignments: %v", assignments))

	return assignments
}

func divideChores(chores []string, peopleAmount int) [][]string {
	rows := (len(chores) + peopleAmount - 1) / peopleAmount
	result := make([][]string, rows)

	for i := 0; i < rows; i++ {
		start := i * peopleAmount
		end := (i + 1) * peopleAmount
		if end > len(chores) {
			end = len(chores)
		}
		chunk := make([]string, peopleAmount)
		// Copy elements from the original slice to the chunk
		copy(chunk, chores[start:end])
		// Append empty strings to fill the chunk
		for j := end - start; j < peopleAmount; j++ {
			chunk[j] = ""
		}
		result[i] = chunk
	}

	return result
}
