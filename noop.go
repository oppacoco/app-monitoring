package app_monitoring

import (
	"time"

	"github.com/gin-gonic/gin"
	ae "github.com/piyushkumar96/app-error"
	pubsub "github.com/piyushkumar96/generic-pubsub"
	"github.com/prometheus/client_golang/prometheus"
)

// NoOpRouterMetrics is a no-operation implementation of RouterMetricsInterface.
// Use this for testing or when you want to disable router metrics collection.
type NoOpRouterMetrics struct{}

// NewNoOpRouterMetrics creates a new no-op router metrics instance.
func NewNoOpRouterMetrics() RouterMetricsInterface {
	return &NoOpRouterMetrics{}
}

// LogMetrics returns a pass-through middleware that does nothing.
func (n *NoOpRouterMetrics) LogMetrics(_ string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// GetHTTPRequestsMetric returns nil.
func (n *NoOpRouterMetrics) GetHTTPRequestsMetric() *prometheus.CounterVec {
	return nil
}

// GetHTTPRequestsLatencyMillisMetric returns nil.
func (n *NoOpRouterMetrics) GetHTTPRequestsLatencyMillisMetric() *prometheus.HistogramVec {
	return nil
}

// GetHTTPRequestSizeBytesMetric returns nil.
func (n *NoOpRouterMetrics) GetHTTPRequestSizeBytesMetric() *prometheus.HistogramVec {
	return nil
}

// GetHTTPResponseSizeBytesMetric returns nil.
func (n *NoOpRouterMetrics) GetHTTPResponseSizeBytesMetric() *prometheus.HistogramVec {
	return nil
}

// NoOpDBMetrics is a no-operation implementation of DBMetricsInterface.
// Use this for testing or when you want to disable database metrics collection.
type NoOpDBMetrics struct{}

// NewNoOpDBMetrics creates a new no-op database metrics instance.
func NewNoOpDBMetrics() DBMetricsInterface {
	return &NoOpDBMetrics{}
}

// LogMetricsPre does nothing and returns the current time.
func (n *NoOpDBMetrics) LogMetricsPre(_ *DBMetricsLabelValues) time.Time {
	return time.Now()
}

// LogMetricsPost does nothing.
func (n *NoOpDBMetrics) LogMetricsPost(_ *ae.AppError, _ *DBMetricsLabelValues, _ time.Time) {
}

// GetOperationsTotalMetric returns nil.
func (n *NoOpDBMetrics) GetOperationsTotalMetric() *prometheus.CounterVec {
	return nil
}

// GetOperationsLatencyMillisMetric returns nil.
func (n *NoOpDBMetrics) GetOperationsLatencyMillisMetric() *prometheus.HistogramVec {
	return nil
}

// NoOpDownstreamServiceMetrics is a no-operation implementation of DownstreamServiceMetricsInterface.
// Use this for testing or when you want to disable downstream service metrics collection.
type NoOpDownstreamServiceMetrics struct{}

// NewNoOpDownstreamServiceMetrics creates a new no-op downstream service metrics instance.
func NewNoOpDownstreamServiceMetrics() DownstreamServiceMetricsInterface {
	return &NoOpDownstreamServiceMetrics{}
}

// LogMetricsPre does nothing.
func (n *NoOpDownstreamServiceMetrics) LogMetricsPre(_ *DownstreamServiceMetricsLabelValues) {
}

// LogMetricsPost does nothing.
func (n *NoOpDownstreamServiceMetrics) LogMetricsPost(_ bool, _ *DownstreamServiceMetricsLabelValues, _ *HTTPMetrics) {
}

// GetHTTPRequestsMetric returns nil.
func (n *NoOpDownstreamServiceMetrics) GetHTTPRequestsMetric() *prometheus.CounterVec {
	return nil
}

// GetHTTPRequestsLatencyMillisMetric returns nil.
func (n *NoOpDownstreamServiceMetrics) GetHTTPRequestsLatencyMillisMetric() *prometheus.HistogramVec {
	return nil
}

// GetHTTPRequestSizeBytesMetric returns nil.
func (n *NoOpDownstreamServiceMetrics) GetHTTPRequestSizeBytesMetric() *prometheus.HistogramVec {
	return nil
}

// GetHTTPResponseSizeBytesMetric returns nil.
func (n *NoOpDownstreamServiceMetrics) GetHTTPResponseSizeBytesMetric() *prometheus.HistogramVec {
	return nil
}

// NoOpCronJobMetrics is a no-operation implementation of CronJobMetricsInterface.
// Use this for testing or when you want to disable cron job metrics collection.
type NoOpCronJobMetrics struct{}

// NewNoOpCronJobMetrics creates a new no-op cron job metrics instance.
func NewNoOpCronJobMetrics() CronJobMetricsInterface {
	return &NoOpCronJobMetrics{}
}

// LogMetricsPre does nothing and returns the current time.
func (n *NoOpCronJobMetrics) LogMetricsPre(_ *CronJobMetricsLabelValues) time.Time {
	return time.Now()
}

// LogMetricsPost does nothing.
func (n *NoOpCronJobMetrics) LogMetricsPost(_ *ae.AppError, _ *CronJobMetricsLabelValues, _ time.Time) {
}

// GetJobExecutionTotalMetric returns nil.
func (n *NoOpCronJobMetrics) GetJobExecutionTotalMetric() *prometheus.CounterVec {
	return nil
}

// GetJobExecutionLatencyMillisMetric returns nil.
func (n *NoOpCronJobMetrics) GetJobExecutionLatencyMillisMetric() *prometheus.HistogramVec {
	return nil
}

// NoOpPSMetrics is a no-operation implementation of PSMetricsInterface.
// Use this for testing or when you want to disable pub/sub metrics collection.
type NoOpPSMetrics struct{}

// NewNoOpPSMetrics creates a new no-op pub/sub metrics instance.
func NewNoOpPSMetrics() PSMetricsInterface {
	return &NoOpPSMetrics{}
}

// LogMetricsPre does nothing and returns the current time.
func (n *NoOpPSMetrics) LogMetricsPre(_ *PSMetricsLabelValues) time.Time {
	return time.Now()
}

// LogMetricsPost does nothing.
func (n *NoOpPSMetrics) LogMetricsPost(_ *PSMetricsLabelValues, _ *pubsub.EventTxnData) {
}

// GetTotalMessagesConsumedMetric returns nil.
func (n *NoOpPSMetrics) GetTotalMessagesConsumedMetric() *prometheus.CounterVec {
	return nil
}

// GetTotalMessagesPublishedMetric returns nil.
func (n *NoOpPSMetrics) GetTotalMessagesPublishedMetric() *prometheus.CounterVec {
	return nil
}

// GetMessagesPublishedLatencyMillisMetric returns nil.
func (n *NoOpPSMetrics) GetMessagesPublishedLatencyMillisMetric() *prometheus.HistogramVec {
	return nil
}

// GetMessagesPublishedSizeBytesMetric returns nil.
func (n *NoOpPSMetrics) GetMessagesPublishedSizeBytesMetric() *prometheus.HistogramVec {
	return nil
}

// NoOpAppMetrics is a no-operation implementation of AppMetricsInterface.
// Use this for testing or when you want to disable application error metrics collection.
type NoOpAppMetrics struct{}

// NewNoOpAppMetrics creates a new no-op application metrics instance.
func NewNoOpAppMetrics() AppMetricsInterface {
	return &NoOpAppMetrics{}
}

// LogMetrics does nothing.
func (n *NoOpAppMetrics) LogMetrics(_ []string) {
}

// DecrementAppErrorCount does nothing.
func (n *NoOpAppMetrics) DecrementAppErrorCount(_ string) {
}

// GetApplicationErrorsCounterMetric returns nil.
func (n *NoOpAppMetrics) GetApplicationErrorsCounterMetric() *prometheus.GaugeVec {
	return nil
}

// Compile-time interface implementation checks for NoOp types
var (
	_ RouterMetricsInterface            = (*NoOpRouterMetrics)(nil)
	_ DBMetricsInterface                = (*NoOpDBMetrics)(nil)
	_ DownstreamServiceMetricsInterface = (*NoOpDownstreamServiceMetrics)(nil)
	_ CronJobMetricsInterface           = (*NoOpCronJobMetrics)(nil)
	_ PSMetricsInterface                = (*NoOpPSMetrics)(nil)
	_ AppMetricsInterface               = (*NoOpAppMetrics)(nil)
)
