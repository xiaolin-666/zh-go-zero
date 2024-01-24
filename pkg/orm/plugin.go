package orm

import (
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func NewTraceAndMetricPlugin() *TraceAndMetricPlugin {
	return &TraceAndMetricPlugin{}
}

type TraceAndMetricPlugin struct {
}

func (t *TraceAndMetricPlugin) Name() string {
	return "TraceAndMetricPlugin"
}

func (t *TraceAndMetricPlugin) Initialize(db *gorm.DB) error {
	// before
	if err := db.Callback().Create().Before("gorm:createBefore").Register("createBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "create")
	}); err != nil {
		return err
	}
	if err := db.Callback().Query().Before("gorm:queryBefore").Register("queryBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "query")
	}); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:updateBefore").Register("updateBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "update")
	}); err != nil {
		return err
	}
	if err := db.Callback().Delete().Before("gorm:deleteBefore").Register("deleteBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "delete")
	}); err != nil {
		return err
	}
	if err := db.Callback().Row().Before("gorm:rowBefore").Register("rowBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "row")
	}); err != nil {
		return err
	}
	if err := db.Callback().Raw().Before("gorm:rawBefore").Register("rawBefore", func(db *gorm.DB) {
		t.beforeCallback(db, "raw")
	}); err != nil {
		return err
	}

	// after
	if err := db.Callback().Create().After("gorm:createAfter").Register("createAfter", func(db *gorm.DB) {
		t.afterCallback(db, "create")
	}); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:queryAfter").Register("queryAfter", func(db *gorm.DB) {
		t.afterCallback(db, "query")
	}); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:updateAfter").Register("updateAfter", func(db *gorm.DB) {
		t.afterCallback(db, "update")
	}); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:deleteAfter").Register("deleteAfter", func(db *gorm.DB) {
		t.afterCallback(db, "delete")
	}); err != nil {
		return err
	}
	if err := db.Callback().Row().After("gorm:rowAfter").Register("rowAfter", func(db *gorm.DB) {
		t.afterCallback(db, "row")
	}); err != nil {
		return err
	}
	if err := db.Callback().Raw().After("gorm:rawAfter").Register("rawAfter", func(db *gorm.DB) {
		t.afterCallback(db, "raw")
	}); err != nil {
		return err
	}

	return nil
}

func (t *TraceAndMetricPlugin) beforeCallback(db *gorm.DB, operation string) {
	startTime := time.Now().Unix()
	db.InstanceSet("gorm:"+operation+"StartTime", startTime)

	ctx := db.Statement.Context
	tracer := trace.TracerFromContext(ctx)
	_, span := tracer.Start(ctx, "gorm:"+operation, oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	db.InstanceSet("gorm:"+operation+"Span", span)
}

func (t *TraceAndMetricPlugin) afterCallback(db *gorm.DB, operation string) {
	startTime, ok := db.InstanceGet("gorm:" + operation + "StartTime")
	if !ok {
		return
	}
	stTime := time.Unix(startTime.(int64), 0)
	metricGormDuration.Observe(time.Since(stTime).Milliseconds(), db.Statement.Table, operation)
	metricGormErrCount.Inc(db.Statement.Table, operation, strconv.FormatBool(db.Statement.Error != nil))

	value, ok := db.InstanceGet("gorm:" + operation + "Span")
	if !ok {
		return
	}
	span := value.(oteltrace.Span)
	if db.Statement.Error != nil {
		span.SetStatus(codes.Error, operation+" operation failed")
		span.RecordError(db.Statement.Error)
	}
	span.SetAttributes(
		semconv.DBSQLTable(db.Statement.Table),
		semconv.DBStatement(db.Statement.SQL.String()),
	)
	span.End()
}
