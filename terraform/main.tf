terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.56.0"
    }
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
