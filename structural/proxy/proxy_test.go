package proxy

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		nginxServer := newNginxServer()
		appStatusURL := "/app/status"
		createuserURL := "/create/user"

		httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	})
}
