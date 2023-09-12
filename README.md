# fuse-go

## Quick start
Documentation for each method, request param, and response field are available in docstrings and will appear on hover in most modern editors.

```
package fuse_go

import (
	"context"
	"fuse"
	"os"
	"testing"
)

var (
	// global client instance
	client *fuse.APIClient
)

// Initialize the global client
func init() {
	cfg := fuse.NewConfiguration()
	cfg.Servers[0].URL = os.Getenv("BASE_URL")
	cfg.AddDefaultHeader("Fuse-Client-Id", "my-fuse-client-id")
	cfg.AddDefaultHeader("Fuse-Api-Key", "my-fuse-api-key")
	cfg.AddDefaultHeader("Plaid-Client-Id", "my-plaid-client-id")
	cfg.AddDefaultHeader("Plaid-Secret", "my-plaid-secret")
	client = fuse.NewAPIClient(cfg)
}

func Run() {
	aggregators := []fuse.Aggregator{
		fuse.AGGREGATOR_PLAID,
	}
	products := []fuse.Product{
		fuse.PRODUCT_ACCOUNT_DETAILS,
		fuse.PRODUCT_BALANCE,
	}

	create_session_request := *fuse.NewCreateSessionRequest(aggregators, products, *fuse.NewEntity("entity-id"))
	create_session_response, _, err := client.FuseAPI.CreateSessionExecute(client.FuseAPI.CreateSession(context.Background()).CreateSessionRequest(create_session_request))

	create_link_token_request := fuse.NewCreateLinkTokenRequest("institution-id", *fuse.NewEntity("entity-id"), "Fuse", create_session_response.ClientSecret)
	create_link_token_response, _, err := client.FuseAPI.CreateLinkTokenExecute(client.FuseAPI.CreateLinkToken(context.Background()).CreateLinkTokenRequest(*create_link_token_request))
}
```


Publishing
```
go mod tidy
```
