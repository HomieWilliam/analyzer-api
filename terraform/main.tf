provider "aws" {
  region = "eu-north-1"
}

resource "aws_instance" "app_server" {
  ami           = "ami-07b36ea9852e986ad"
  instance_type = "t3.micro"
  key_name      = "analyzer-api-key"

  tags = {
    Name = "analyzer-api"
  }
}