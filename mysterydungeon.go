package main

import (
	"fmt"
)

func askQuestion1() string {
	
	fmt.Println("Question 1: When faced with danger, what is your first instinct?")
	fmt.Println("a) Stand and fight")
	fmt.Println("b) Try to negotiate")
	fmt.Println("c) Run away")
	fmt.Println("")

	
	var response1 string
	fmt.Scanln(&response1)
	if response1 != "a" && response1 != "b" && response1 != "c" {
		fmt.Println("Invalid response. Please enter a, b, or c.")
		response1 = ""
	}

	return response1
}

func askQuestion2() string {
	
	fmt.Println("Question 2: Which of these environments do you feel most comfortable in?")
	fmt.Println("a) Forest")
	fmt.Println("b) City")
	fmt.Println("c) Beach")
	fmt.Println("")

	
	var response2 string
	fmt.Scanln(&response2)
	if response2 != "a" && response2 != "b" && response2 != "c" {
		fmt.Println("Invalid response. Please enter a, b, or c.")
		response2 = ""
	}

	return response2
}

func askQuestion3() string {
	
	fmt.Println("Question 3: Which of these abilities would you prefer to have?")
	fmt.Println("a) Invisibility")
	fmt.Println("b) Super strength")
	fmt.Println("c) Teleportation")
	fmt.Println("")

	
	var response3 string
	fmt.Scanln(&response3)
	if response3 != "a" && response3 != "b" && response3 != "c" {
		fmt.Println("Invalid response. Please enter a, b, or c.")
		response3 = ""
	}

	return response3
}

func main() {
	
	response1 := askQuestion1()
	response2 := askQuestion2()
	response3 := askQuestion3()

	
	fullResponse := response1 + response2 + response3

	
	var pokemon string
	switch fullResponse {
	case "aaa":
		pokemon = "Charmander"
	case "aab":
		pokemon = "Bulbasaur"
	case "aac":
		pokemon = "Squirtle"
	case "abb":
		pokemon = "Chikorita"
	case "abc":
		pokemon = "Cyndaquil"
	case "acc":
		pokemon = "Totodile"
	case "bbb":
		pokemon = "Treeko"
	case "bbc":
		pokemon = "Torchic"
	case "bcc":
		pokemon = "Mudkip"
	case "ccc":
		pokemon = "Pikachu"
	default:
		fmt.Println("Sorry, there was an error. Please try again.")
		return
	}

	
	fmt.Println("Congratulations, you are a", pokemon, "!")
}
