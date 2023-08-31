data "slack_cmd_exec" "example" {
  provider = slack-cmd-exec
  
  slack_webhook_url = "<slack_hook>"
  os_command        = "uname -a"
}

