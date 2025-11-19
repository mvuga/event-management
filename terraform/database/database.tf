variable "name" {
    description = "Database name"
    default = "event_management"
}

variable "instance_name" {
    default = "postgres-15"
}

variable "region" {}

variable "database_version" {
    default = "POSTGRES_15"
}

variable "tier" {
    description = "Database instance type"
    default = "db-g1-small"
}

variable "database_user" {
    description = "Database username"
    default = "eventusr"
}

variable "database_deletion_policy" {
    default = "ABANDON" # DEFAULT for POSTGRES databases
}

resource "random_password" "database_password" {
  length           = 20
  special          = false
}


resource "google_sql_database" "database" {
  name     = var.name
  instance = google_sql_database_instance.main.name
  deletion_policy = var.database_deletion_policy
  # charset = "en_US.UTF8"
}

resource "google_sql_database_instance" "main" {
  name             = var.instance_name
  region           = var.region
  database_version = var.database_version
  settings {
    tier = var.tier
  }

  deletion_protection  = true
}

resource "google_sql_user" "users" {
  name     = var.database_user
  instance = google_sql_database_instance.main.name
  password = random_password.database_password.result
}

output "database_password" {

    value = google_sql_user.users.password
    sensitive = true
  
}

output "database_dns_name" {

    value = google_sql_database_instance.main.connection_name
  
}