package test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// 🧠 define the same directory for multiple functions in global scope (Short variable declarations can only be used inside functions)
var tfdir = "../examples/concat-variables"
var varfile_location = "location.tfvars"

func TestConcat_Variables_Output(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tfdir,
		VarFiles:     []string{varfile_location},
		Vars: map[string]interface{}{
			"type":    "vnet",
			"product": "fraudprevention",
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "output_from_module")
	assert.Equal(t, "vnet-northeurope-fraudprevention", output)
}

func TestConcat_Variables_FileResource(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tfdir,
		VarFiles:     []string{varfile_location},
		Vars: map[string]interface{}{
			"type":    "vnet",
			"product": "fraudprevention",
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	file, err := os.Open("../examples/concat-variables/resource_name.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file content: %s", err)
	}

	content_string := string(content)

	assert.Equal(t, "vnet-northeurope-fraudprevention", content_string)
}
