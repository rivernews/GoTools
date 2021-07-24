package GoTools

func GetSlackWebhookURL() string {
	url := getEnvVarOrDefault("SLACK_WEBHOOK_URL", "")
	if url == "" {
		Logger("ERROR", "Missing environment variable `SLACK_WEBHOOK_URL`")
	}
	return url
}

func SendSlackMessage(message string) {
	Fetch(FetchOption{
		URL: GetSlackWebhookURL(),
		Method: "POST",
		PostData: map[string]string{
			"text": message,
			// "channel": "#build",
			// "username": "Kubernetes Cluster Head Service",
			// "icon_url": "https://github.com/kubernetes/kubernetes/raw/master/logo/logo.png",
		},
	})
}