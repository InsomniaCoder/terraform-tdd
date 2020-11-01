package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformACRProvisionShouldHaveCorrectSetup(t *testing.T) {
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	expectedName := "unitTestACR"
	expectedResourceGroup := "test-rg"
	expectedLocation := "southeastasia"
	expectedSKU := "Standard"

	actualName := terraform.Output(t, terraformOptions, "acr_name")
	actualResourceGroup := terraform.Output(t, terraformOptions, "rg_name")
	actualLocation := terraform.Output(t, terraformOptions, "acr_location")
	actualSKU := terraform.Output(t, terraformOptions, "acr_sku")

	assert.Equal(t, expectedName, actualName)
	assert.Equal(t, expectedResourceGroup, actualResourceGroup)
	assert.Equal(t, expectedLocation, actualLocation)
	assert.Equal(t, expectedSKU, actualSKU)
}
