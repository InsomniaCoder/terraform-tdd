provider "azurerm" {
  version = "=2.32.0"
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "test-rg"
  location = "Southeast Asia"
}

resource "azurerm_container_registry" "example" {
  name                = "unitTestACR"
  resource_group_name = "${azurerm_resource_group.example.name}"
  location            = "${azurerm_resource_group.example.location}"
  sku                 = "Standard"
}
