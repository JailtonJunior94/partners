output "client_certificate" {
  value = azurerm_kubernetes_cluster.partner_aks.kube_config.0.client_certificate
}

output "kube_config" {
  value = nonsensitive(azurerm_kubernetes_cluster.partner_aks.kube_config_raw)
}
