package states

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	warp10 "github.com/miton18/go-warp10/base"
)

type S string

func (s S) String() string {
	return string(s)
}

func TestNewState(t *testing.T) {
	type args struct {
		options      StateOptions
		initialValue fmt.Stringer
	}
	tests := []struct {
		name string
		args args
		want State
	}{{
		name: "main",
		args: args{
			options: StateOptions{
				Namespace: "app",
				System:    "sys",
				Name:      "state",
			},
			initialValue: S("READY"),
		},
		want: &iState{
			options: StateOptions{
				Namespace: "app",
				System:    "sys",
				Name:      "state",
			},
			separator: DefaultSeparator,
			value: "READY",
		},
	}, {
		name: "with init value",
		args: args{
			options: StateOptions{
				Namespace: "app",
				System:    "sys",
				Name:      "state",
			},
			initialValue: nil,
		},
		want: &iState{
			options: StateOptions{
				Namespace: "app",
				System:    "sys",
				Name:      "state",
			},
			separator: DefaultSeparator,
			value: DefaultStateValue.String(),
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewState(tt.args.options, tt.args.initialValue)

			if got.Help() != tt.want.Help() {
				t.Errorf(
					"NewState() Help = %+v, want %+v",
					got.Help(),
					tt.want.Help(),
				)
			}

			if got.String() != tt.want.String() {
				t.Errorf(
					"NewState() string = %+v, want %+v",
					got.String(),
					tt.want.String(),
				)
			}

			if got.Metric().ClassName != tt.want.Metric().ClassName {
				t.Errorf(
					"NewState() metric.class = %+v, want %+v",
					got.Metric().ClassName,
					tt.want.Metric().ClassName,
				)
			}

			if got.Metric().SensisionSelector(true) != tt.want.Metric().SensisionSelector(true) {
				t.Errorf(
					"NewState() metric.class = %+v, want %+v",
					got.Metric().SensisionSelector(true),
					tt.want.Metric().SensisionSelector(true),
				)
			}

			if !reflect.DeepEqual(got.Metric().Values[0][1], tt.want.Metric().Values[0][1]) {
				t.Errorf(
					"NewState() metric.values.v = %+v, want %+v",
					got.Metric().Values,
					tt.want.Metric().Values,
				)
			}
		})
	}
}

//nolint:govet
func Test_iState_Metric(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		options StateOptions
		value   string
		time    time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *warp10.GTS
	}{{
		name: "main",
		fields: fields{
			options: StateOptions{Name: "mystate"},
			value:   DefaultStateValue.String(),
		},
		want: &warp10.GTS{
			ClassName: "mystate",
			Values: warp10.Datapoints{{
				0, "UNKNOWN",
			}},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &iState{
				options: tt.fields.options,
				value:   tt.fields.value,
				time:    tt.fields.time,
			}

			got := state.Metric()

			if got.SensisionSelector(true) != tt.want.SensisionSelector(true) {
				t.Errorf("iState.Metric() selector = %v, want %v", got, tt.want)
			}

			if len(tt.want.Values) == 0 {
				return
			}

			if got.Values[0][1] != tt.want.Values[0][1] {
				t.Errorf("iState.Metric() selector = %v, want %v", got, tt.want)
			}
		})
	}
}
