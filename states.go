package states

import (
	"fmt"

	warp10 "github.com/miton18/go-warp10/base"
)

type (

	// State give the current state of a system
	State interface {
		// Metric return a timeseries
		Metric() *warp10.GTS
		// String pretty print the namespace, system, name and value of the state
		String() string
		// Set update the current state
		Set(fmt.Stringer)
		// Print informations about this state
		Help() string
	}

	// States is a list of States
	States []State

	// StateOptions is used to declare a new State
	StateOptions struct {
		Namespace string `json:"namespace,omitempty"`
		System    string `json:"system,omitempty"`
		Name      string `json:"name"`
		Help      string `json:"help,omitempty"`
	}

	// Registry collect a set of app states
	Registry interface {
		// Register a new state to this registry
		// Can fail if another state with the same namespace/system/name exists
		Register(State) error
		// Register a new state to this registry
		// Panic if another state with the same namespace/system/name exists
		MustRegister(State)
		// Gather a list of all states timeseries
		Gather() warp10.GTSList
	}
)
