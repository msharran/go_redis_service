output "server_dns" {
  value = "http://${aws_instance.go_redis_instance.public_dns}"
}