package handler

import (
	"fmt"
	"kolesagpt/ui"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/olahol/melody.v1"
)

func (h *Handler) Routes() http.Handler {
	r := httprouter.New()
	ws := melody.New()
	r.NotFound = http.HandlerFunc(notFoundResponse)
	r.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	r.ServeFiles("/static/*filepath", http.Dir("./ui/static"))

	r.HandlerFunc(http.MethodGet, "/", HTML("home"))
	r.HandlerFunc(http.MethodGet, "/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	}))

	ws.HandleClose(h.HandleClose)
	ws.HandleConnect(h.HandleConnect)
	ws.HandleMessage(h.HandleMessage)
	ws.HandleDisconnect(h.HandleDisconnect)

	return enableCORS(r)
}

func HTML(templateName string, data ...any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ui.RenderTemplate(w, r, templateName, data)
	}
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "<h1>404</h1><hr>The requested resource could not be found"
	fmt.Fprint(w, http.StatusNotFound, message)
}
func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("<h1>405</h1><hr>The %s method is not supported for this resource", r.Method)
	fmt.Fprint(w, http.StatusMethodNotAllowed, message)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
