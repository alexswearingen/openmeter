package billingservice

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/openmeterio/openmeter/openmeter/app"
	"github.com/openmeterio/openmeter/openmeter/billing"
	"github.com/openmeterio/openmeter/openmeter/billing/service/invoicecalc"
	"github.com/openmeterio/openmeter/openmeter/billing/service/lineservice"
	"github.com/openmeterio/openmeter/openmeter/customer"
	customerentity "github.com/openmeterio/openmeter/openmeter/customer/entity"
	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/openmeter/productcatalog/feature"
	"github.com/openmeterio/openmeter/openmeter/streaming"
	"github.com/openmeterio/openmeter/openmeter/watermill/eventbus"
	"github.com/openmeterio/openmeter/pkg/framework/transaction"
)

var _ billing.Service = (*Service)(nil)

type Service struct {
	adapter            billing.Adapter
	customerService    customer.CustomerService
	appService         app.Service
	logger             *slog.Logger
	invoiceCalculator  invoicecalc.Calculator
	featureService     feature.FeatureConnector
	meterRepo          meter.Repository
	streamingConnector streaming.Connector

	lineService *lineservice.Service
	publisher   eventbus.Publisher
}

type Config struct {
	Adapter            billing.Adapter
	CustomerService    customer.CustomerService
	AppService         app.Service
	Logger             *slog.Logger
	FeatureService     feature.FeatureConnector
	MeterRepo          meter.Repository
	StreamingConnector streaming.Connector
	Publisher          eventbus.Publisher
}

func (c Config) Validate() error {
	if c.Adapter == nil {
		return errors.New("adapter cannot be null")
	}

	if c.CustomerService == nil {
		return errors.New("customer service cannot be null")
	}

	if c.AppService == nil {
		return errors.New("app service cannot be null")
	}

	if c.Logger == nil {
		return errors.New("logger cannot be null")
	}

	if c.FeatureService == nil {
		return errors.New("feature connector cannot be null")
	}

	if c.MeterRepo == nil {
		return errors.New("meter repo cannot be null")
	}

	if c.StreamingConnector == nil {
		return errors.New("streaming connector cannot be null")
	}

	if c.Publisher == nil {
		return errors.New("publisher cannot be null")
	}

	return nil
}

func New(config Config) (*Service, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	svc := &Service{
		adapter:            config.Adapter,
		customerService:    config.CustomerService,
		appService:         config.AppService,
		logger:             config.Logger,
		featureService:     config.FeatureService,
		meterRepo:          config.MeterRepo,
		streamingConnector: config.StreamingConnector,
		publisher:          config.Publisher,
	}

	lineSvc, err := lineservice.New(lineservice.Config{
		BillingAdapter:     config.Adapter,
		FeatureService:     config.FeatureService,
		MeterRepo:          config.MeterRepo,
		StreamingConnector: config.StreamingConnector,
	})
	if err != nil {
		return nil, fmt.Errorf("creating line service: %w", err)
	}

	svc.lineService = lineSvc

	calculator, err := invoicecalc.New(invoicecalc.Config{
		LineService: lineSvc,
	})
	if err != nil {
		return nil, fmt.Errorf("creating invoice calculator: %w", err)
	}

	svc.invoiceCalculator = calculator

	return svc, nil
}

func (s Service) WithInvoiceCalculator(calc invoicecalc.Calculator) *Service {
	s.invoiceCalculator = calc

	return &s
}

func (s Service) InvoiceCalculator() invoicecalc.Calculator {
	return s.invoiceCalculator
}

// TranscationForGatheringInvoiceManipulation is a helper function that wraps the given function in a transaction and ensures that
// an update lock is held on the customer record. This is useful when you need to manipulate the gathering invoices, as we cannot lock an
// invoice, that doesn't exist yet.
func TranscationForGatheringInvoiceManipulation[T any](ctx context.Context, svc *Service, customer customerentity.CustomerID, fn func(ctx context.Context) (T, error)) (T, error) {
	if err := customer.Validate(); err != nil {
		var empty T
		return empty, fmt.Errorf("validating customer: %w", err)
	}

	// NOTE: This should not be in transaction, or we can get a conflict for parallel writes
	err := svc.adapter.UpsertCustomerOverride(ctx, customer)
	if err != nil {
		var empty T
		return empty, fmt.Errorf("upserting customer override: %w", err)
	}

	return transaction.Run(ctx, svc.adapter, func(ctx context.Context) (T, error) {
		if err := svc.adapter.LockCustomerForUpdate(ctx, customer); err != nil {
			var empty T
			return empty, fmt.Errorf("locking customer for update: %w", err)
		}

		return fn(ctx)
	})
}
