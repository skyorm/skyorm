package skyorm

var (
	// DefaultLogger is a dummy logger that does nothing. It is used
	// in case nil is passed into provider.
	DefaultLogger Logger = new(dl)
)

// Logger interface.
type Logger interface {
	Printf(format string, v ...interface{})
}

// dl is "dummy" Logger implementation, used as DefaultLogger.
type dl struct{}

// Printf implements Logger interface.
func (*dl) Printf(_ string, _ ...interface{}) {
}
