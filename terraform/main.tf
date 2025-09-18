terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "7.3.0"
    }
  }
}

provider "google" {
  project     = var.project_id
  region      = var.region
}

resource "google_project_service" "enabled_apis" {
  for_each = toset([
    "container.googleapis.com",
    "sqladmin.googleapis.com",
    "artifactregistry.googleapis.com"
  ])
  project = var.project_id
  service = each.key
}

module "database" {
    source = "./database"
    instance_name = "postgres-14"
    database_version = "PPOSTGRES_15"
    region = var.region
}

module "gke" {
  source = "./gke"
  region = var.region
  machine_type = var.gke_machine_type

}