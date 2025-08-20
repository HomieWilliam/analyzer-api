variable "instance_types" {
  type    = list(string)
  default = ["t3.nano", "t3.micro", "t2.micro"]  # lista de tipos aceitáveis
}

# Data source para pegar a primeira instância disponível
data "aws_instance_type_offerings" "available" {
  filter {
    name   = "instance-type"
    values = var.instance_types
  }
  location_type = "region"
}

resource "aws_instance" "app_server" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = data.aws_instance_type_offerings.available.instance_types[0]
  key_name      = "analyzer-api-key"

  tags = {
    Name = "analyzer-api"
  }
}
