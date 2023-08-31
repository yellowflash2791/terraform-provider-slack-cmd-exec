package slack_exec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func dataSourceSlackNotify() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackNotifyRead,

		Schema: map[string]*schema.Schema{
			"slack_webhook_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackNotifyRead(d *schema.ResourceData, m interface{}) error {
	// Get the webhook URL and OS command from the schema
	webhookURL := d.Get("slack_webhook_url").(string)
	osCommand := d.Get("os_command").(string)

	// Execute the OS command
	cmd := exec.Command("/bin/bash", "-c", osCommand+" > output.txt")
	if err := cmd.Run(); err != nil {
		return err
	}

	output, err := exec.Command("/bin/bash", "-c", "cat output.txt").Output()
	if err != nil {
		return err
	}

	escapedOutput := fmt.Sprintf("%s", output)
	escapedOutput = fmt.Sprintf("%q", escapedOutput)

	// Prepare the Slack payload
	payload := map[string]interface{}{
		"text": "Output: " + escapedOutput,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Send the Slack notification
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d.SetId("slack_notification_sent")
	d.Set("status", "sent")

	return nil
}
