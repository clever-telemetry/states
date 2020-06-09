package states

import (
	"sync"

	warp10 "github.com/miton18/go-warp10/base"
	"github.com/pkg/errors"
)

type iRegistry struct {
	sync.RWMutex
	states States
}

// NewRegistry create a new Registry
func NewRegistry() Registry {
	return &iRegistry{
		states: States{},
	}
}

func (registry *iRegistry) Register(state State) error {
	if state.Metric().ClassName == "" {
		return errors.Errorf("at least a name is required to define a state")
	}

	registry.Lock()
	defer registry.Unlock()

	for _, stt := range registry.states {
		s1 := stt.Metric().SensisionSelector(false)
		s2 := state.Metric().SensisionSelector(false)
		if s1 == s2 {
			return AlreadyExistsErr
		}
	}

	registry.states = append(registry.states, state)

	return nil
}

func (registry *iRegistry) MustRegister(state State) {
	err := registry.Register(state)
	if err != nil {
		panic(errors.Wrapf(err, "cannot register state (%s)", state.String()))
	}
}

func (registry *iRegistry) Gather() warp10.GTSList {
	registry.RLock()
	defer registry.RUnlock()

	gts := make(warp10.GTSList, len(registry.states))
	for i, state := range registry.states {
		gts[i] = state.Metric()
	}

	return gts
}
