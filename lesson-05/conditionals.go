package main

import "fmt"

func main() {

	/* single if statement */
	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	/* if - else block */
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote yet")
	}

	/* if - else if - else block */
	score := 95
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	/* short if declaration */
	if x, y := 10, 20; x < y {
		fmt.Println("x is smaller than y")
	}

	/*
		mixing types is not allowed in short if declaration -->

		if x := 10, y := "Go"; x < 20 {
			fmt.Println(y)
		}
	*/

	/* switch statements */
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Weekend is near!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("A regular day")
	}

	/* we can combine multiple cases */
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day")
	}
}
