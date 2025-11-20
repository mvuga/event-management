variable "project_id" {
    description = "GCP Project ID"
    default = "gcloud-showcase"
    type = string
}

variable "region" {
    description = "GCP deployment region"
    default = "europe-west3-c"
    type= string
}

variable "gke_machine_type" {
    description = "GKE machine type"
    default = "e2-small"
}