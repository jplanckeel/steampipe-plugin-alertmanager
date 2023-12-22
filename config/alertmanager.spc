connection "steampipe-plugin-alertmanager" {
  plugin  = "jplanckeel/alertmanager"

  # address to access the API (required).
  address = "alertmanager.localhost.com:8080"

  # API token for your opsgenie instance (required).
  # See https://docs.opsgenie.com/docs/api-access-management
  schemes = "https"

  # To filter request you can add opsgenie query
  #path = "/"
}
