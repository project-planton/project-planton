variable "gcp_credential" {
  description = "GCP Credential data"
  type = object({
    # The Google Service Account Base64 Encoded Key, which is used to authenticate API requests to GCP services.
    # This is a required field, and the value must be a valid base64 encoded string.
    service_account_key_base64 = string
  })
}

provider "google" {
}
