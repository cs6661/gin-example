package db

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"

	log "gin-example/pkg/logger"
)

const (
	Silent logger.LogLevel = iota + 1
	Error
	Warn
	Info
)

type Config struct {
	SlowThreshold time.Duration
	//Colorful                  bool
	IgnoreRecordNotFoundError bool
	LogLevel                  logger.LogLevel
}

func NewLogger(config Config) logger.Interface {
	var (
		infoStr      = "%s [info] "
		warnStr      = "%s [warn] "
		errStr       = "%s [error] "
		traceStr     = "file:%s runtime:%.3fms rows:%v sql:%s"
		traceWarnStr = "file:%s err:%s runtime:%.3fms rows:%v sql:%s"
		traceErrStr  = "file:%s err:%s runtime:%.3fms rows:%v sql:%s"
	)

	return &Logger{
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type Logger struct {
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		logs := log.Ctxs(ctx)
		logs.Info(l.infoStr + msg)
		//logs.Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		logs := log.Ctxs(ctx)
		logs.Warn(l.warnStr + msg)
		//logs.Warnf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		logs := log.Ctxs(ctx)
		logs.Error(l.warnStr + msg)
		//logs.Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= Silent {
		return
	}
	elapsed := time.Since(begin)
	lg := log.Ctxs(ctx)

	switch {
	case err != nil && l.LogLevel >= Error && !l.IgnoreRecordNotFoundError:
		if err == gorm.ErrRecordNotFound {
			sql, rows := fc()
			lg.Info("[SQL]",
				zap.String("file", utils.FileWithLineNum()),
				zap.Float64("runtime", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("rows", rows),
				zap.String("sql", sql),
			)
			break
		}
		sql, rows := fc()
		lg.Error("[SQL]",
			zap.String("file", utils.FileWithLineNum()),
			zap.String("error", err.Error()),
			zap.Float64("runtime", float64(elapsed.Nanoseconds())/1e6),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		lg.Warn("[SQL]",
			zap.String("file", utils.FileWithLineNum()),
			zap.String("error", slowLog),
			zap.Float64("runtime", float64(elapsed.Nanoseconds())/1e6),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	case l.LogLevel == Info:
		sql, rows := fc()
		lg.Debug("[SQL]",
			zap.String("file", utils.FileWithLineNum()),
			zap.Float64("runtime", float64(elapsed.Nanoseconds())/1e6),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	}
}
