
resource "aws_vpc" "default" {
  count = var.vpc_enabled ? 1 : 0

  cidr_block                       = var.cidr_block
  instance_tenancy                 = var.instance_tenancy
  enable_dns_hostnames             = var.enable_dns_hostnames
  enable_dns_support               = var.enable_dns_support
  enable_classiclink               = var.enable_classiclink
  enable_classiclink_dns_support   = var.enable_classiclink_dns_support
  assign_generated_ipv6_cidr_block = true
  tags = {
    "Name" = "dhyanio-infra"
  }
  lifecycle {
    ignore_changes = [
      tags,
      tags["kubernetes.io"],
      tags["SubnetType"],
    ]
  }
}

resource "aws_internet_gateway" "default" {
  count = var.vpc_enabled ? 1 : 0

  vpc_id = element(aws_vpc.default.*.id, count.index)
  tags = {
    "Name" = "dhyanio-infra"
  }
}


resource "aws_flow_log" "vpc_flow_log" {
  count = var.vpc_enabled && var.enable_flow_log == true ? 1 : 0

  log_destination      = var.s3_bucket_arn
  log_destination_type = "s3"
  traffic_type         = var.traffic_type
  vpc_id               = element(aws_vpc.default.*.id, count.index)
}