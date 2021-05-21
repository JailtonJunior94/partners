output "api_address_ip_address" {
  value = "${azurerm_container_group.api_address_aci.ip_address}"
}

output "api_address_fqdn" {
  value = "${azurerm_container_group.api_address_aci.fqdn}"
}

output "api_partner_ip_address" {
  value = "${azurerm_container_group.api_partner_aci.ip_address}"
}

output "api_partner_fqdn" {
  value = "${azurerm_container_group.api_partner_aci.fqdn}"
}