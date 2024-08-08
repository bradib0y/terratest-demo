module "concat" {
  source = "../.."
}

output "output_from_module" {
  value = module.concat.resource_name
}