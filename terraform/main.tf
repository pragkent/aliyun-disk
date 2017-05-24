provider "alicloud" {
  region = "cn-shanghai"
}

data "alicloud_images" "image" {
  most_recent = true
  owners      = "self"
  name_regex  = "^ubuntu-server-"
}

resource "alicloud_security_group" "test" {
  name        = "internet_gateway"
  description = "security group for internet gateway servers"
}

resource "alicloud_security_group_rule" "test_allow_ingress_tcp_ssh" {
  type              = "ingress"
  ip_protocol       = "tcp"
  nic_type          = "internet"
  policy            = "accept"
  port_range        = "22/22"
  priority          = 50
  security_group_id = "${alicloud_security_group.test.id}"
  cidr_ip           = "0.0.0.0/0"
}

resource "alicloud_security_group_rule" "test_allow_ingress_icmp_all" {
  type              = "ingress"
  ip_protocol       = "icmp"
  nic_type          = "internet"
  policy            = "accept"
  port_range        = "-1/-1"
  priority          = 100
  security_group_id = "${alicloud_security_group.test.id}"
  cidr_ip           = "0.0.0.0/0"
}

resource "alicloud_security_group_rule" "test_allow_egress_all" {
  type              = "egress"
  ip_protocol       = "all"
  nic_type          = "internet"
  policy            = "accept"
  port_range        = "-1/-1"
  priority          = 100
  security_group_id = "${alicloud_security_group.test.id}"
  cidr_ip           = "0.0.0.0/0"
}

resource "alicloud_instance" "test" {
  availability_zone = "cn-shanghai-b"
  instance_name        = "aliyun-disk-test"
  host_name            = "aliyun-disk-test"
  instance_type        = "ecs.n4.small"
  io_optimized         = "optimized"
  image_id             = "${data.alicloud_images.image.images.0.id}"
  system_disk_category = "cloud_efficiency"
  security_groups      = ["${alicloud_security_group.test.*.id}"]
  internet_charge_type = "PayByTraffic"
  allocate_public_ip = true
  internet_max_bandwidth_out = 10

  tags {
    hostname = "aliyun-disk-test"
  }
}

resource "alicloud_disk" "disk" {
  availability_zone = "cn-shanghai-b"
  name              = "aliyun-disk-test"
  description       = "aliyun-disk-test"
  category          = "cloud_efficiency"
  size              = "30"
}

output "instance_id" {
  value = "${alicloud_instance.test.id}"
}

output "instance_public_ip" {
  value = "${alicloud_instance.test.public_ip}"
}

output "disk_id" {
  value = "${alicloud_disk.disk.id}"
}
