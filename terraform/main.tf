provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "app_server" {
  ami           = "ami-0c4fc5dcabc9df21d"
  instance_type = "t3.micro"
  key_name      = "analyzer-api-key"

  tags = {
    Name = "analyzer-api"
  }
}