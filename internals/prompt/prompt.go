package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectBackend(options []string)(string, error){
	fmt.Println("Choose a backend starter")

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


func SelectDependecies(options []string)([]string, error){
	fmt.Println(options)
	fmt.Println("Select additional dependencies..")

	for i, opts := range options{
		fmt.Printf(" %d) %s\n", i+1, opts)
	}

	fmt.Print("\n Enter numbers (eg 1,2,3): ")

	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return nil,nil
	}

	var dep []string

	for _, part := range strings.Split(input, ","){
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil && num >=1 && num <= len(options){
			dep = append(dep, options[num - 1])
		}
	}

	return  dep, nil
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