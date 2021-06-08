variable "aws_region" {
  description = "Region in which aws resources will be created"
  default = "ap-south-1"
}

variable "aws_cli_profile" {
  description = "Profile name of Aws CLI"
  default = "default"
}

variable "aws_ubuntu_ami" {
  description = "Ubuntu AMI Id"
  default = "ami-0c1a7f89451184c8b"
}

variable "environment" {
  description = "Tag used for environment segregation"
  default = "development"
}

variable "server_port" {
  description = "Server port for security group"
  default = 80
}