provider "aws" {
  region = "us-east-1"
  profile = "default"
}

resource "aws_instance" "go_redis_instance" {
	ami = "ami-09e67e426f25ce0d7"
	instance_type = "t2.micro"
	key_name = "sharrans-root-useast1-kp"
	user_data = file("user_data.sh")
	tags = {
		Environment = "dev"
	}
	vpc_security_group_ids = [aws_security_group.go_redis_service_sg.id]
}