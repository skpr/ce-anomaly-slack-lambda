package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/skpr/slack"
)

const (
	EnvSlackWebhook   = "SKPR_CE_ANOMALY_LAMBDA_SLACK_WEBHOOK"
	EnvSlackIcon      = "SKPR_CE_ANOMALY_LAMBDA_SLACK_ICON"
	EnvSlackDashboard = "SKPR_CE_ANOMALY_LAMBDA_SLACK_DASHBOARD"
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

// HandleLambdaEvent will respond to a CloudWatch Alarm, check for rate limited IP addresses and send a message to Slack.
func HandleLambdaEvent(ctx context.Context, event *Event) error {
	client, err := slack.NewClient([]string{
		os.Getenv(EnvSlackWebhook),
	})
	if err != nil {
		return fmt.Errorf("failed to setup slack client: %w", err)
	}

	dashboard := os.Getenv(EnvSlackDashboard)

	// Use the anomaly detection link as a fallback for the dashboard link.
	if dashboard == "" {
		dashboard = event.AnomalyDetailsLink
	}

	params := slack.PostMessageParams{
		Context: map[string]string{
			"Account":    event.AccountID,
			"Anomaly ID": event.AnomalyId,
		},
		Description: "Cost anomaly has been detected!",
		Reason:      fmt.Sprintf("%s has increased by %f%%", event.DimensionalValue, event.Impact.TotalImpactPercentage),
		Dashboard:   dashboard,
		Icon:        os.Getenv(EnvSlackIcon),
	}

	return client.PostMessage(params)
}
