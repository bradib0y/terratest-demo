package test

import (
	"fmt"
	"os"
	"testing"
)

//var args = os.Args

func TestDummy(t *testing.T) {
	// Get the command-line arguments passed to the test
	args := os.Args

	// Print the arguments to the console
	fmt.Println("Args:", args)

	// Process the arguments
	for i, arg := range args {
		fmt.Printf("Arg %d: %s\n", i+1, arg)
	}

	// conditional
	var tfLocalInfo string
	if IsTfLocal() {
		tfLocalInfo = "TfLocal: YES"
	} else {
		tfLocalInfo = "TfLocal: NO"
	}

	fmt.Println("--- ", tfLocalInfo)

}

func IsTfLocal() bool {
	args := os.Args

	for _, argValue := range args {
		if argValue == "tf-local-dev" {
			return true
		}
	}

	return false
}
