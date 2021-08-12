output "acr_url" {
  value = azurerm_container_registry.partner-acr.login_server
}

output "acr_username" {
  value = azurerm_container_registry.partner-acr.admin_username
}

output "acr_password" {
  value = nonsensitive(azurerm_container_registry.partner-acr.admin_password)
}

output "cosmosdb_connectionstrings" {
  value = nonsensitive(azurerm_cosmosdb_account.partners_account.connection_strings)
}
