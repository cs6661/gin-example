package logger

import (
	"fmt"
	"gin-example/config"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()
	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。
	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewConsoleEncoder(encodeConfig)
	return encoder
}

// 负责日志写入的位置
func getLogWriter(config *config.LogConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", config.Path, config.LogFileName), // 文件位置
		MaxSize:    config.LumberJack.MaxSize,                             // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     config.LumberJack.MaxAge,                              // 保留旧文件的最大天数
		MaxBackups: config.LumberJack.MaxBackups,                          // 保留旧文件的最大个数
		Compress:   false,                                                 // 是否压缩/归档旧文件
	}
	if config.Stdout {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(lumberJackLogger))

	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitLogger(conf *config.LogConfig) {
	// 获取日志写入位置
	writeSyncer := getLogWriter(conf)
	// 获取日志编码格式
	encoder := getEncoder()

	// 获取日志最低等级，即>=该等级，才会被写入。
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(conf.Level))
	if err != nil {
		return
	}
	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, l)
	Logger = zap.New(core, zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(Logger)

}
