package orm

import "github.com/zeromicro/go-zero/core/metric"

const gormNameSpace = "gorm_client"

var metricGormDuration = metric.NewHistogramVec(&metric.HistogramVecOpts{
	Namespace: gormNameSpace,
	Subsystem: "requests",
	Name:      "duration_ms",
	Help:      "gorm client duration(ms)",
	Labels:    []string{"table", "method"},
	Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
})

var metricGormErrCount = metric.NewCounterVec(&metric.CounterVecOpts{
	Namespace: gormNameSpace,
	Subsystem: "requests",
	Name:      "error_count",
	Help:      "gorm client requests error count",
	Labels:    []string{"table", "method", "is_error"},
})
