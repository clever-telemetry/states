package http

import (
	"fmt"
	"net/http"

	"github.com/clever-telemetry/states"
)

// Handler return an HTTP handler using the default Registry
func Handler() http.Handler {
	return HandlerFor(states.DefaultRegistry)
}

// HandlerFor return an http handler, using the given Registry
func HandlerFor(registry states.Registry) http.Handler {
	h := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		gtsList := registry.Gather()
		for i := range gtsList {
			for j := range gtsList[i].Values {
				if len(gtsList[i].Values[j]) > 1 {
					gtsList[i].Values[j] = gtsList[i].Values[j][len(gtsList[i].Values[j])-1:]
				}
			}
		}

		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusOK)

		res.Write([]byte(fmt.Sprintf("# %d states\n", len(gtsList))))

		_, _ = res.Write([]byte(gtsList.Sensision()))
	})
	return h
}
