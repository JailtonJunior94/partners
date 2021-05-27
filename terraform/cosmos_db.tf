resource "azurerm_cosmosdb_account" "go_challenge_account" {
  name                      = "go-challenge-account"
  resource_group_name       = azurerm_resource_group.go_challenge_rg.name
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

resource "azurerm_cosmosdb_mongo_database" "go_challenge_db" {
  name                = "GoChallengeDB"
  resource_group_name = azurerm_cosmosdb_account.go_challenge_account.resource_group_name
  account_name        = azurerm_cosmosdb_account.go_challenge_account.name
  throughput          = 400
}

resource "azurerm_cosmosdb_mongo_collection" "go_challenge_partner" {
  name                = "Partners"
  resource_group_name = azurerm_cosmosdb_account.go_challenge_account.resource_group_name
  account_name        = azurerm_cosmosdb_account.go_challenge_account.name
  database_name       = azurerm_cosmosdb_mongo_database.go_challenge_db.name

  default_ttl_seconds = "777"
  shard_key           = "_id"
  throughput          = 400

  lifecycle {
    ignore_changes = [index]
  }

  depends_on = [azurerm_cosmosdb_mongo_database.go_challenge_db]
}
