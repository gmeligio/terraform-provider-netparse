data "netparse_cidr" "example1" {
  cidr = "192.0.2.1/24"
}

output "cidr1" {
  value = data.netparse_cidr.example

  # {
  #   cidr    = "192.0.2.1/24"
  #   ip      = "192.0.2.1"
  #   network = "192.0.2.0/24"
  # }
}

data "netparse_cidr" "example2" {
  cidr = "192.0.2.1/24"
}

output "cidr2" {
  value = data.netparse_cidr.example

  # {
  #   cidr    = "2001:db8:a0b:12f0::1/32"
  #   ip      = "2001:db8:a0b:12f0::1"
  #   network = "2001:db8::/32"
  # }
}
