variable "name" {
  type    = string
  default = ""
}

variable "application" {
  type    = string
  default = ""
}

variable "environment" {
  type    = string
  default = ""
}


variable "label_order" {
  type    = list
  default = []
}



variable "attributes" {
  type    = list
  default = []
}

variable "tags" {
  type    = map
  default = {}
}


variable "vpc_enabled" {
  type    = bool
  default = true
}

variable "cidr_block" {
  type    = string
  default = ""
}

variable "instance_tenancy" {
  type    = string
  default = "default"

}

variable "enable_dns_hostnames" {
  type    = bool
  default = true
}

variable "enable_dns_support" {
  type    = bool
  default = true
}

variable "enable_classiclink" {
  type    = bool
  default = false
}

variable "enable_classiclink_dns_support" {
  type    = bool
  default = false
}

variable "enable_flow_log" {
  type    = bool
  default = false
}

variable "s3_bucket_arn" {
  type    = string
  default = ""

}

variable "traffic_type" {
  type    = string
  default = "ALL"
}
