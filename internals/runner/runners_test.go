package runner

import (
	"testing"
)


var authTest = []struct{
	name string
	projectDir string
	stack string
	authName string
	wantErr bool
	errorMsg string
	
}{
	{
        name:       "success case",
        projectDir: ".",
        stack:      "go", 
        authName:   "jwt",
        wantErr:    false,
    },
    {
        name:       "unknown auth",
        projectDir: ".",
        stack:      "fastapi",
        authName:   "invalid-auth",
        wantErr:    true,
        errorMsg:   "unknown auth library: invalid-auth",
    },
    {
        name:       "no auth selected",
        projectDir: ".", 
        stack:      "ts",
        authName:   "none",
        wantErr:    false,
    },	
}


//testing for RunAuth
func TestRunAuth_Cases(t *testing.T){
	for _, tt := range authTest{
		t.Run(tt.name, func(t *testing.T) {
			err := RunAuth(tt.projectDir, tt.stack, tt.authName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected an error but go none")
				}else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("Expected error %q but got %q ", tt.errorMsg, err.Error())
				}
			}else {
				if err != nil {
					t.Errorf("unexpected error: %v",err)
				}
			}
		})
	}
}


//test for run validator
var validatorTest = []struct{
	name string
	projectDir string
	stack string
	validatorName string
	wantErr bool
	errorMsg string
	
}{
	{
        name:       "success case",
        projectDir: ".",
        stack:      "ts", 
        validatorName:   "zod",
        wantErr:    false,
    },
    {
        name:       "unknown validator",
        projectDir: ".",
        stack:      "fastapi",
        validatorName:   "invalid-validator",
        wantErr:    true,
        errorMsg:   "unknown validator: invalid-validator",
    },
    {
        name:       "no validator selected",
        projectDir: ".", 
        stack:      "ts",
        validatorName:   "none",
        wantErr:    false,
    },	
}


func TestRunValidator_Cases(t *testing.T){
	for _, tt := range validatorTest{
		t.Run(tt.name, func(t *testing.T) {
			err := RunValidator(tt.projectDir, tt.stack, tt.validatorName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected an error but go none")
				}else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("Expected error %q but got %q ", tt.errorMsg, err.Error())
				}
			}else {
				if err != nil {
					t.Errorf("unexpected error: %v",err)
				}
			}
		})
	}
}


//test for lint
var lintTest = []struct{
	name string
	projectDir string
	stack string
	lintName string
	wantErr bool
	errorMsg string
	
}{
	{
        name:       "success case",
        projectDir: ".",
        stack:      "ts", 
        lintName:   "eslint",
        wantErr:    false,
    },
    {
        name:       "unknown linter",
        projectDir: ".",
        stack:      "fastapi",
        lintName:   "invalid-linter",
        wantErr:    true,
        errorMsg:   "unknown linter: invalid-linter",
    },
    {
        name:       "no linter selected",
        projectDir: ".", 
        stack:      "ts",
        lintName:   "none",
        wantErr:    false,
    },	
}


func TestRunLint_Cases(t *testing.T){
	for _, tt := range lintTest{
		t.Run(tt.name, func(t *testing.T) {
			err := RunLint(tt.projectDir, tt.stack, tt.lintName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected an error but go none")
				}else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("Expected error %q but got %q ", tt.errorMsg, err.Error())
				}
			}else {
				if err != nil {
					t.Errorf("unexpected error: %v",err)
				}
			}
		})
	}
}



//test for databas
var databaseTest = []struct{
	name string
	projectDir string
	stack string
	databaseName string
	wantErr bool
	errorMsg string
	
}{
	{
        name:       "success case",
        projectDir: ".",
        stack:      "ts", 
        databaseName:   "mongodb",
        wantErr:    false,
    },
    {
        name:       "unknown database",
        projectDir: ".",
        stack:      "fastapi",
        databaseName:   "invalid-database",
        wantErr:    true,
        errorMsg:   "unknown database: invalid-database",
    },
    {
        name:       "no database selected",
        projectDir: ".", 
        stack:      "ts",
        databaseName:   "none",
        wantErr:    false,
    },	
}


func TestRunDatabase_Cases(t *testing.T){
	for _, tt := range databaseTest{
		t.Run(tt.name, func(t *testing.T) {
			err := RunDatabase(tt.projectDir, tt.stack, tt.databaseName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected an error but go none")
				}else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("Expected error %q but got %q ", tt.errorMsg, err.Error())
				}
			}else {
				if err != nil {
					t.Errorf("unexpected error: %v",err)
				}
			}
		})
	}
}

