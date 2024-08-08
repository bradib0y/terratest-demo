variable "postfix" {
  description = "postfix"
  default     = "john99"
}

variable "location" {
  description = "location"
  default     = "West Europe"
}

module "storage_account" {
  source   = "../.."
  location = var.location
  postfix  = var.postfix
}

output "storage_id" {
  value = module.storage_account.storage_account_id
}

output "stac_object" {
  description = "Output object of Storage Account"
  value       = module.storage_account.stac_object
  sensitive   = true
}
