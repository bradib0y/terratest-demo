module "hello" {
  source = "../.."
  subject = "Hungary"
}

output "message_from_module" {
  value = module.hello.message
}