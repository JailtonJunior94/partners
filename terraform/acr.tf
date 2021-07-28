resource "azurerm_container_registry" "partner-acr" {
  name                = "partnersacr"
  resource_group_name = azurerm_resource_group.partners_rg.name
  location            = var.location
  sku                 = "Basic"
}
