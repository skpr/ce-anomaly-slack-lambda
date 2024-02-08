Cost Explorer Anomaly Detection Slack Lambda
============================================

A Lambda function for posting Cost Explorer anomalies to Slack

## Environment Variables

* **SKPR_CE_ANOMALY_LAMBDA_SLACK_WEBHOOK** (Required) - Configure the webhook which will be used to post the message.
* **SKPR_CE_ANOMALY_LAMBDA_SLACK_ICON** (Optional) - Configure the icon which will be used for the post.
* **SKPR_CE_ANOMALY_LAMBDA_SLACK_DASHBOARD** (Optional) - Configure the dashboard link. Will default to the anomaly details link.
