resource "azurerm_kubernetes_cluster" "partner_aks" {
  name                = "partner-aks"
  location            = var.location
  resource_group_name = azurerm_resource_group.partners_rg.name
  dns_prefix          = "partner-aks"
  sku_tier            = "Free"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_B2s"
  }

  identity {
    type = "SystemAssigned"
  }
}
