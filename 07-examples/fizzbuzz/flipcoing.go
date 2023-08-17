package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func flipCoin() string {
	if rand.Float64() < 0.5 {
		return "Heads"
	}
	return "Tails"
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <coin count> <experiment count> <one tail | all tails>")
		return
	}

	coinCount, err := strconv.Atoi(os.Args[1])
	if err != nil || coinCount <= 0 {
		fmt.Println("Invalid coin count. Please enter a positive integer.")
		return
	}

	experimentCount, err := strconv.Atoi(os.Args[2])
	if err != nil || experimentCount <= 0 {
		fmt.Println("Invalid experiment count. Please enter a positive integer.")
		return
	}

	targetOutcome := os.Args[3]
	if targetOutcome != "one" && targetOutcome != "all" {
		fmt.Println("Invalid target outcome. Please enter 'one' or 'all'.")
		return
	}

	tailCount := 0

	for i := 0; i < experimentCount; i++ {
		tails := 0
		for j := 0; j < coinCount; j++ {
			if flipCoin() == "Tails" {
				tails++
			}
		}

		outcomeAchieved := false
		if targetOutcome == "one" && tails > 0 {
			outcomeAchieved = true
		} else if targetOutcome == "all" && tails == coinCount {
			outcomeAchieved = true
		}

		if outcomeAchieved {
			tailCount++
		}

		fmt.Printf("Experiment %d: %d tails\n", i+1, tails)
	}

	tailProbability := float64(tailCount) / float64(experimentCount)
	fmt.Printf("\nTail Experiment Probability: %.3f", tailProbability)
	allTheoProbability := math.Pow(0.5, float64(coinCount))
	oneTheoProbability := 1 - allTheoProbability
	if targetOutcome == "one" {
		fmt.Printf("\nTail Theoretical one Probability: %.3f", oneTheoProbability)
		fmt.Printf("\nError: %.3f\n", math.Abs(tailProbability-oneTheoProbability))

	} else if targetOutcome == "all" {
		fmt.Printf("\nTail Theoretical all Probability: %.3f", allTheoProbability)
		fmt.Printf("\nError: %.3f\n", math.Abs(tailProbability-allTheoProbability))

	}
}
