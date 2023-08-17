package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateCalories(weight float64, distance float64, speed string) float64 {
	caloriesPerKm := 0.75 // Adjust this value based on your needs
	speedMultiplier := map[string]float64{
		"normal": 1.0,
		"quick":  1.3, // 30% more intense than normal walking
		"fast":   1.5, // 50% more intense than normal walking
		"run30":  2.0, // Example value for running pace
		"run50":  2.5, // Example value for running pace
		"run70":  3.0, // Example value for running pace
		"run100": 4.0, // Alien Mode
	}

	if multiplier, ok := speedMultiplier[strings.ToLower(speed)]; ok {
		return caloriesPerKm * distance * weight * multiplier
	}
	return 0.0
}

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Usage: weight_loss_app <weight in kg> <distance in km> <speed> <days> <daily intake calories>")
		return
	}

	weight, err1 := strconv.ParseFloat(os.Args[1], 64)
	distance, err2 := strconv.ParseFloat(os.Args[2], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input for weight or distance.")
		return
	}

	speed := os.Args[3]
	days, err3 := strconv.Atoi(os.Args[4])
	dailyIntakeCalories, err4 := strconv.Atoi(os.Args[5])

	if err3 != nil || err4 != nil {
		fmt.Println("Invalid input for days or daily intake calories.")
		return
	}

	fmt.Printf("| %-4s | %-8s | %-15s | %-20s | %-22s | %-24s |\n", "Days", "Weight (kg)", "Calorie Burned", "Daily Calories Intake", "Cumulative Calorieburned", "Cumulative Calories Intake")
	fmt.Println(strings.Repeat("-", 120))

	file, err := os.Create("weight_loss_table.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Days", "Weight (kg)", "Calorie Burned", "Daily Calories Intake", "Cumulative Calorieburned", "Cumulative Calories Intake"})

	cumulativeCaloriesBurned := 0.0
	cumulativeCaloriesIntake := 0.0
	totalWeightLoss := 0.

	for day := 1; day <= days; day++ {
		caloriesBurned := calculateCalories(weight, distance, speed)
		caloriesIntake := float64(dailyIntakeCalories)

		cumulativeCaloriesBurned += caloriesBurned
		cumulativeCaloriesIntake += caloriesIntake

		weightLoss := caloriesBurned / 7700  // 7700 calories ≈ 1 kg
		updatedWeight := weight - weightLoss // Subtract weight loss from current weight

		if day%2 == 0 {
			formatString := "| %-4d | %-8.2f   | %-15.2f  | %-20d  | %-22.2f   | %-24.2f   |\n"
			formattedRow := fmt.Sprintf(formatString, day, updatedWeight, caloriesBurned, dailyIntakeCalories, cumulativeCaloriesBurned, cumulativeCaloriesIntake)
			fmt.Printf("\x1b[32m%s\x1b[0m", formattedRow)

		} else {
			fmt.Printf("| %-4d | %-8.2f   | %-15.2f  | %-20d  | %-22.2f   | %-24.2f   |\n", day, updatedWeight, caloriesBurned, dailyIntakeCalories, cumulativeCaloriesBurned, cumulativeCaloriesIntake)
		}

		row := []string{
			strconv.Itoa(day),
			fmt.Sprintf("%.2f", updatedWeight),
			fmt.Sprintf("%.2f", caloriesBurned),
			strconv.Itoa(dailyIntakeCalories),
			fmt.Sprintf("%.2f", cumulativeCaloriesBurned),
			fmt.Sprintf("%.2f", cumulativeCaloriesIntake),
		}

		writer.Write(row)

		weight = updatedWeight
		totalWeightLoss += weightLoss
	}

	fmt.Println(strings.Repeat("-", 120))

	fmt.Println("Total Weight Loss:", totalWeightLoss, "kg")
	fmt.Println("CSV file 'weight_loss_table.csv' created successfully.")
}
