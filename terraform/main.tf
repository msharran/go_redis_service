provider "aws" {
  region = var.aws_region
  profile = var.aws_cli_profile
}

resource "aws_instance" "go_redis_instance" {
	ami = var.aws_ubuntu_ami
	instance_type = "t2.micro"
	user_data = file("user_data.sh")
	tags = {
		Environment = var.environment
	}
	vpc_security_group_ids = [aws_security_group.go_redis_service_sg.id]
}