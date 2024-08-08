variable "subject" {
  description = "Subject of welcome message"
  type        = string
  default     = "World"

  validation {
    condition     = length(var.subject) <= 20
    error_message = "The 'subject' variable must not be longer than 20 characters."
  }
}

output "message" {
  value = "Hello, ${var.subject}!"
}