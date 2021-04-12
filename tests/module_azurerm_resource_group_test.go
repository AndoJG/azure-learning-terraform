package test

import (
	"azure_terraform_modules_tests/helpers"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestModuleAzurermResourceGroup(t *testing.T) {
	// The value for subscriptionID is replaced by the environment var ARM_SUBSCRIPTION_ID
	subscriptionID := ""
	moduleDir := "../modules/azurerm_resource_group"
	resourceGroupName := "module_test_rg"
	location := "eastus"

	// Write provider configuration to module directory
	helpers.ConfigureProvider(moduleDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to our Terraform module
		TerraformDir: moduleDir,

		// Variables to pass to Terraform
		Vars: map[string]interface{}{
			"name":     resourceGroupName,
			"location": location,
		},
	})

	// Destroy resources when done
	defer terraform.Destroy(t, terraformOptions)

	// Init and apply Terraform code
	terraform.InitAndApply(t, terraformOptions)

	// Check if resource group exists, fails if it does not
	azure.ResourceGroupExists(t, resourceGroupName, subscriptionID)
}
