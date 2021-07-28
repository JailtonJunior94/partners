resource "azurerm_cosmosdb_account" "partners_account" {
  name                      = "partners-account"
  resource_group_name       = azurerm_resource_group.partners_rg.name
  location                  = var.location
  offer_type                = "Standard"
  kind                      = "MongoDB"
  enable_automatic_failover = true
  enable_free_tier          = true

  capabilities {
    name = "EnableMongo"
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 400
    max_staleness_prefix    = 200000
  }

  geo_location {
    location          = var.location
    failover_priority = 0
  }
}

resource "azurerm_cosmosdb_mongo_database" "partners_db" {
  name                = "PartnerDB"
  resource_group_name = azurerm_cosmosdb_account.partners_account.resource_group_name
  account_name        = azurerm_cosmosdb_account.partners_account.name
  throughput          = 400
}

resource "azurerm_cosmosdb_mongo_collection" "partners_partner" {
  name                = "Partners"
  resource_group_name = azurerm_cosmosdb_account.partners_account.resource_group_name
  account_name        = azurerm_cosmosdb_account.partners_account.name
  database_name       = azurerm_cosmosdb_mongo_database.partners_db.name

  default_ttl_seconds = "777"
  shard_key           = "_id"
  throughput          = 400

  lifecycle {
    ignore_changes = [index]
  }

  depends_on = [azurerm_cosmosdb_mongo_database.partners_db]
}
