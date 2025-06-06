package clients

import (
	"context"

	"github.com/cryptellation/candlesticks/api"
	temporalclient "go.temporal.io/sdk/client"
)

// Client is a client for the cryptellation candlesticks service.
type Client interface {
	// ListCandlesticks calls the candlesticks list workflow.
	ListCandlesticks(
		ctx context.Context,
		params api.ListCandlesticksWorkflowParams,
	) (res api.ListCandlesticksWorkflowResults, err error)
	// Info calls the service info.
	Info(ctx context.Context) (api.ServiceInfoResults, error)
}

type client struct {
	temporal temporalclient.Client
}

// New creates a new client to execute temporal workflows.
func New(cl temporalclient.Client) Client {
	return &client{temporal: cl}
}

// ListCandlesticks calls the candlesticks list workflow.
func (c client) ListCandlesticks(
	ctx context.Context,
	params api.ListCandlesticksWorkflowParams,
) (res api.ListCandlesticksWorkflowResults, err error) {
	workflowOptions := temporalclient.StartWorkflowOptions{
		TaskQueue: api.WorkerTaskQueueName,
	}

	// Execute workflow
	exec, err := c.temporal.ExecuteWorkflow(ctx, workflowOptions, api.ListCandlesticksWorkflowName, params)
	if err != nil {
		return api.ListCandlesticksWorkflowResults{}, err
	}

	// Get result and return
	err = exec.Get(ctx, &res)
	return res, err
}

// Info calls the service info.
func (c client) Info(ctx context.Context) (res api.ServiceInfoResults, err error) {
	workflowOptions := temporalclient.StartWorkflowOptions{
		TaskQueue: api.WorkerTaskQueueName,
	}

	// Execute workflow
	exec, err := c.temporal.ExecuteWorkflow(ctx, workflowOptions, api.ServiceInfoWorkflowName)
	if err != nil {
		return api.ServiceInfoResults{}, err
	}

	// Get result and return
	err = exec.Get(ctx, &res)
	return res, err
}
