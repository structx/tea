package teapot

// Attr
type Attr struct {
	key    string
	kfield kField

	intValue         int
	int16Value       int16
	int32Value       int32
	int64Value       int64
	stringValue      string
	boolValue        bool
	stringSliceValue []string

	anyValue any

	errValue error
}

// Int
func Int(key string, value int) Attr {
	return Attr{key: key, kfield: kint, intValue: value}
}

// Int16
func Int16(key string, value int16) Attr {
	return Attr{key: key, kfield: kint16, int16Value: value}
}

// Int32
func Int32(key string, value int32) Attr {
	return Attr{key: key, kfield: kint32, int32Value: value}
}

// Int64
func Int64(key string, value int64) Attr {
	return Attr{key: key, kfield: kint64, int64Value: value}
}

// Bool
func Bool(key string, value bool) Attr {
	return Attr{key: key, kfield: kbool, boolValue: value}
}

// String
func String(key, value string) Attr {
	return Attr{key: key, kfield: kstring, stringValue: value}
}

// Any
func Any(key string, value any) Attr {
	return Attr{key: key, kfield: kany, anyValue: value}
}

// StringSlice
func StringSlice(key string, value []string) Attr {
	return Attr{key: key, kfield: kstringslice, stringSliceValue: value}
}

// Error
func Error(err error) Attr {
	return Attr{key: "error", kfield: kerr, errValue: err}
}
