package zaplog

import (
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
	"bytes"
	"path/filepath"
)

type Log struct {
	logger *zap.Logger
}

//var Log *zap.Logger //全局日志

//创建日志
func New(opts ...Option)  *Log {
	o := &Options{
		LogPath:"tmp/log/",
		LogName:"output",
		LogLevel:zapcore.DebugLevel,
		MaxSize:100,
		MaxAge:7,
		Stacktrace:zapcore.ErrorLevel,
		IsStdOut:true,
	}
	for _,opt := range opts {
		opt(o)
	}
	writers := []zapcore.WriteSyncer{newRollingFile(o.LogPath,o.LogName,o.MaxSize,o.MaxAge)}
	if o.IsStdOut {
		writers = append(writers, os.Stdout)
	}
	logger := newZapLogger(o.LogLevel,o.Stacktrace, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)
	return &Log{logger:logger}
}

func newZapLogger(level,stacktrace zapcore.Level, output zapcore.WriteSyncer) (*zap.Logger) {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "app",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		//EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		//},
		EncodeTime: zapcore.ISO8601TimeEncoder,
	}

	var encoder zapcore.Encoder
	dyn := zap.NewAtomicLevel()
	//encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	//encoder = zapcore.NewJSONEncoder(encCfg) // zapcore.NewConsoleEncoder(encCfg)
	dyn.SetLevel(level)
	encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoder = zapcore.NewJSONEncoder(encCfg)

	return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller(),zap.AddStacktrace(stacktrace),zap.AddCallerSkip(1))
}


//创建分割日志的writer
func newRollingFile(path,srvname string,maxSize,MaxAge int) zapcore.WriteSyncer {
	if err := os.MkdirAll(path, 0766); err != nil {
		panic(err)
		return nil
	}

	return newLumberjackWriteSyncer(&lumberjack.Logger{
		Filename:  filepath.Join(path,srvname + ".log"),
		MaxSize:   maxSize, //megabytes
		MaxAge:    MaxAge,   //days
		LocalTime: true,
		Compress:  false,
	})
}

type lumberjackWriteSyncer struct {
	*lumberjack.Logger
	buf       *bytes.Buffer
	logChan   chan []byte
	closeChan chan interface{}
	maxSize   int
}

func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
	ws := &lumberjackWriteSyncer{
		Logger:    l,
		buf:       bytes.NewBuffer([]byte{}),
		logChan:   make(chan []byte, 5000),
		closeChan: make(chan interface{}),
		maxSize:   1024,
	}
	go ws.run()
	return ws
}

func (l *lumberjackWriteSyncer) run() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			if l.buf.Len() > 0 {
				l.sync()
			}
		case bs := <-l.logChan:
			_, err := l.buf.Write(bs)
			if err != nil {
				continue
			}
			if l.buf.Len() > l.maxSize {
				l.sync()
			}
		case <-l.closeChan:
			l.sync()
			return
		}
	}
}

func (l *lumberjackWriteSyncer) Stop() {
	close(l.closeChan)
}

func (l *lumberjackWriteSyncer) Write(bs []byte) (int, error) {
	b := make([]byte, len(bs))
	for i, c := range bs {
		b[i] = c
	}
	l.logChan <- b
	return 0, nil
}

func (l *lumberjackWriteSyncer) Sync() error {
	return nil
}

func (l *lumberjackWriteSyncer) sync() error {
	defer l.buf.Reset()
	_, err := l.Logger.Write(l.buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

