resource "aws_cloudwatch_metric_alarm" "recovery1"{
    alarm_name = "recovery1"
    evaluation_periods = "2"
    period = "180"
    alarm_description = "auto recovery alarm"
    alarm_actions = ["arn:aws:automate:${var.region}:ec2:recover"]
    comparison_operator = "GreaterThanThreshold"
    threshold = "0.0"
    metric_name = "StatusCheckFailed_System"

    dimensions =  {
      InstanceId = aws_instance.nginx1.id
    }
}

resource "aws_cloudwatch_metric_alarm" "recovery2"{
    alarm_name = "recovery2"
    evaluation_periods = "2"
    period = "180"
    alarm_description = "auto recovery alarm"
    alarm_actions = ["arn:aws:automate:${var.region}:ec2:recover"]
    comparison_operator = "GreaterThanThreshold"
    threshold = "0.0"
    metric_name = "StatusCheckFailed_System"

    dimensions = {
      InstanceId = aws_instance.nginx2.id
    }
    
}
