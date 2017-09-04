variable "profile" {}
variable "region" {}
variable "shared_credentials_file" {}

variable "dd_api_key" {}
variable "image_id" {}
variable "instance_type" {}
variable "asg_size" { default = "1" }

provider "aws" {
  shared_credentials_file = "${var.shared_credentials_file}"
  profile = "${var.profile}"
  region = "${var.region}"
}

resource "aws_launch_configuration" "watirspec-launch-config" {
  name = "watirspec-launch-config"
  image_id =  "${var.image_id}"
  instance_type = "${var.instance_type}"
  user_data = "${replace(file("cloudconfig.yml"), "DD_API_KEY", "${var.dd_api_key}")}"
  security_groups = ["allow_all_for_watirspec"]
}

resource "aws_autoscaling_group" "watirspec-asg" {
  availability_zones = ["us-west-2a", "us-west-2b", "us-west-2c"]
  name = "watirspec-asg"
  max_size = "${var.asg_size}"
  min_size = "${var.asg_size}"
  health_check_grace_period = 600
  health_check_type = "ELB"
  desired_capacity = "${var.asg_size}"
  force_delete = true
  launch_configuration = "${aws_launch_configuration.watirspec-launch-config.name}"
  load_balancers = ["watirspec-elb"]
  tag {
    key = "Name"
    value = "watirspec"
    propagate_at_launch = true
  }
}
