# terraform.tfvars# Replace these values with your ownresource_group_name  = "Test"     # Change to your RG namestorage_account_name = "nandustoragetest"          # Change to unique namesubscription_id      = "9f0d6fb2-38cc-4057-a6c9-81d97f97d6a9"  # Your sub IDsubscription_name    = "Azure subscription 1"  # Your sub namelocation            = "ukwest"
# terraform.tfvars
resource_group_name  = "test"
storage_account_name = "sttest001nandu"
subscription_id      = "user your value"
subscription_name    = "Azure subscription 1"
location            = "ukwest"
replication_type = "LRS"
kind = "StorageV2"

# Default values for UT-002 - these are the defaults but explicitly defined
access_tier = "Hot"
min_tls_version          = "TLS1_2"
https_traffic_only       = true
allow_blob_public_access = false

# For UT-003
enable_access_key_output = true

# ST-004: Private endpoint values
vnet_name                     = "vnet-test"
subnet_name                   = "snet-private"
private_endpoint_name         = "pe-sttest001"
private_dns_zone_name         = "privatelink.blob.core.windows.net"
existing_private_dns_zone     = false
subnet_id                     = ""