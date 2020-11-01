output "rg_name" {
  value = azurerm_resource_group.example.name
}
output "acr_name" {
  value = azurerm_container_registry.example.name
}
output "acr_location" {
  value = azurerm_container_registry.example.location
}
output "acr_sku" {
  value = azurerm_container_registry.example.sku
}

