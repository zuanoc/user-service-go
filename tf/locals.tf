locals {
  env = terraform.workspace
  service_name = "user-service"
  region = "eu-north-1"
  account_id = "781831274713"
  common_tags = {
    CREATED_BY: "Terraform"
    SERVICE: local.service_name
  }
}
