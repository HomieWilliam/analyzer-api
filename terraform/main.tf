provider "aws" {
  region = "eu-north-1"
}

variable "docker_registry" {
  type = string
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }
}

resource "aws_instance" "app_server" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.nano"
  key_name      = "analyzer-api-key"

  # User data para instalar Docker e rodar o container
  user_data = <<-EOF
              #!/bin/bash
              apt-get update
              apt-get install -y docker.io
              systemctl enable docker
              systemctl start docker
              docker run -d -p 8080:8080 ${var.docker_registry}/analyzer-api:latest
              EOF

  tags = {
    Name = "analyzer-api"
  }
}