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
  access_key = "AKIAZXX7XTCCXGIOZDUK"
  secret_key = "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"
}

# Create a VPC
resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  enable_classiclink_dns_support = true
  enable_dns_hostnames = true
}

resource "aws_internet_gateway" "gw"{
  vpc_id = aws_vpc.vpc.id
}

resource "aws_route_table" "rtable"{
  vpc_id = aws_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }
}


resource "aws_subnet" "subnet1"{
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-west-2a"
}

resource "aws_route_table_association" "a"{
  subnet_id  = aws_subnet.subnet1.id
  route_table_id = aws_route_table.rtable.id
}

resource "aws_security_group" "sg"{
  name = "group1 "
  vpc_id  = aws_vpc.vpc.id

  ingress{
    description = "ssh"
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress{
    description = "http"
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress{
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource  "aws_network_interface" "inteface"{
  subnet_id = aws_subnet.subnet1.id
  private_ips = ["10.0.1.50"]
  security_groups = [aws_security_group.sg.id]
}

resource "aws_eip" "one"{
  vpc = true
  depends_on = [aws_internet_gateway.gw]
  network_interface = aws_network_interface.inteface.id
  associate_with_private_ip = "10.0.1.50"

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

resource "aws_iam_role" "CF2TF-IAM-Role" {
  name = "CF2TF-IAM-Role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_policy" "CF2TF-IAM-Policy" {
  name = "CF2TF-IAM-Policy"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "s3:*",
      "Effect": "Allow",
      "Resource": "*"
     }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "CF2TF-IAM-PA" {
  depends_on = ["aws_iam_role.CF2TF-IAM-Role", "aws_iam_policy.CF2TF-IAM-Policy"]
  role = "${aws_iam_role.CF2TF-IAM-Role.name}"
  policy_arn = "${aws_iam_policy.CF2TF-IAM-Policy.arn}"
}

resource "aws_iam_instance_profile" "CF2TF" {
  name = "CF2TF-IAM-IP"
  depends_on = ["aws_iam_role.CF2TF-IAM-Role"]
  role = "${aws_iam_role.CF2TF-IAM-Role.id}"
}



resource "aws_instance" "testec2"{
  ami = "ami-0c2d06d50ce30b442"
  instance_type = "t2.micro"
  key_name = "Mermaid"
  # subnet_id = aws_subnet.subnet1.id
  # security_groups = [aws_security_group.sg.id]
  # associate_public_ip_address = true
  iam_instance_profile = aws_iam_instance_profile.CF2TF.id

  network_interface {
    network_interface_id = aws_network_interface.inteface.id
    device_index = 0
  }

  tags = {
    Name = "Sasha"
  }

  user_data = <<EOF
              #!/bin/bash
              sudo yum -y update
              sudo yum -y install nginx      
              sudo amazon-linux-extras install -y nginx1
              sudo service nginx start     
              sudo systemctl start nginx
              sudo systemctl enable nginx
              sudo aws s3 cp s3://indexbuckettask7/index.html /usr/share/nginx/html/index.html 
              EOF
  
}