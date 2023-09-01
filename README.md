# Slack Exec Terraform Provider: Automate Command Execution and Slack Notifications

## Description

This Terraform provider allows you to execute shell commands and send their output to a Slack channel automatically.

## 🚨 Disclaimer 🚨

This project is intended for educational purposes only. The authors and contributors are not responsible for any malicious use or damages that may be done using this tool. Use responsibly.

## Prerequisites
Go 1.16+

Terraform 0.13+

Slack Workspace with the ability to create webhooks

## Installation

### Clone the Repository

    git clone https://github.com/your-username/slack-exec-terraform-provider.git

### Build the Provider

      cd slack-exec-terraform-provider

      go mod init terraform-provider-slack-cmd-exec

      go mod tidy 

      make install


## Usage 

### Refer examples/main.tf and examples/data.tf

    terraform {  
        required_providers {
            slack-cmd-exec = {
              source = "github.com/terraform-provider-slack-cmd-exec/slack-cmd-exec"
              version = "0.2.0" # use the version that you need
            }
        }
     }
            
    
### Input the slack hook and command to execute in the data.tf

    data "slack_cmd_exec" "example" {
      provider = slack-cmd-exec

### Now run the below terraform commands

    terraform init
    terraform plan
      slack_webhook_url = "<slack_hook>"
      os_command        = "uname -a"
    }

