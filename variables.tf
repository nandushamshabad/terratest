# variables.tf
variable "resource_group_name" {
  description = "Existing resource group name"
  type        = string
}

variable "storage_account_name" {
  description = "Storage account name"
  type        = string
}

variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
  sensitive   = true
}

variable "subscription_name" {
  description = "Azure subscription name"
  type        = string
}

variable "location" {
  description = "Azure region"
  type        = string
  default     = "eastus"
}
variable "kind" {
  description = "Storage account kind"
  type        = string
}

variable "replication_type" {
  description = "Replication type (LRS, GRS, etc.)"
  type        = string
}

variable "access_tier" {
  description = "Access tier for the storage account. Defaults to 'Hot'"
  type        = string
  default     = "Hot"
}

variable "min_tls_version" {
  description = "Minimum TLS version for the storage account. Defaults to 'TLS1_2'"
  type        = string
  default     = "TLS1_2"
}

variable "https_traffic_only" {
  description = "Enforce HTTPS traffic only. Defaults to true"
  type        = bool
  default     = true
}

variable "allow_blob_public_access" {
  description = "Allow public access to blobs. Defaults to false"
  type        = bool
  default     = false
}

variable "enable_access_key_output" {
  description = "Enable output of primary access key (sensitive)"
  type        = bool
  default     = true
}

# ST-004: Private endpoint configuration
variable "vnet_name" {
  description = "Virtual network name"
  type        = string
}

variable "subnet_name" {
  description = "Subnet name for private endpoint"
  type        = string
}

variable "private_endpoint_name" {
  description = "Private endpoint name"
  type        = string
  default     = "pe-storage-blob"
}

variable "private_dns_zone_name" {
  description = "Private DNS zone name for blob"
  type        = string
  default     = "privatelink.blob.core.windows.net"
}

variable "existing_private_dns_zone" {
  description = "Use existing private DNS zone or create new"
  type        = bool
  default     = false
}

variable "subnet_id" {
  description = "Existing subnet ID for private endpoint"
  type        = string
  default     = ""
}