package api

import (
	"time"

	"github.com/cryptellation/candlesticks/pkg/candlestick"
	"github.com/cryptellation/candlesticks/pkg/period"
)

const (
	// WorkerTaskQueueName is the name of the task queue for the cryptellation worker.
	WorkerTaskQueueName = "CryptellationCandlesticksTaskQueue"
)

const (
	// ListCandlesticksWorkflowName is the name of the workflow to get candlesticks.
	ListCandlesticksWorkflowName = "ListCandlesticksWorkflow"
)

type (
	// ListCandlesticksWorkflowParams is the parameters of the ListCandlesticks workflow.
	ListCandlesticksWorkflowParams struct {
		Exchange string
		Pair     string
		Period   period.Symbol
		Start    *time.Time
		End      *time.Time
		Limit    uint
	}

	// ListCandlesticksWorkflowResults is the result of the ListCandlesticks workflow.
	ListCandlesticksWorkflowResults struct {
		List *candlestick.List
	}
)

const (
	// ServiceInfoWorkflowName is the name of the workflow to get the service info.
	ServiceInfoWorkflowName = "ServiceInfoWorkflow"
)

type (
	// ServiceInfoParams contains the parameters of the service info workflow.
	ServiceInfoParams struct{}

	// ServiceInfoResults contains the result of the service info workflow.
	ServiceInfoResults struct {
		Version string
	}
)
