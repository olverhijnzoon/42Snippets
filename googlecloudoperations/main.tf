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

  boot_disk {
    initialize_params {
      image = "debian-11-bullseye"
    }
  }

  network_interface {
    network = "default"

    # Enable external IP
    access_config {}
  }
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

resource "google_monitoring_notification_channel" "default" {
  display_name = "notification-channel"
  type         = "email"
  labels = {
    email_address = var.email_address
  }
}

resource "google_monitoring_alert_policy" "default" {
  display_name = "alert-policy"
  combiner     = "OR"

  conditions {
    display_name = "uptime-check"
    condition_threshold {
      filter     = "metric.type=\"monitoring.googleapis.com/uptime_check/check_passed\" resource.type=\"uptime_url\""
      comparison = "COMPARISON_LT"
      threshold_value = 1
      duration   = "60s"

      aggregations {
        alignment_period   = "60s"
        per_series_aligner = "ALIGN_NEXT_OLDER"
      }
    }
  }

  notification_channels = [
    google_monitoring_notification_channel.default.name
  ]
}