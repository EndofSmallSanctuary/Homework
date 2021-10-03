resource "aws_instance" "nginx1"{
  ami = "ami-0c2d06d50ce30b442"
  instance_type = "t2.micro"
  key_name = "Rose"
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
  key_name = "Rose"
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