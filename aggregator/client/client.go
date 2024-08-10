package client

import (
    "context"

    "github.com/emmanueluwa/tolling/types"
)

type Client interface {
    Aggregate(context.Context, *types.AggregateRequest) error
}
