package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectBackend(options []string)(string, error){
	fmt.Println("-----Choose a backend starter------")

	for i, opt := range options {
		fmt.Printf("%d), %s\n",i+1, opt)
	}

	fmt.Print("\nEnter choice (number):")

	reader := bufio.NewReader((os.Stdin))
	input, err := reader.ReadString('\n')
	if err != nil {
		return "",err
	}

	input = strings.TrimSpace(input)

	choice,err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(options){
		return "", fmt.Errorf("invalid choice: %s", input)
	}

	return options[choice -1], nil
}

func AskProjectName()(string, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter project name: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "",err
	}

	return strings.TrimSpace(input),nil
}


//ask for database
func AskDatabase( option []string)(string, error){
	fmt.Println("-----Choice a database------")
	for i, opt := range option {
		fmt.Printf(" %d). %s\n",i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you choice: ")
	choice,err := reader.ReadString('\n')
	if err != nil{
		return "",err
	}
	choice = strings.TrimSpace(choice)

	value,err := strconv.Atoi(choice)
	if err != nil && value > len(option) && value < 1 {
		return "",err
	}

	return option[value - 1], nil
}

//ask for auth
func AskForAuth(options []string)(string,error){
	fmt.Println("-----Choose an Authentication libray-----")
	for i ,opt := range options{
		fmt.Printf(" %d). %v\n",i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your choice: ")
	input,err := reader.ReadString('\n')
	if err != nil {
		return "",err
	}
	input = strings.TrimSpace(input)

	choice,err := strconv.Atoi(input)
	if err != nil && choice > len(options) && choice < 1 {
		return "", nil
	}

	return options[choice - 1], nil
}

//ask for linting
//ask for validation
func AskForValidator(options []string)(string,error){
	fmt.Println("-----Choose a Validator libray-----")
	for i ,opt := range options{
		fmt.Printf(" %d). %v\n",i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your choice: ")
	input,err := reader.ReadString('\n')
	if err != nil {
		return "",err
	}
	input = strings.TrimSpace(input)

	choice,err := strconv.Atoi(input)
	if err != nil && choice > len(options) && choice < 1 {
		return "", nil
	}

	return options[choice - 1], nil
}


func ConfirmGit()(bool, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Initialize Git repository? (y/n)")
	input,err := reader.ReadString('\n')
	if err != nil {
		return false,err
	}

	input = strings.TrimSpace(input)

	return input == "y" || input == "Y", nil
}