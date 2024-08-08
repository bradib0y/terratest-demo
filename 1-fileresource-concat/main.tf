variable "type" {
  description = "Resource type"
  type        = string
  default     = "vm"
}

variable "location" {
  description = "Location"
  type        = string
  default     = "westeurope"
}

variable "product" {
  description = "Product"
  type        = string
  default     = "netbank"
}

locals {
  resource_name = lower("${var.type}-${var.location}-${var.product}")
}

resource "local_file" "resource_name" {
  content  = local.resource_name
  filename = "resource_name.txt"
}

output "resource_name" {
  value = local.resource_name
}
