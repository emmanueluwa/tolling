package client

import (
    "net/http"
    "context"
    "encoding/json"
    "bytes"
    "fmt"

    "github.com/emmanueluwa/tolling/types"

)

type HTTPClient struct {
    Endpoint string
}


func NewHTTPClient(endpoint string) *HTTPClient {
    return &HTTPClient{
        Endpoint: endpoint,
    }
}


func (c *HTTPClient) Aggregate(ctx context.Context, aggReq *types.AggregateRequest) error {
    b, err := json.Marshal(aggReq)
    if err != nil {
        return err
    }

    //prepares request
    req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
    if err != nil {
        return err
    }

    //send req
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("the service responded with non 200 status code %d", resp.StatusCode)
    }
    return nil
}
