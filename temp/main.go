package main

import (
    "log"
    "time"
    "context"

    "github.com/emmanueluwa/tolling/types"
    "github.com/emmanueluwa/tolling/aggregator/client"
)

func main() {
    c, err := client.NewGRPCClient(":3001")
    if err != nil {
        log.Fatal(err)
    }

    if _, err := c.Aggregate(context.Background(), &types.AggregateRequest{
        ObuID: 1,
        Value: 101.11,
        Unix: time.Now().UnixNano(),
    }); err != nil {
        log.Fatal(err)
    }

}


//if _, err := c.Aggregate(context.Background(), &types.AggregateRequest{
//        ObuID: 1,
//        Value: 101.11,
//        Unix: time.Now().UnixNano(),
//    }); err != nil {
//        log.Fatal(err)
//    }

