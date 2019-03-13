package cl

// StringClosure is a function that returns a string, used to defer execution of expensive logging operations
type StringClosure func() string

// Value is the generic list of things processed by the log chan
type Value []interface{}

// Fatal is a log value that indicates level and how to interpret the interface slice
type Fatal Value

// Fatalf is a log value that indicates level and how to interpret the interface slice
type Fatalf Value

// Ftl is a log type that is just one string
type Ftl string

// Fatalc is for passing a closure when the log entry is expensive to compute
type Fatalc StringClosure

// Error is a log value that indicates level and how to interpret the interface slice
type Error Value

// Errorf is a log value that indicates level and how to interpret the interface slice
type Errorf Value

// Err is a log type that is just one string
type Err string

// Errorc is for passing a closure when the log entry is expensive to compute
type Errorc StringClosure

// Warn is a log value that indicates level and how to interpret the interface slice
type Warn Value

// Warnf is a log value that indicates level and how to interpret the interface slice
type Warnf Value

// Wrn is a log type that is just one string
type Wrn string

// Warnc is for passing a closure when the log entry is expensive to compute
type Warnc StringClosure

// Info is a log value that indicates level and how to interpret the interface slice
type Info Value

// Infof is a log value that indicates level and how to interpret the interface slice
type Infof Value

// Inf is a log type that is just one string
type Inf string

// Infoc is for passing a closure when the log entry is expensive to compute
type Infoc StringClosure

// Debug is a log value that indicates level and how to interpret the interface slice
type Debug Value

// Debugf is a log value that indicates level and how to interpret the interface slice
type Debugf Value

// Dbg is a log type that is just one string
type Dbg string

// Debugc is for passing a closure when the log entry is expensive to compute
type Debugc StringClosure

// Trace is a log value that indicates level and how to interpret the interface slice
type Trace Value

// Tracef is a log value that indicates level and how to interpret the interface slice
type Tracef Value

// Trc is a log type that is just one string
type Trc string

// Tracec is for passing a closure when the log entry is expensive to compute
type Tracec StringClosure

// A SubSystem is a logger with a specific prefix name prepended  to the entry
type SubSystem struct {
	Name        string
	Ch          chan interface{}
	Level       int
	LevelString string
}

type Registry map[string]*SubSystem
