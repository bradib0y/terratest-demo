package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTooLongErrorMessage(t *testing.T) {
	// t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/hello-too-long",
	}

	// 🧠 Run terraform init and plan, this time with the "E" version of the function
	_, err := terraform.InitAndPlanE(t, terraformOptions)

	// 🧠 Check if an error occurred
	assert.Error(t, err)

	// 🧠 Check if the error message matches the expected message
	expectedErrorMessage := "The 'subject' variable must not be longer than 20 characters."
	assert.True(t, strings.Contains(err.Error(), expectedErrorMessage), "The error message does not match the expected error message.")
}
