resource "azurerm_container_group" "api_partner_aci" {
  name                = "api-partner"
  location            = azurerm_resource_group.go_challenge_rg.location
  resource_group_name = azurerm_resource_group.go_challenge_rg.name
  ip_address_type     = "public"
  dns_name_label      = "api-partner"
  os_type             = "linux"

  container {
    name   = "api-partner"
    image  = "jailtonjunior/partner-api:latest"
    cpu    = "0.5"
    memory = "0.5"
    ports {
      port     = 80
      protocol = "TCP"
    }
    ports {
      port     = 8000
      protocol = "TCP"
    }
    environment_variables = tomap({
      "PORT" = "8000"
    })
  }
}
