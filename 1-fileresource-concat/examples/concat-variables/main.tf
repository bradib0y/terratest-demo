variable "type" {
  description = "Resource type"
  type        = string
}

variable "location" {
  description = "Location"
  type        = string
}

variable "product" {
  description = "Product"
  type        = string
}

module "concat" {
  source = "../.."
  type = var.type
  location = var.location
  product = var.product
}

output "output_from_module" {
  value = module.concat.resource_name
}