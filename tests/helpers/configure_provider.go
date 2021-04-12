package helpers

import (
	"io"
	"log"
	"os"
)

// Generates provider.tf configuration in module directory
func ConfigureProvider(moduleDir string) {
	providerFile := moduleDir + "/provider.tf"

	providerContent :=
		`terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.49.0"
    }
  }
}
provider "azurerm" {
  features {}
}`

	err := WriteToFile(providerFile, providerContent)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
