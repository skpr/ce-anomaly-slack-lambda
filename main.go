package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/skpr/slack"
)

const (
	// EnvSlackWebhook is used to define the Slack webhook URL.
	EnvSlackWebhook = "SKPR_CE_ANOMALY_LAMBDA_SLACK_WEBHOOK"
	// EnvSlackIcon is used to define the Slack icon.
	EnvSlackIcon = "SKPR_CE_ANOMALY_LAMBDA_SLACK_ICON"
	// EnvSlackIconDefault is used to define the default Slack icon.
	EnvSlackIconDefault = "https://raw.githubusercontent.com/skpr/slack/main/icons/aws_cost_explorer.png"
	// EnvSlackDashboard is used to define the Slack dashboard URL.
	EnvSlackDashboard = "SKPR_CE_ANOMALY_LAMBDA_SLACK_DASHBOARD"
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

// HandleLambdaEvent will respond to a Cost Explorer anomaly and post it to Slack.
func HandleLambdaEvent(ctx context.Context, event *Event) error {
	client, err := slack.NewClient([]string{
		os.Getenv(EnvSlackWebhook),
	})
	if err != nil {
		return fmt.Errorf("failed to setup slack client: %w", err)
	}

	var (
		dashboard = os.Getenv(EnvSlackDashboard)
		icon      = os.Getenv(EnvSlackIcon)
	)

	// Use the anomaly detection link as a fallback for the dashboard link.
	if dashboard == "" {
		dashboard = event.AnomalyDetailsLink
	}

	// Use the default icon if none is provided.
	if icon == "" {
		icon = EnvSlackIconDefault
	}

	params := slack.PostMessageParams{
		Context: map[string]string{
			"Account":    event.AccountID,
			"Anomaly ID": event.AnomalyID,
		},
		Description: "Cost anomaly has been detected!",
		Reason:      fmt.Sprintf("%s has increased by %f%%", event.DimensionalValue, event.Impact.TotalImpactPercentage),
		Dashboard:   dashboard,
		Icon:        icon,
	}

	return client.PostMessage(params)
}
