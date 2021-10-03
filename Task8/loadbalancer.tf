
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
