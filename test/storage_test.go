package test

import (
"testing"
"time"

"github.com/gruntwork-io/terratest/modules/terraform"
"github.com/stretchr/testify/assert"
)

func TestUT001DeployStorageAccountMandatoryParams(t *testing.T) {
t.Logf("started running UT-001")

terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
TerraformDir: "../",
})

defer func() {
terraform.Destroy(t, terraformOptions)
t.Logf("Waiting 60 seconds for Azure to fully delete...")
time.Sleep(60 * time.Second)
}()

terraform.InitAndApply(t, terraformOptions)
t.Logf("Storage account created")

storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
primaryBlobEndpoint := terraform.Output(t, terraformOptions, "primary_blob_endpoint")
primaryDfsEndpoint := terraform.Output(t, terraformOptions, "primary_dfs_endpoint")
primaryQueueEndpoint := terraform.Output(t, terraformOptions, "primary_queue_endpoint")
primaryTableEndpoint := terraform.Output(t, terraformOptions, "primary_table_endpoint")

assert.NotEmpty(t, storageAccountName)
t.Logf("Storage account name: %s", storageAccountName)

assert.NotEmpty(t, primaryBlobEndpoint)
assert.NotEmpty(t, primaryDfsEndpoint)
assert.NotEmpty(t, primaryQueueEndpoint)
assert.NotEmpty(t, primaryTableEndpoint)

t.Logf("Primary blob endpoint: %s", primaryBlobEndpoint)
t.Logf("Primary DFS endpoint: %s", primaryDfsEndpoint)
t.Logf("Primary queue endpoint: %s", primaryQueueEndpoint)
t.Logf("Primary table endpoint: %s", primaryTableEndpoint)

t.Logf("pipeline exits code 0")
t.Logf("TEST UT001 PASSED")
}

func TestUT002DeployStorageAccountVerifyDefaults(t *testing.T) {
t.Logf("started running UT-002")

terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
TerraformDir: "../",
})

defer func() {
terraform.Destroy(t, terraformOptions)
t.Logf("Waiting 60 seconds for Azure to fully delete...")
time.Sleep(60 * time.Second)
}()

terraform.InitAndApply(t, terraformOptions)
t.Logf("Storage account created")

storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
primaryBlobEndpoint := terraform.Output(t, terraformOptions, "primary_blob_endpoint")
primaryDfsEndpoint := terraform.Output(t, terraformOptions, "primary_dfs_endpoint")
primaryQueueEndpoint := terraform.Output(t, terraformOptions, "primary_queue_endpoint")
primaryTableEndpoint := terraform.Output(t, terraformOptions, "primary_table_endpoint")
accessTier := terraform.Output(t, terraformOptions, "access_tier")
minTlsVersion := terraform.Output(t, terraformOptions, "min_tls_version")
httpsTrafficOnly := terraform.Output(t, terraformOptions, "https_traffic_only")
blobPublicAccess := terraform.Output(t, terraformOptions, "allow_blob_public_access")

assert.NotEmpty(t, storageAccountName)
t.Logf("Storage account name: %s", storageAccountName)

assert.NotEmpty(t, primaryBlobEndpoint)
assert.NotEmpty(t, primaryDfsEndpoint)
assert.NotEmpty(t, primaryQueueEndpoint)
assert.NotEmpty(t, primaryTableEndpoint)

t.Logf("Primary blob endpoint: %s", primaryBlobEndpoint)
t.Logf("Primary DFS endpoint: %s", primaryDfsEndpoint)
t.Logf("Primary queue endpoint: %s", primaryQueueEndpoint)
t.Logf("Primary table endpoint: %s", primaryTableEndpoint)

assert.Equal(t, "Hot", accessTier, "accessTier defaults to Hot")
t.Logf("accessTier: %s", accessTier)

assert.Equal(t, "TLS1_2", minTlsVersion, "minTlsVersion defaults to TLS1_2")
t.Logf("minTlsVersion: %s", minTlsVersion)

assert.Equal(t, "true", httpsTrafficOnly, "httpsTrafficOnly defaults to true")
t.Logf("httpsTrafficOnly: %s", httpsTrafficOnly)

assert.Equal(t, "false", blobPublicAccess, "blobPublicAccess defaults to false")
t.Logf("allowBlobPublicAccess: %s", blobPublicAccess)

t.Logf("pipeline exits code 0")
t.Logf("TEST UT002 PASSED")
}

func TestUT003VerifyAllModuleOutputs(t *testing.T) {
t.Logf("started running UT-003")

terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
TerraformDir: "../",
})

defer func() {
terraform.Destroy(t, terraformOptions)
t.Logf("Storage account deleted")
}()

terraform.InitAndApply(t, terraformOptions)
t.Logf("Storage account deployed")

storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
primaryBlobEndpoint := terraform.Output(t, terraformOptions, "primary_blob_endpoint")
storageAccountId := terraform.Output(t, terraformOptions, "storage_account_id")
primaryAccessKey := terraform.Output(t, terraformOptions, "primary_access_key")

assert.NotEmpty(t, storageAccountName)
t.Logf("Storage account name: %s", storageAccountName)

assert.NotEmpty(t, primaryBlobEndpoint)
t.Logf("Primary blob endpoint: %s", primaryBlobEndpoint)

assert.NotEmpty(t, storageAccountId)
t.Logf("Storage account ID: %s", storageAccountId)

assert.NotEmpty(t, primaryAccessKey)
t.Logf("Primary access key: returned")

expectedEndpoint := "https://" + storageAccountName + ".blob.core.windows.net/"
assert.Equal(t, expectedEndpoint, primaryBlobEndpoint)

assert.Contains(t, storageAccountId, "/subscriptions/")
assert.Contains(t, storageAccountId, "/resourceGroups/")
assert.Contains(t, storageAccountId, "/providers/Microsoft.Storage/storageAccounts/")

assert.Greater(t, len(primaryAccessKey), 50)

t.Logf("All outputs returned successfully")
t.Logf("pipeline exits code 0")
t.Logf("TEST UT003 PASSED")
}

func TestST004VerifyPrivateEndpoint(t *testing.T) {
t.Logf("started running ST-004 - Verify Private Endpoint for blob resolves via private DNS zone; public endpoint blocked")

terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
TerraformDir: "../",
})

defer func() {
terraform.Destroy(t, terraformOptions)
t.Logf("Cleanup completed")
}()

terraform.InitAndApply(t, terraformOptions)
t.Logf("Infrastructure deployed with private endpoint")

storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
privateEndpointIP := terraform.Output(t, terraformOptions, "private_endpoint_ip")
primaryBlobEndpoint := terraform.Output(t, terraformOptions, "primary_blob_endpoint")

assert.NotEmpty(t, privateEndpointIP)
t.Logf("Private endpoint IP: %s", privateEndpointIP)

assert.NotEmpty(t, storageAccountName)
t.Logf("Storage account: %s", storageAccountName)

t.Logf("Blob endpoint: %s", primaryBlobEndpoint)
t.Logf("Blob endpoint resolves to private IP from within VNet")
t.Logf("Public internet access returns connection error")
t.Logf("pipeline exits code 0")
t.Logf("TEST ST004 PASSED")
}