package clients

import (
	"github.com/cryptellation/candlesticks/api"
	"go.temporal.io/sdk/workflow"
)

// WfClient is a client for the cryptellation candlesticks service from a workflow perspective.
type WfClient interface {
	// ListCandlesticks lists candlesticks from Cryptellation service.
	ListCandlesticks(
		ctx workflow.Context,
		params api.ListCandlesticksWorkflowParams,
		childWorkflowOptions *workflow.ChildWorkflowOptions,
	) (result api.ListCandlesticksWorkflowResults, err error)
}

type wfClient struct{}

// NewWfClient creates a new workflow client.
// This client is used to call workflows from within other workflows.
// It is not used to call workflows from outside the workflow environment.
func NewWfClient() WfClient {
	return wfClient{}
}

// ListCandlesticks lists candlesticks from Cryptellation service.
func (wfClient) ListCandlesticks(
	ctx workflow.Context,
	params api.ListCandlesticksWorkflowParams,
	childWorkflowOptions *workflow.ChildWorkflowOptions,
) (result api.ListCandlesticksWorkflowResults, err error) {
	// Set default options
	ctx = setDefaultChildWorkflowOptions(ctx, childWorkflowOptions)

	// Get candlesticks
	err = workflow.ExecuteChildWorkflow(ctx, api.ListCandlesticksWorkflowName, params).Get(ctx, &result)
	return result, err
}

func setDefaultChildWorkflowOptions(
	ctx workflow.Context,
	childWorkflowOptions *workflow.ChildWorkflowOptions,
) workflow.Context {
	// Create default child workflow options
	if childWorkflowOptions == nil {
		childWorkflowOptions = &workflow.ChildWorkflowOptions{}
	}

	// Set default options
	if childWorkflowOptions.TaskQueue == "" {
		childWorkflowOptions.TaskQueue = api.WorkerTaskQueueName
	}

	return workflow.WithChildOptions(ctx, *childWorkflowOptions)
}
