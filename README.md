Overview
The Slack Exec Terraform Provider allows you to execute shell commands and send their outputs to a designated Slack channel. This enables you to automate the monitoring and notification process within your infrastructure as code (IAC) projects. Whether you want to get server statistics, application statuses, or any other information, this provider integrates these functionalities directly into your Terraform workflow.

Prerequisites
Go 1.16+
Terraform 0.13+
Slack Workspace with the ability to create webhooks

Installation

Clone the Repository

git clone https://github.com/your-username/slack-exec-terraform-provider.git

Build the Provider

cd slack-exec-terraform-provider

go mod init terraform-provider-slack-cmd-exec

go mod tidy 

make install


Now implement the 
