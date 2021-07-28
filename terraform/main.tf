terraform {
  backend "azurerm" {
    resource_group_name  = "jj-rg"
    storage_account_name = "storagepartners"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscriptionid
}

resource "azurerm_resource_group" "partners_rg" {
  name     = "partners-rg"
  location = var.location
}
