package binance

import (
	"context"
	"errors"
	"io"
	"time"

	client "github.com/adshao/go-binance/v2"
	"github.com/cryptellation/candlesticks/svc/exchanges"
	"github.com/cryptellation/candlesticks/svc/exchanges/binance/entities"
	"go.temporal.io/sdk/worker"
)

// Activities is the struct that contains all the activities for the Binance exchange.
type Activities struct {
	Client *client.Client
}

// New creates a new Binance activities struct.
func New(apiKey, secretKey string) (*Activities, error) {
	// Validate that API key and secret key are not empty
	if apiKey == "" {
		return nil, entities.WrapError(errors.New("API key cannot be empty"))
	}
	if secretKey == "" {
		return nil, entities.WrapError(errors.New("secret key cannot be empty"))
	}

	c := client.NewClient(apiKey, secretKey)
	c.Logger.SetOutput(io.Discard)

	// Return service
	return &Activities{
		Client: c,
	}, nil
}

// Name returns the name of the Binance activities.
func (a *Activities) Name() string {
	return entities.ExchangeName
}

// Register registers the Binance activities with the given worker.
func (a *Activities) Register(_ worker.Worker) {
	// No need to register the Binance activities, they are already registered
	// with its parent.
}

// GetCandlesticksActivity gets the candlesticks for the given pair and period.
func (a *Activities) GetCandlesticksActivity(
	ctx context.Context,
	params exchanges.GetCandlesticksActivityParams,
) (exchanges.GetCandlesticksActivityResults, error) {
	a.Client.Debug = true

	service := a.Client.NewKlinesService()

	// Set symbol
	service.Symbol(entities.BinanceSymbol(params.Pair))

	// Set interval
	binanceInterval, err := entities.PeriodToInterval(params.Period)
	if err != nil {
		return exchanges.GetCandlesticksActivityResults{}, entities.WrapError(err)
	}
	service.Interval(binanceInterval)

	// Set start and end time
	service.StartTime(entities.TimeToKLineTime(params.Start))
	service.EndTime(entities.TimeToKLineTime(params.End))

	// Set limit
	if params.Limit > 0 {
		service.Limit(params.Limit)
	}

	// Get KLines
	kl, err := service.Do(ctx)
	if err != nil {
		return exchanges.GetCandlesticksActivityResults{}, entities.WrapError(err)
	}

	// Change them to right format
	list, err := entities.KLinesToCandlesticks(params.Pair, params.Period, kl, time.Now())
	if err != nil {
		return exchanges.GetCandlesticksActivityResults{}, entities.WrapError(err)
	}

	return exchanges.GetCandlesticksActivityResults{
		List: list,
	}, nil
}
