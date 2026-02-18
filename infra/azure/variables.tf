variable "location" {
  description = "The Azure Region in which all resources in this example should be created."
  default     = "East US"
}

variable "client_id" {
  description = "The Client ID (appId) for the Service Principal."
}

variable "client_secret" {
  description = "The Client Secret (password) for the Service Principal."
}
