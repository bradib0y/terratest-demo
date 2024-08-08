# We strongly recommend using the required_providers block to set the
# Azure Provider source and version being used
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.47.0"
    }
  }
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}

  subscription_id = "28cfe00f-ec8c-4a2c-a32d-037342c9671d"
  tenant_id       = "ec6edc5f-33e6-46a3-afe2-7ba945d30c35"
}