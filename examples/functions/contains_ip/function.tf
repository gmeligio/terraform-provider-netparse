locals {
  example1 = provider::netparse::contains_ip("192.0.2.0/24", "192.0.2.1") # true

  example2 = provider::netparse::contains_ip("192.0.2.0/24", "192.1.0.0") # false
}
