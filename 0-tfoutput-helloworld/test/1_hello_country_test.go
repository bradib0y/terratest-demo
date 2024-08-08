package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestHelloCountry(t *testing.T) {
	// 🧠 The tests which are marked with this will be run in parallel - saves time.
	// 🧠 Make sure that your parallel tests are not interfering with each other in any way
	// t.Parallel()

	// 🧠 Terraform options - these are manipulating how we run Terraform. (eg. which folder, which .tfvars file)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello-country",
	})

	// 🧠 Terraform destroy command - cleans up the resources created during the test
	// 🧠 Use 'defer' keyword to run this after current scope, after test function returned, or even in case of error
	defer terraform.Destroy(t, terraformOptions)

	// 🧠 Terraform main commands
	terraform.InitAndApply(t, terraformOptions)

	// 🧠 Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "message_from_module")
	assert.Equal(t, "Hello, Hungary!", output)
}
