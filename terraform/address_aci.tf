resource "azurerm_container_group" "api_address_aci" {
  name                = "api-address"
  location            = azurerm_resource_group.go_challenge_rg.location
  resource_group_name = azurerm_resource_group.go_challenge_rg.name
  ip_address_type     = "public"
  dns_name_label      = "api-address"
  os_type             = "linux"

  container {
    name   = "api-address"
    image  = "jailtonjunior/address-api:latest"
    cpu    = 1
    memory = 0.5
    ports {
      port     = 80
      protocol = "TCP"
    }
    ports {
      port     = 9000
      protocol = "TCP"
    }
    environment_variables = tomap({
      "PORT" = "9000"
    })
  }
}
