# Slack Exec Terraform Provider: Automate Command Execution and Slack Notifications

## Description

This Terraform provider allows you to execute shell commands and send their output to a Slack channel automatically.


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

  
