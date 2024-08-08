package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestStorage_Basic_WestEurope(t *testing.T) {
	// 🧠 Enabling this test function to run in parallel with others
	t.Parallel()

	// 🧠 Random ID in name allows parallel tests with the same name without collision
	uniqueID := random.UniqueId()
	postfix := strings.ToLower(fmt.Sprintf("netbank%s", uniqueID))
	location := "West Europe"

	tfOptions := &terraform.Options{
		TerraformDir: "../examples/basic-westeurope",
		Vars: map[string]interface{}{
			"postfix":  postfix,
			"location": location,
		},
	}

	defer terraform.Destroy(t, tfOptions)
	terraform.InitAndApply(t, tfOptions)
	resulted_resource_id := terraform.Output(t, tfOptions, "storage_id")
	expected_result := fmt.Sprintf(
		"/subscriptions/28cfe00f-ec8c-4a2c-a32d-037342c9671d/resourceGroups/rg-demo/providers/Microsoft.Storage/storageAccounts/st%s",
		postfix,
	)

	assert.Equal(t, resulted_resource_id, expected_result)

	resulted_stac_object := terraform.OutputMap(t, tfOptions, "stac_object")

	fmt.Print(resulted_stac_object)

	assert.Equal(t, resulted_stac_object["sftp_enabled"], "false")


	// pretty print the result
	jsonString, err := json.MarshalIndent(resulted_stac_object, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))

}