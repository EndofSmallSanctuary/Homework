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
  access_key = "AKIAZXX7XTCC6U3SZR7N"
  secret_key = "Nh/DnldVqIad+WCfOL55LJSp+QQIrovyQjlbhrI"
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

resource "aws_subnet" "subnet2"{
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "us-west-2b"
}

resource "aws_route_table_association" "a"{
  subnet_id  = aws_subnet.subnet1.id
  route_table_id = aws_route_table.rtable.id
}

resource "aws_route_table_association" "b"{
  subnet_id  = aws_subnet.subnet2.id
  route_table_id = aws_route_table.rtable.id
}

resource "aws_security_group" "sg"{
  name = "sshhttpgroup"
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

resource  "aws_network_interface" "inteface2"{
  subnet_id = aws_subnet.subnet2.id
  private_ips = ["10.0.2.50"]
  security_groups = [aws_security_group.sg.id]
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

resource "aws_iam_role" "s3bucketrole" {
  name = "s3Role"
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

resource "aws_iam_policy" "s3bucketpolicy" {
  name = "s3Policy"
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

resource "aws_iam_role_policy_attachment" "policyattachment" {
  depends_on = [aws_iam_role.s3bucketrole, aws_iam_policy.s3bucketpolicy]
  role = aws_iam_role.s3bucketrole.name
  policy_arn = aws_iam_policy.s3bucketpolicy.arn
}

resource "aws_iam_instance_profile" "s3profile" {
  name = "s3profile"
  depends_on = [aws_iam_role.s3bucketrole]
  role = aws_iam_role.s3bucketrole.id
}



resource "aws_instance" "nginx1"{
  ami = "ami-0c2d06d50ce30b442"
  instance_type = "t2.micro"
  key_name = "Mermaid"
 // subnet_id = aws_subnet.subnet1.id
  # vpc_security_grous_ids = [aws_security_group.sg.id]
  # associate_public_ip_address = true
  iam_instance_profile = aws_iam_instance_profile.s3profile.id

  network_interface {
    network_interface_id = aws_network_interface.inteface.id
    device_index = 0
  }

  tags = {
    Name = "Nginx1"
  }

  user_data = <<EOF
              #!/bin/bash
              sudo yum -y update
              sudo yum -y install nginx      
              sudo amazon-linux-extras install -y nginx1
              sudo service nginx start     
              sudo systemctl start nginx
              sudo systemctl enable nginx
              sudo aws s3 cp s3://${var.bucketname}/index.html /usr/share/nginx/html/index.html 
              EOF
  
}



resource "aws_eip" "one"{
  vpc = true
  depends_on = [aws_instance.nginx1]
  network_interface = aws_network_interface.inteface.id
  associate_with_private_ip = "10.0.1.50"

}

resource "aws_instance" "nginx2"{
  ami = "ami-0c2d06d50ce30b442"
  instance_type = "t2.micro"
  key_name = "Mermaid"
 // subnet_id = aws_subnet.subnet2.id
  # vpc_security_grous_ids = [aws_security_group.sg.id]
  # associate_public_ip_address = true
  iam_instance_profile = aws_iam_instance_profile.s3profile.id

  network_interface {
    network_interface_id = aws_network_interface.inteface2.id
    device_index = 0
  }

  tags = {
    Name = "Nginx2"
  }

  user_data = <<EOF
              #!/bin/bash
              sudo yum -y update
              sudo yum -y install nginx      
              sudo amazon-linux-extras install -y nginx1
              sudo service nginx start     
              sudo systemctl start nginx
              sudo systemctl enable nginx
              sudo aws s3 cp s3://${var.bucketname}/index.html /usr/share/nginx/html/index.html 
              EOF
  
}

resource "aws_eip" "two"{
  vpc = true
  depends_on = [aws_instance.nginx2]
  network_interface = aws_network_interface.inteface2.id
  associate_with_private_ip = "10.0.2.50"

}

resource "aws_elb" "balancer"{
  name = "balancer"
  subnets = [aws_subnet.subnet1.id,aws_subnet.subnet2.id]
  security_groups = [aws_security_group.sg.id]

  listener {
    instance_port = 80
    instance_protocol = "http"
    lb_port = 80
    lb_protocol = "http"
  }

  health_check {
    healthy_threshold   = 2
    unhealthy_threshold = 2
    timeout             = 3
    target              = "HTTP:80/"
    interval            = 30
  }

  instances = [aws_instance.nginx1.id,aws_instance.nginx2.id]
  cross_zone_load_balancing = true
  idle_timeout = 600
  connection_draining = true
  connection_draining_timeout = 600
}
