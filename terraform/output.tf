output "server_dns" {
  value = aws_instance.go_redis_instance.public_dns
}