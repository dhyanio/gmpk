
output "vpc_id" {
  value = join("", aws_vpc.default.*.id)
}

output "vpc_cidr_block" {
  value = join("", aws_vpc.default.*.cidr_block)
}

output "vpc_main_route_table_id" {
  value = join("", aws_vpc.default.*.main_route_table_id)
}

output "vpc_default_network_acl_id" {
  value = join("", aws_vpc.default.*.default_network_acl_id)
}

output "vpc_default_security_group_id" {
  value = join("", aws_vpc.default.*.default_security_group_id)
}

output "vpc_default_route_table_id" {
  value = join("", aws_vpc.default.*.default_route_table_id)
}

output "vpc_ipv6_association_id" {
  value = join("", aws_vpc.default.*.ipv6_association_id)
}

output "ipv6_cidr_block" {
  value = join("", aws_vpc.default.*.ipv6_cidr_block)
}


output "igw_id" {
  value = join("", aws_internet_gateway.default.*.id)
}