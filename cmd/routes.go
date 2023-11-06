package main

import (
	"fmt"
	"kolesagpt/ui"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/olahol/melody.v1"
)

type Userdata struct {
	username                    string
	YearsOfCommercialExperience string
	CurrentPosition             string
	DesiredPosition             string
	stack                       string
}

var (
	userdata = &Userdata{}
)

func routes() http.Handler {
	r := httprouter.New()
	m := melody.New()
	// Convert the notFoundResponse() helper to a http.Handler using the
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	r.NotFound = http.HandlerFunc(notFoundResponse)
	r.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	r.ServeFiles("/static/*filepath", http.Dir("./ui/static"))

	r.HandlerFunc(http.MethodGet, "/", HTML("home"))
	r.HandlerFunc(http.MethodGet, "/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	}))
	m.HandleClose(func(s *melody.Session, i int, ss string) error {
		logINFO.Println("Closing chat with:", userdata.username)
		return s.Close()
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m := string(msg)
		wsdata := strings.Split(m, ",")
		switch wsdata[0] {
		case "chat":

		case "start":
			userdata.username = wsdata[0]
			userdata.YearsOfCommercialExperience = wsdata[1]
			userdata.CurrentPosition = wsdata[2]
			userdata.DesiredPosition = wsdata[3]
			userdata.stack = wsdata[4]
			logINFO.Println("got:", userdata)
			// s.Write([]byte(userdata.username))
			startChat(userdata, s)
		}
	})

	return enableCORS(r)
}

func HTML(templateName string, data ...any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ui.RenderTemplate(w, r, templateName, data)
	}
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	fmt.Fprint(w, http.StatusNotFound, message)
}
func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
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
