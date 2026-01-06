package teapot

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelName = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

type Format int

const (
	TEXT Format = iota
	JSON
)

type kField int

const (
	kint kField = iota
	kint16
	kint32
	kint64
	kbool
	kany
	kstring
	kstringslice
	kerr
)
