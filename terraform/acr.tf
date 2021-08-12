resource "azurerm_container_registry" "partner-acr" {
  name                = "partnersregistry"
  resource_group_name = azurerm_resource_group.partners_rg.name
  location            = var.location
  sku                 = "Basic"
  admin_enabled       = true
}
