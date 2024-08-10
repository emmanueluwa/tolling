package client

import (
    "net/http"
    "encoding/json"
    "bytes"
    "fmt"

    "github.com/emmanueluwa/tolling/types"

)

type Client struct {
    Endpoint string
}


func NewClient(endpoint string) *Client {
    return &Client{
        Endpoint: endpoint,
    }
}


func (c *Client) AggregateInvoice(distance types.Distance) error {
    b, err := json.Marshal(distance)
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
