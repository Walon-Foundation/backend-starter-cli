package prompt

import (
	"os"
	"testing"

)

var option = []string{
	"fastapi",
	"honojs",
	"nextjs",
}

func TestSelectBackend(t *testing.T){
	oldStdin := os.Stdin
	defer func(){ os.Stdin = oldStdin}()

	reader, writer,err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = reader

	go func(){
		defer writer.Close()
		writer.WriteString("2\n")
	}()

	result, err := SelectBackend(option)
	if err != nil {
		t.Errorf("unexpected error: %v\n",err)
		return
	}

	if result != option[1]{
		t.Errorf("Expected %q, got %q\n", option[1], result)
	}
}


func TestProjectName(t *testing.T){
	oldStdin := os.Stdin
	go func(){ os.Stdin = oldStdin}()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	go func(){
		defer writer.Close()
		writer.WriteString("Walon\n")
	}()

	os.Stdin = reader

	result,err := AskProjectName()
	if err != nil {
		t.Errorf("unexpected error: %v",err)
		return
	}

	if result != "Walon" {
		t.Errorf("expected %q, got %q\n", "Walon", result)
	}
}


func TestConfirmGit(t *testing.T){
	oldStdin := os.Stdin
	go func(){ os.Stdin = oldStdin}()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	go func(){
		defer writer.Close()
		writer.WriteString("y\n")
	}()

	os.Stdin = reader

	ok,err := ConfirmGit()
	if err != nil {
		t.Errorf("unexpected error: %v",err)
	}

	if !ok {
		t.Error("expect true but got no")
	}
}