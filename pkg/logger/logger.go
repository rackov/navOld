package logger

import (
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Уровни логирования, соответствующие logrus
const (
	TRACE = logrus.TraceLevel
	DEBUG = logrus.DebugLevel
	INFO  = logrus.InfoLevel
	WARN  = logrus.WarnLevel
	ERROR = logrus.ErrorLevel
)

type Logger struct {
	*logrus.Logger
}

var (
	defaultLogger *Logger
	once          sync.Once
)

// New создает и настраивает новый экземпляр логгера.
// filePath - путь к файлу для записи логов. Если пустой, логи пишутся только в stdout.
func New(level logrus.Level, filePath ...string) *Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	l.SetOutput(os.Stdout)
	l.SetLevel(level)

	if len(filePath) > 0 && filePath[0] != "" {
		l.SetFormatter(&logrus.JSONFormatter{ //TextFormatter{
			// DisableColors:   true,
			//  FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
		multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename:   filePath[0],
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		})
		l.SetOutput(multiWriter)
	}

	return &Logger{l}
}

// Init инициализирует глобальный логгер-синглтон.
// Ее нужно вызвать один раз при старте приложения.
func Init(level logrus.Level, filePath ...string) {
	once.Do(func() {
		defaultLogger = New(level, filePath...)
	})
}

// ParseLevel преобразует строку в logrus.Level.
func ParseLevel(levelStr string) (logrus.Level, error) {
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return DEBUG, err // Возвращаем уровень по умолчанию в случае ошибки
	}
	return level, nil
}

// --- Методы для установки уровня ---

// SetLevel динамически меняет уровень логирования
func (l *Logger) SetLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

// SetGlobalLevel устанавливает уровень для глобального логгера
func SetGlobalLevel(level logrus.Level) {
	defaultLogger.SetLevel(level)
}

// --- Глобальные функции-хелперы для удобства ---

// Они вызывают методы у глобального экземпляра defaultLogger
func WithField(key string, value interface{}) *logrus.Entry {
	return defaultLogger.WithField(key, value)
}
func WithFields(fields logrus.Fields) *logrus.Entry {
	return defaultLogger.WithFields(fields)
}
func WithError(err error) *logrus.Entry {
	return defaultLogger.WithError(err)
}
func Trace(args ...interface{}) {
	defaultLogger.Trace(args...)
}
func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args...)
}
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// GetLogger возвращает экземпляр глобального логгера.
// Это может быть полезно, если какой-то компонент хочет получить
// прямой доступ к логгеру, а не использовать глобальные функции.
func GetLogger() *Logger {
	return defaultLogger
}
