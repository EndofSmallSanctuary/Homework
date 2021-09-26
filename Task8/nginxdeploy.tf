terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-west-2"
  access_key = "AKIAZXX7XTCC327D2D6H"
  secret_key = "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"
}

# Create a VPC
resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "subnet1"{
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Subnet 1 Terraform"
  }
}

resource "aws_subnet" "subnet2"{
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "Subnet 2 Terraform"
  }
}

resource "aws_security_group" "security_group"{
  name = "sshttpgroup"
  description = "allow security group"
  vpc_id = aws_vpc.vpc.id

  ingress {
    description = "SSH"
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTP"
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
     from_port        = 0
     to_port          = 0
     protocol         = "-1"
     cidr_blocks      = ["0.0.0.0/0"]
     ipv6_cidr_blocks = ["::/0"]
  }
  



}

# resource "tls_private_key" "pk" {
#   algorithm = "RSA"
#   rsa_bits  = 4096
# }

# resource "aws_key_pair" "kp" {
#   key_name   = "myKey"       # Create a "myKey" to AWS!!
#   public_key = tls_private_key.pk.public_key_openssh

#   provisioner "local-exec" { # Create a "myKey.pem" to your computer!!
#     command = "echo '${tls_private_key.pk.private_key_pem}' > myKey.pem"
#   }
# }



resource "aws_instance" "testec2"{
  ami = "ami-0c2d06d50ce30b442"
  instance_type = "t2.micro"
  key_name = "Mermaid"
  subnet_id = aws_subnet.subnet1.id
  associate_public_ip_address = true
  security_groups = [aws_security_group.security_group.id]

  tags = {
    Name = "Sasha"
  }

  user_data = <<-EOF
              #!/bin/bash
              sudo yum -y update
              sudo yum -y install nginx      
              sudo service nginx start     
              sudo systemctl start nginx
              sudo systemctl enable nginx
              sudo aws s3 cp s3://indexbuckettask7/index.html /usr/share/nginx/html/index.html 
              EOF
  
}