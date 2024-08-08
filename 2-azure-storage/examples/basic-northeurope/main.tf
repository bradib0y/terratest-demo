variable "postfix" {
  description = "postfix"
  default     = "john99"
}

variable "location" {
  description = "location"
  default     = "North Europe"
}

module "storage_account" {
  source   = "../.."
  location = var.location
  postfix   = var.postfix
}

output "storage_id" {
  value = module.storage_account.storage_account_id
}
