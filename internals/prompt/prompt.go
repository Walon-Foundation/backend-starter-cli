package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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