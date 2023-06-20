terraform {
  required_version = "~> 1.3.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.40"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "google_storage_bucket" "terraform_state" {
  name     = "42snippets-state-bucket"
  location = "EU"
}

resource "google_compute_instance" "default" {
  name         = "snippets42"
  machine_type = var.machine_type
  tags = ["http-server"]

  boot_disk {
    initialize_params {
      image = "debian-11"
    }
  }

  network_interface {
    network = "default"

    # Enable external IP
    access_config {}
  }

  metadata_startup_script = <<SCRIPT
  apt-get update
  apt-get install -y nginx
  systemctl start nginx
  SCRIPT
}

resource "google_compute_firewall" "default" {
  name    = "nginx-allow-http"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["http-server"]
  source_ranges = ["0.0.0.0/0"]
}

resource "google_monitoring_uptime_check_config" "default" {
  display_name = "uptime-check"
  timeout       = "10s"
  period        = "60s"

  http_check {
    path = "/"
    port = "80"
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "snippets-390309"
      # Reference the external IP of the compute instance
      host       = google_compute_instance.default.network_interface[0].access_config[0].nat_ip
    }
  }

  # Explicitly specify dependency on the Compute Instance
  depends_on = [google_compute_instance.default]
}

module "alerts" {
  source = "./modules/alerts"
  email_address = var.email_address
}

module "dashboards" {
  source = "./modules/dashboards"
}