resource "aws_security_group" "go_redis_service_sg" {
  name        = "go_redis_service_sg"
  description = "controls access to the ALB"
  vpc_id      = aws_default_vpc.default.id

  ingress {
    protocol    = "tcp"
    from_port   = var.server_port
    to_port     = var.server_port
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}

#AWS Managed default VPC
resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}