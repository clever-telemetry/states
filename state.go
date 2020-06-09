package states

import (
	"fmt"
	"strings"
	"sync"
	"time"

	warp10 "github.com/miton18/go-warp10/base"
)

type iState struct {
	sync.RWMutex
	options StateOptions
	value   string
	time    time.Time
}

// NewState create a new state
func NewState(options StateOptions, initialValue fmt.Stringer) State {
	state := &iState{
		options: options,
		value:   DefaultStateValue.String(),
		time:    time.Now(),
	}

	if initialValue != nil {
		state.Set(initialValue)
	}

	return state
}

func (state *iState) Metric() *warp10.GTS {
	state.RLock()
	defer state.RUnlock()

	labels := warp10.Labels(state.options.Labels)
	gts := warp10.NewGTSWithLabels(state.key("."), labels)

	if state.options.Help != "" {
		gts.Attributes = warp10.Attributes{
			"help": state.options.Help,
		}
	}

	gts.Values.Add(state.time, state.value)

	return gts
}

func (state *iState) String() string {
	state.RLock()
	defer state.RUnlock()

	return fmt.Sprintf(
		"%s->%s",
		state.key("/"),
		state.value,
	)
}

func (state *iState) Set(newState fmt.Stringer) {
	state.Lock()
	defer state.Unlock()

	state.value = newState.String()
	state.time = time.Now()
}

func (state *iState) Help() string {
	state.RLock()
	defer state.RUnlock()

	return state.options.Help
}

func (state *iState) key(separator string) string {
	state.RLock()
	defer state.RUnlock()

	name := []string{}

	if state.options.Namespace != "" {
		name = append(name, state.options.Namespace)
	}
	if state.options.System != "" {
		name = append(name, state.options.System)
	}

	name = append(name, state.options.Name)

	return strings.Join(name, separator)
}

// DefaultStateValue return an unknown state
type defaultStateValue struct{}

func (d defaultStateValue) String() string {
	return "UNKNOWN"
}
