output "domain" {
  value = provider::netparse::parse_cidr("192.0.2.1/24")

  # {
  #   ip      = "192.0.2.1"
  #   network = "192.0.2.0/24"
  # }
}
