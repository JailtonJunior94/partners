terraform {
  backend "azurerm" {
    resource_group_name  = "bookings-rg"
    storage_account_name = "bookingstoragejj"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscriptionid
}

resource "azurerm_resource_group" "go_challenge_rg" {
  name     = "go-challenge-rg"
  location = var.location
}
