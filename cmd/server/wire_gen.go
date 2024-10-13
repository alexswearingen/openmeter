// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	kafka2 "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/openmeterio/openmeter/config"
	"github.com/openmeterio/openmeter/openmeter/app"
	"github.com/openmeterio/openmeter/openmeter/ent/db"
	"github.com/openmeterio/openmeter/openmeter/ingest"
	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/openmeter/namespace"
	"github.com/openmeterio/openmeter/openmeter/streaming"
	"github.com/openmeterio/openmeter/openmeter/watermill/driver/kafka"
	"github.com/openmeterio/openmeter/openmeter/watermill/eventbus"
	"github.com/openmeterio/openmeter/pkg/kafka/metrics"
	"go.opentelemetry.io/otel/metric"
	"log/slog"
)

// Injectors from wire.go:

func initializeApplication(ctx context.Context, conf config.Configuration, logger *slog.Logger) (Application, func(), error) {
	telemetryConfig := conf.Telemetry
	metricsTelemetryConfig := telemetryConfig.Metrics
	appMetadata := metadata(conf)
	resource := app.NewTelemetryResource(appMetadata)
	meterProvider, cleanup, err := app.NewMeterProvider(ctx, metricsTelemetryConfig, resource, logger)
	if err != nil {
		return Application{}, nil, err
	}
	traceTelemetryConfig := telemetryConfig.Trace
	tracerProvider, cleanup2, err := app.NewTracerProvider(ctx, traceTelemetryConfig, resource, logger)
	if err != nil {
		cleanup()
		return Application{}, nil, err
	}
	textMapPropagator := app.NewDefaultTextMapPropagator()
	globalInitializer := app.GlobalInitializer{
		Logger:            logger,
		MeterProvider:     meterProvider,
		TracerProvider:    tracerProvider,
		TextMapPropagator: textMapPropagator,
	}
	aggregationConfiguration := conf.Aggregation
	clickHouseAggregationConfiguration := aggregationConfiguration.ClickHouse
	v, err := app.NewClickHouse(clickHouseAggregationConfiguration)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	v2 := conf.Meters
	inMemoryRepository := app.NewMeterRepository(v2)
	clickhouseConnector, err := app.NewClickHouseStreamingConnector(aggregationConfiguration, v, inMemoryRepository, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	postgresConfig := conf.Postgres
	meter := app.NewMeter(meterProvider, appMetadata)
	driver, cleanup3, err := app.NewPostgresDriver(ctx, postgresConfig, meterProvider, meter, tracerProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	db := app.NewDB(driver)
	entPostgresDriver, cleanup4 := app.NewEntPostgresDriver(db, logger)
	client := app.NewEntClient(entPostgresDriver)
	health := app.NewHealthChecker(logger)
	telemetryHandler := app.NewTelemetryHandler(metricsTelemetryConfig, health)
	v3, cleanup5 := app.NewTelemetryServer(telemetryConfig, telemetryHandler)
	producer, err := app.NewKafkaProducer(conf, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	metrics, err := app.NewKafkaMetrics(meter)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	eventsConfiguration := conf.Events
	ingestConfiguration := conf.Ingest
	kafkaIngestConfiguration := ingestConfiguration.Kafka
	kafkaConfiguration := kafkaIngestConfiguration.KafkaConfiguration
	logTelemetryConfig := telemetryConfig.Log
	brokerOptions := app.NewBrokerConfiguration(kafkaConfiguration, logTelemetryConfig, appMetadata, logger, meter)
	v4 := app.ServerProvisionTopics(eventsConfiguration)
	adminClient, err := app.NewKafkaAdminClient(kafkaConfiguration)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	topicProvisionerConfig := kafkaIngestConfiguration.TopicProvisionerConfig
	kafkaTopicProvisionerConfig := app.NewKafkaTopicProvisionerConfig(adminClient, logger, meter, topicProvisionerConfig)
	topicProvisioner, err := app.NewKafkaTopicProvisioner(kafkaTopicProvisionerConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	publisherOptions := kafka.PublisherOptions{
		Broker:           brokerOptions,
		ProvisionTopics:  v4,
		TopicProvisioner: topicProvisioner,
	}
	publisher, cleanup6, err := app.NewServerPublisher(ctx, eventsConfiguration, publisherOptions, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	eventbusPublisher, err := app.NewEventBusPublisher(publisher, eventsConfiguration, logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	namespacedTopicResolver, err := app.NewNamespacedTopicResolver(conf)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	collector, err := app.NewKafkaIngestCollector(producer, namespacedTopicResolver)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	ingestCollector, cleanup7, err := app.NewIngestCollector(conf, collector, logger, meter)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	namespaceHandler, err := app.NewKafkaNamespaceHandler(namespacedTopicResolver, topicProvisioner, conf)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	v5 := app.NewNamespaceHandlers(namespaceHandler, clickhouseConnector)
	namespaceConfiguration := conf.Namespace
	manager, err := app.NewNamespaceManager(v5, namespaceConfiguration)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Application{}, nil, err
	}
	v6 := app.NewTelemetryRouterHook(meterProvider, tracerProvider)
	application := Application{
		GlobalInitializer:  globalInitializer,
		StreamingConnector: clickhouseConnector,
		MeterRepository:    inMemoryRepository,
		EntClient:          client,
		TelemetryServer:    v3,
		KafkaProducer:      producer,
		KafkaMetrics:       metrics,
		EventPublisher:     eventbusPublisher,
		IngestCollector:    ingestCollector,
		NamespaceHandlers:  v5,
		NamespaceManager:   manager,
		Meter:              meter,
		RouterHook:         v6,
	}
	return application, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// TODO: is this necessary? Do we need a logger first?
func initializeLogger(conf config.Configuration) *slog.Logger {
	telemetryConfig := conf.Telemetry
	logTelemetryConfig := telemetryConfig.Log
	appMetadata := metadata(conf)
	resource := app.NewTelemetryResource(appMetadata)
	logger := app.NewLogger(logTelemetryConfig, resource)
	return logger
}

// wire.go:

type Application struct {
	app.GlobalInitializer

	StreamingConnector streaming.Connector
	MeterRepository    meter.Repository
	EntClient          *db.Client
	TelemetryServer    app.TelemetryServer
	KafkaProducer      *kafka2.Producer
	KafkaMetrics       *metrics.Metrics
	EventPublisher     eventbus.Publisher

	IngestCollector ingest.Collector

	NamespaceHandlers []namespace.Handler
	NamespaceManager  *namespace.Manager

	Meter metric.Meter

	RouterHook func(chi.Router)
}

func metadata(conf config.Configuration) app.Metadata {
	return app.Metadata{
		ServiceName:       "openmeter",
		Version:           version,
		Environment:       conf.Environment,
		OpenTelemetryName: "openmeter.io/backend",
	}
}
