module "hello" {
  source = "../.."
  subject = "John"
}

output "message_from_module" {
  value = module.hello.message
}