resource "azurerm_storage_account" "example" {
  name                     = "st${var.postfix}"
  resource_group_name      = "rg-demo"
  location                 = var.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  network_rules {
    default_action = "Deny"
    ip_rules       = ["94.27.0.0/16"]
  }
}

output "storage_account_id" {
  description = "ID of storage account"
  value       = azurerm_storage_account.example.id
}

output "stac_object" {
  description = "Output object of Storage Account"
  value       = azurerm_storage_account.example
}