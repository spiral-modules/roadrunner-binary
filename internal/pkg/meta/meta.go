package meta

import "strings"

// version value will be set during compilation.
var (
	version   = "local"
	buildTime = "development"
)

// Version returns version value (without `v` prefix).
func Version() string { return strings.TrimLeft(version, "vV ") } // TODO check 2nd char "is numeric?"

// BuildTime ... // FIXME
func BuildTime() string { return buildTime } // TODO test me
