module "hello" {
  source = "../.."
  subject = "VeryVeryLongNameJohnny"
}

output "message_from_module" {
  value = module.hello.message
}