package prompt

import (
	"os"
	"testing"

)

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