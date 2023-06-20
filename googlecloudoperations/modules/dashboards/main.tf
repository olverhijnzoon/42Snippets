resource "google_monitoring_dashboard" "dashboard_simple_demo" {
  dashboard_json = <<EOF
{
  "displayName": "Demo Dashboard",
  "gridLayout": {
    "widgets": [
      {
        "blank": {}
      }
    ]
  }
}

EOF
}

resource "google_monitoring_dashboard" "dashboard_grid_layout_demo" {
  dashboard_json = <<EOF
{
  "dashboardFilters": [],
  "displayName": "Grid Layout Demo",
  "mosaicLayout": {
    "columns": 12,
    "tiles": [
      {
        "height": 4,
        "widget": {
          "title": "Widget 1",
          "xyChart": {
            "chartOptions": {
              "mode": "COLOR"
            },
            "dataSets": [
              {
                "plotType": "LINE",
                "targetAxis": "Y1",
                "timeSeriesQuery": {
                  "timeSeriesQueryLanguage": "fetch gce_instance::compute.googleapis.com/instance/cpu/utilization\n| filter instance_name =~ 'snippets.*'",
                  "unitOverride": ""
                }
              }
            ],
            "thresholds": [],
            "timeshiftDuration": "0s",
            "yAxis": {
              "label": "",
              "scale": "LINEAR"
            }
          }
        },
        "width": 6
      },
      {
        "height": 4,
        "widget": {
          "text": {
            "content": "Widget 2",
            "format": "MARKDOWN",
            "style": {
              "backgroundColor": "",
              "fontSize": "FONT_SIZE_UNSPECIFIED",
              "horizontalAlignment": "H_LEFT",
              "padding": "PADDING_SIZE_UNSPECIFIED",
              "textColor": "",
              "verticalAlignment": "V_TOP"
            }
          }
        },
        "width": 6,
        "xPos": 6
      },
      {
        "height": 4,
        "widget": {
          "title": "Widget 3",
          "xyChart": {
            "chartOptions": {
              "mode": "COLOR"
            },
            "dataSets": [
              {
                "breakdowns": [],
                "dimensions": [],
                "measures": [],
                "plotType": "STACKED_BAR",
                "targetAxis": "Y1",
                "timeSeriesQuery": {
                  "timeSeriesFilter": {
                    "aggregation": {
                      "perSeriesAligner": "ALIGN_RATE"
                    },
                    "filter": "metric.type=\"agent.googleapis.com/nginx/connections/accepted_count\""
                  },
                  "unitOverride": "1"
                }
              }
            ],
            "thresholds": [],
            "timeshiftDuration": "0s",
            "yAxis": {
              "label": "y1Axis",
              "scale": "LINEAR"
            }
          }
        },
        "width": 6,
        "yPos": 4
      }
    ]
  }
}

EOF
}