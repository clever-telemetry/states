package states

import (
	"net/http"
)

// Our custom state
// The best way to have enums in Go
type MyState string

// Implement fmt.Stringer
func (m MyState) String() string {
	return string(m)
}

// Well defined state names
const (
	READY   = MyState("ready")
	ERROR   = MyState("error")
	PENDING = MyState("pending")
)

// High-level example for an app
func Example() {
	// Create a new state
	mystate := NewState(StateOptions{
		Namespace: "my.app",
		System:    "http.server",
		Name:      "state",
	}, PENDING)

	// Register the new state to the default registry
	MustRegister(mystate)

	// later in the code
	mystate.Set(READY)

	//http.Handle("/states", stateshttp.Handler())
	// http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
