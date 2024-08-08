package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestHelloPerson(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello-person",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "message_from_module")
	assert.Equal(t, "Hello, John!", output)
}

// 🧠 multiple Test functions
func TestHelloPerson2(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello-person",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// 🧠 multiple asserts
	output := terraform.Output(t, terraformOptions, "message_from_module")
	assert.Contains(t, output, "Hello")
	assert.Contains(t, output, ",")
	assert.Contains(t, output, "John")
	assert.Contains(t, output, "!")
}
