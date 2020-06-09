package states

import (
	"sync"
	"testing"
	"time"

	warp10 "github.com/miton18/go-warp10/base"
)

//nolint:govet
func Test_iRegistry_Register(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		states  States
	}
	type args struct {
		state State
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{{
		name: "main",
		fields: fields{
			states: States{},
		},
		args: args{
			NewState(StateOptions{Name: "mystate"}, nil),
		},
		wantErr: false,
	}, {
		name: "collistion",
		fields: fields{
			states: States{
				NewState(StateOptions{Name: "mystate"}, nil),
			},
		},
		args: args{
			NewState(StateOptions{Name: "mystate"}, nil),
		},
		wantErr: true,
	}, {
		name: "no name",
		fields: fields{
			states: States{},
		},
		args: args{
			NewState(StateOptions{Name: ""}, nil),
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := &iRegistry{
				RWMutex: tt.fields.RWMutex,
				states:  tt.fields.states,
			}

			err := registry.Register(tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("iRegistry.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//nolint:govet
func Test_iRegistry_MustRegister(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		states  States
	}
	type args struct {
		state State
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectPanic bool
	}{{
		name: "main",
		fields: fields{
			states: States{},
		},
		args: args{
			state: NewState(StateOptions{Name: "mystate"}, nil),
		},
		expectPanic: false,
	}, {
		name: "no name",
		fields: fields{
			states: States{},
		},
		args: args{
			state: NewState(StateOptions{Name: ""}, nil),
		},
		expectPanic: true,
	}}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			registry := &iRegistry{
				RWMutex: tt.fields.RWMutex,
				states:  tt.fields.states,
			}
			registry.MustRegister(tt.args.state)
		})
	}
}

//nolint:govet
func Test_iRegistry_Gather(t *testing.T) {
	testGts := warp10.NewGTS("mystate")
	testGts.Values.Add(time.Now(), "UNKNOWN")

	type fields struct {
		RWMutex sync.RWMutex
		states  States
	}
	tests := []struct {
		name   string
		fields fields
		want   warp10.GTSList
	}{{
		name: "main",
		fields: fields{
			states: States{},
		},
		want: warp10.GTSList{},
	}, {
		name: "main",
		fields: fields{
			states: States{
				NewState(StateOptions{Name: "mystate"}, nil),
			},
		},
		want: warp10.GTSList{testGts},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := &iRegistry{
				RWMutex: tt.fields.RWMutex,
				states:  tt.fields.states,
			}

			got := registry.Gather()

			if len(got) != len(tt.want) {
				t.Errorf("iRegistry.Gather() len = %d, want %d", len(got), len(tt.want))
			}

			if len(tt.want) == 0 {
				return
			}

			if got[0].ClassName != tt.want[0].ClassName {
				t.Errorf(
					"iRegistry.Gather() .class = %v, want %v",
					got[0].ClassName,
					tt.want[0].ClassName,
				)
			}

			if got[0].SensisionSelector(true) != tt.want[0].SensisionSelector(true) {
				t.Errorf(
					"iRegistry.Gather() .class = %v, want %v",
					got[0].SensisionSelector(true),
					tt.want[0].SensisionSelector(true),
				)
			}

			if got[0].Values[0][1] != tt.want[0].Values[0][1] {
				t.Errorf(
					"iRegistry.Gather() .value = %v, want %v",
					got[0].Values[0],
					tt.want[0].Values[0],
				)
			}
		})
	}
}
