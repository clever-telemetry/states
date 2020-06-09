package states

import (
	"reflect"
	"testing"

	warp10 "github.com/miton18/go-warp10/base"
)

func TestRegister(t *testing.T) {
	type args struct {
		state State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "main",
		args: args{
			NewState(StateOptions{Name: "mystate"}, nil),
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Register(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGather(t *testing.T) {
	DefaultRegistry = NewRegistry()

	tests := []struct {
		name string
		want warp10.GTSList
	}{{
		name: "main",
		want: warp10.GTSList{},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gather()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gather() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustRegister(t *testing.T) {
	type args struct {
		state State
	}
	tests := []struct {
		name string
		args args
	}{{
		name: "main",
		args: args{
			state: NewState(StateOptions{Name: "mystate"}, nil),
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustRegister(tt.args.state)
		})
	}
}
