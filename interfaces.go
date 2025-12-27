package app_monitoring

import (
	"time"

	"github.com/gin-gonic/gin"
	ae "github.com/piyushkumar96/app-error"
	pubsub "github.com/piyushkumar96/generic-pubsub"
	"github.com/prometheus/client_golang/prometheus"
)

// RouterMetricsInterface defines the contract for router-level HTTP metrics.
// Implement this interface to provide custom router metrics or mock implementations for testing.
type RouterMetricsInterface interface {
	// LogMetrics returns a Gin middleware that logs HTTP request metrics.
	LogMetrics(metricsPath string) gin.HandlerFunc

	// GetHTTPRequestsMetric returns the underlying HTTP requests counter metric.
	GetHTTPRequestsMetric() *prometheus.CounterVec

	// GetHTTPRequestsLatencyMillisMetric returns the underlying HTTP request latency histogram.
	GetHTTPRequestsLatencyMillisMetric() *prometheus.HistogramVec

	// GetHTTPRequestSizeBytesMetric returns the underlying HTTP request size histogram.
	GetHTTPRequestSizeBytesMetric() *prometheus.HistogramVec

	// GetHTTPResponseSizeBytesMetric returns the underlying HTTP response size histogram.
	GetHTTPResponseSizeBytesMetric() *prometheus.HistogramVec
}

// DBMetricsInterface defines the contract for database operation metrics.
// Implement this interface to provide custom database metrics or mock implementations for testing.
type DBMetricsInterface interface {
	// LogMetricsPre should be called before a database operation.
	// Returns the start time for latency calculation.
	LogMetricsPre(dbMetricsLabelValues *DBMetricsLabelValues) time.Time

	// LogMetricsPost should be called after a database operation completes.
	LogMetricsPost(appErr *ae.AppError, dbMetricsLabelValues *DBMetricsLabelValues, opsExecTime time.Time)

	// GetOperationsTotalMetric returns the underlying operations counter metric.
	GetOperationsTotalMetric() *prometheus.CounterVec

	// GetOperationsLatencyMillisMetric returns the underlying operations latency histogram.
	GetOperationsLatencyMillisMetric() *prometheus.HistogramVec
}

// DownstreamServiceMetricsInterface defines the contract for downstream HTTP service metrics.
// Implement this interface to provide custom downstream metrics or mock implementations for testing.
type DownstreamServiceMetricsInterface interface {
	// LogMetricsPre should be called before making a downstream HTTP call.
	LogMetricsPre(dssMetricsLabelValues *DownstreamServiceMetricsLabelValues)

	// LogMetricsPost should be called after a downstream HTTP call completes.
	LogMetricsPost(success bool, dssMetricsLabelValues *DownstreamServiceMetricsLabelValues, httpMetrics *HTTPMetrics)

	// GetHTTPRequestsMetric returns the underlying HTTP requests counter metric.
	GetHTTPRequestsMetric() *prometheus.CounterVec

	// GetHTTPRequestsLatencyMillisMetric returns the underlying HTTP request latency histogram.
	GetHTTPRequestsLatencyMillisMetric() *prometheus.HistogramVec

	// GetHTTPRequestSizeBytesMetric returns the underlying HTTP request size histogram.
	GetHTTPRequestSizeBytesMetric() *prometheus.HistogramVec

	// GetHTTPResponseSizeBytesMetric returns the underlying HTTP response size histogram.
	GetHTTPResponseSizeBytesMetric() *prometheus.HistogramVec
}

// CronJobMetricsInterface defines the contract for cron job execution metrics.
// Implement this interface to provide custom cron job metrics or mock implementations for testing.
type CronJobMetricsInterface interface {
	// LogMetricsPre should be called at the start of a cron job execution.
	// Returns the start time for latency calculation.
	LogMetricsPre(cjMetricsLabelValues *CronJobMetricsLabelValues) time.Time

	// LogMetricsPost should be called after a cron job execution completes.
	LogMetricsPost(appErr *ae.AppError, cjMetricsLabelValues *CronJobMetricsLabelValues, opsExecTime time.Time)

	// GetJobExecutionTotalMetric returns the underlying job execution counter metric.
	GetJobExecutionTotalMetric() *prometheus.CounterVec

	// GetJobExecutionLatencyMillisMetric returns the underlying job execution latency histogram.
	GetJobExecutionLatencyMillisMetric() *prometheus.HistogramVec
}

// PSMetricsInterface defines the contract for pub/sub messaging metrics.
// Implement this interface to provide custom pub/sub metrics or mock implementations for testing.
type PSMetricsInterface interface {
	// LogMetricsPre should be called before publishing a message or when starting to process a consumed message.
	// Returns the start time for latency calculation.
	LogMetricsPre(psMetricsLabelValues *PSMetricsLabelValues) time.Time

	// LogMetricsPost should be called after a pub/sub operation completes.
	LogMetricsPost(psMetricsLabelValues *PSMetricsLabelValues, eventTxnData *pubsub.EventTxnData)

	// GetTotalMessagesConsumedMetric returns the underlying messages consumed counter metric.
	GetTotalMessagesConsumedMetric() *prometheus.CounterVec

	// GetTotalMessagesPublishedMetric returns the underlying messages published counter metric.
	GetTotalMessagesPublishedMetric() *prometheus.CounterVec

	// GetMessagesPublishedLatencyMillisMetric returns the underlying message publish latency histogram.
	GetMessagesPublishedLatencyMillisMetric() *prometheus.HistogramVec

	// GetMessagesPublishedSizeBytesMetric returns the underlying published message size histogram.
	GetMessagesPublishedSizeBytesMetric() *prometheus.HistogramVec
}

// AppMetricsInterface defines the contract for application-level error metrics.
// Implement this interface to provide custom app metrics or mock implementations for testing.
type AppMetricsInterface interface {
	// LogMetrics increments the application error counter for each provided error code.
	LogMetrics(errCodes []string)

	// DecrementAppErrorCount decrements the application error counter for a specific error code.
	DecrementAppErrorCount(errCode string)

	// GetApplicationErrorsCounterMetric returns the underlying application errors gauge metric.
	GetApplicationErrorsCounterMetric() *prometheus.GaugeVec
}

// Compile-time interface implementation checks
var (
	_ RouterMetricsInterface            = (*RouterMetrics)(nil)
	_ DBMetricsInterface                = (*DBMetrics)(nil)
	_ DownstreamServiceMetricsInterface = (*DownstreamServiceMetrics)(nil)
	_ CronJobMetricsInterface           = (*CronJobMetrics)(nil)
	_ PSMetricsInterface                = (*PSMetrics)(nil)
	_ AppMetricsInterface               = (*AppMetrics)(nil)
)
