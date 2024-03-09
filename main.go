package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	// Beispiel, wie man eine Gewohnheit mit einem spezifischen Ziel erstellen könnte
	habits := []Habit{
		{
			Name: "Lernen",
			Goal: QuanityGoal{Unit: "Seiten", Quantity: 10, Timeframe: Day},
		},
		{
			Name: "Laufen",
			Goal: NTimesInTimeframeGoal{Quantity: 3, Timeframe: Week},
		},
		{
			Name: "Saugen",
			Goal: EveryNDaysGoal{IntervalDays: 2},
		},
	}

	for _, habit := range habits {
		logHabit(habit)
	}

	habits[0].newAction(2) // zwei Seiten gelesen
	time.Sleep(2 * time.Second)
	habits[0].newAction(1) // zwei Seiten gelesen

	log.Println("----")

	for _, habit := range habits {
		logHabit(habit)
	}
}

func logHabit(habit Habit) {
	log.Println(habit.Name)
	log.Printf("Goal: %s\n", habit.Goal.Describe())
	if habit.History.Actions == nil {
		log.Println("No history yet")
		return
	}
	log.Printf("History: %s\n", habit.History.Describe())
	log.Println(habit.History.Details())
}
