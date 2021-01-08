package states

import warp10 "github.com/miton18/go-warp10/base"

const (
	DefaultSeparator = "."
)

var (
	// DefaultRegistry is ...
	DefaultRegistry = NewRegistry()

	// DefaultStateValue is used when the state is not known
	DefaultStateValue = defaultStateValue{}
)

// MustRegister register a state to the default Registry
// Panic if the state already exists
func MustRegister(state State) {
	DefaultRegistry.MustRegister(state)
}

// Register register a state to the default Registry
func Register(state State) error {
	return DefaultRegistry.Register(state)
}

// Gather return a GTS list of all his states
func Gather() warp10.GTSList {
	return DefaultRegistry.Gather()
}
