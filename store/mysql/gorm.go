package mysql

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	gormSpanKey        = "__gorm_span"
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

//use db.WithContext change the session context
var op gorm.Plugin = &opentracingPlugin{}

func InitPlugin(db *gorm.DB) {
	db.Use(op)
}

func before(db *gorm.DB) {
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")
	db.InstanceSet(gormSpanKey, span)
	return
}

func after(db *gorm.DB) {
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	defer span.Finish()

	ext.DBInstance.Set(span, db.Name())
	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}
	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	return
}

type opentracingPlugin struct{}

func (op *opentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *opentracingPlugin) Initialize(db *gorm.DB) (err error) {

	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}
