# Go API client for metasv

API definition for MetaSV provided apis


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/metasv/metasv-go-sdk 
```

## Documentation For Authorization


### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), metasv.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

## Usage Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/metasv/metasv-go-sdk"
	"os"
)

func main() {
	authCtx := context.WithValue(context.Background(), metasv.ContextAccessToken, "YOUR_METASV_TOKEN")
	configuration := metasv.NewConfiguration()
	client := metasv.NewAPIClient(configuration)

	_, response, err := client.BlockApi.BlockGet(authCtx).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BlockInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", response)
	}
	// response from `BlockInfoGet`: BlockchainInfo
	fmt.Fprintf(os.Stdout, "Response from `BlockApi.BlockInfoGet`: %v\n", response)
}

```

## Author

heqiming@metasv.com

