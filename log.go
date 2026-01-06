package teapot

import (
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Option
type Option func(*Logger)

// WithLevel
func WithLevel(level Level) Option { return func(l *Logger) { l.lvl = level } }

// WithWriter
func WithWriter(w io.Writer) Option { return func(l *Logger) { l.w = w } }

// Logger
type Logger struct {
	mu sync.Mutex
	p  *sync.Pool

	w io.Writer

	lvl Level
	f   Format
}

// New
func New(opts ...Option) *Logger {
	l := &Logger{
		w:   os.Stdout,
		lvl: DEBUG,
		f:   JSON,
		mu:  sync.Mutex{},
		p: &sync.Pool{
			New: func() any {
				buf := make([]byte, 0, 1024)
				return &buf
			},
		},
	}
	for _, o := range opts {
		o(l)
	}
	return l
}

// Debug
func (l *Logger) Debug(format string, attrs ...Attr) {
	l.printf(format, DEBUG, attrs...)
}

// Info
func (l *Logger) Info(format string, attrs ...Attr) {
	l.printf(format, INFO, attrs...)
}

// Error
func (l *Logger) Error(format string, attrs ...Attr) {
	l.printf(format, ERROR, attrs...)
}

// Fatal
func (l *Logger) Fatal(format string, attrs ...Attr) {
	l.printf(format, FATAL, attrs...)
	os.Exit(1)
}

func (l *Logger) printf(format string, level Level, attrs ...Attr) {
	if level < l.lvl {
		return
	}

	ptr := l.p.Get().(*[]byte)
	buf := (*ptr)[:0]

	switch l.f {
	// TODO
	// implement text logger
	// case TEXT:
	case JSON:
		buf = append(buf, `{"time":"`...)
		buf = time.Now().UTC().AppendFormat(buf, time.RFC3339Nano)
		buf = append(buf, `", "level":"`...)
		buf = append(buf, levelName[l.lvl]...)
		buf = append(buf, `", "message":"`...)
		buf = append(buf, format...)
		buf = append(buf, `"`...)
		for _, attr := range attrs {
			buf = append(buf, `,"`...)
			buf = append(buf, attr.key...)
			buf = append(buf, `":`...)

			switch attr.kfield {
			case kint:
				buf = strconv.AppendInt(buf, int64(attr.intValue), 10)
			case kint16:
				buf = strconv.AppendInt(buf, int64(attr.int16Value), 10)
			case kint32:
				buf = strconv.AppendInt(buf, int64(attr.int32Value), 10)
			case kint64:
				buf = strconv.AppendInt(buf, attr.int64Value, 10)
			case kbool:
				buf = strconv.AppendBool(buf, attr.boolValue)
			case kstring:
				buf = strconv.AppendQuote(buf, attr.stringValue)
			case kstringslice:
				s := strings.Join(attr.stringSliceValue, ",")
				buf = strconv.AppendQuote(buf, s)
			case kerr:
				buf = strconv.AppendQuote(buf, attr.errValue.Error())
			}
		}

		if level >= ERROR {
			buf = append(buf, `, "stacktrace":`...)

			var sb = make([]byte, 2048)
			n := runtime.Stack(sb, false)
			buf = strconv.AppendQuote(buf, string(sb[:n]))
		}

		buf = append(buf, "}"...)
	default:
		return
	}

	buf = append(buf, '\n')

	l.mu.Lock()
	l.w.Write(buf)
	l.mu.Unlock()

	*ptr = buf
	l.p.Put(ptr)
}
