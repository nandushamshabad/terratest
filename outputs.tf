# outputs.tf
output "storage_account_name" {
  value = azurerm_storage_account.main.name
}

output "resource_group_name" {
  value = var.resource_group_name
}

output "subscription_id_used" {
  value = var.subscription_id
  sensitive = true
}

output "subscription_name_used" {
  value = var.subscription_name
}
output "primary_blob_endpoint" {
  value = azurerm_storage_account.main.primary_blob_endpoint
}

output "primary_dfs_endpoint" {
  value = azurerm_storage_account.main.primary_dfs_endpoint
}

output "primary_queue_endpoint" {
  value = azurerm_storage_account.main.primary_queue_endpoint
}

output "primary_table_endpoint" {
  value = azurerm_storage_account.main.primary_table_endpoint
}

output "access_tier" {
  value = azurerm_storage_account.main.access_tier
}

output "min_tls_version" {
  value = azurerm_storage_account.main.min_tls_version
}

output "https_traffic_only" {
  value = azurerm_storage_account.main.https_traffic_only_enabled
}

output "allow_blob_public_access" {
  value = lookup(azurerm_storage_account.main, "allow_blob_public_access", false)
}

# New outputs for UT-003
output "storage_account_id" {
  description = "ARM resource ID of the storage account"
  value = azurerm_storage_account.main.id
}

output "primary_access_key" {
  description = "Primary access key for the storage account"
  value = var.enable_access_key_output ? azurerm_storage_account.main.primary_access_key : null
  sensitive = true
}

# ST-004: Private endpoint outputs
output "private_endpoint_ip" {
  value = azurerm_private_endpoint.main.private_service_connection[0].private_ip_address
}

output "private_dns_zone_id" {
  value = azurerm_private_dns_zone.main.id
}

output "private_endpoint_id" {
  value = azurerm_private_endpoint.main.id
}

output "private_endpoint_fully_qualified_domain_name" {
  value = azurerm_private_endpoint.main.private_service_connection[0].private_ip_address
}