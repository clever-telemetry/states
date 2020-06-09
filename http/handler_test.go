package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/clever-telemetry/states"
)

func TestHandlerFor(t *testing.T) {
	type args struct {
		states states.States
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "main",
		args: args{
			states: states.States{
				states.NewState(states.StateOptions{
					Name: "mystate",
				}, nil),
			},
		},
		want: "# 1 states\n// mystate{} 'UNKNOWN'\n\n",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HandlerFor(states.DefaultRegistry)
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			for _, st := range tt.args.states {
				states.MustRegister(st)
			}

			got.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("HandlerFor() status = %v, want %v", rr.Code, http.StatusOK)
			}

			s := rr.Body.String()
			if !strings.HasPrefix(s, tt.want) {
				t.Errorf("HandlerFor() = '%+v'", s)
			}

		})
	}
}
