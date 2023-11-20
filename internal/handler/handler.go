package handler

import (
	"fmt"
	"kolesagpt/config"
	"kolesagpt/internal/api"
	"kolesagpt/internal/models"
	"log"
	"net/http"
	"strings"

	"gopkg.in/olahol/melody.v1"
)

type Handler struct {
	Client *http.Client
	Cfg    *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		Cfg:    cfg,
		Client: &http.Client{},
	}
}

func (h *Handler) StartChat(userdata *models.User, s *melody.Session) {
	log.Println("starting chat for", userdata.Username)

	s.Write([]byte(userdata.Username + ": Let's simulate a real life interview for " + userdata.DesiredPosition))
	gptRes, err := api.PromptGPT(h.Client, h.Cfg, userdata)
	if err != nil {
		s.Write([]byte("AI: Sorry, I am not availabe at the moment. Please, try again later..."))
		return
	}
	s.Write([]byte("AI: " + gptRes.Choices[len(gptRes.Choices)-1].Message.Content))
	s.Set("user", userdata)
}

func (h *Handler) Chat(s *melody.Session, r *models.UserResponse) {
	userdataIn, ok := s.Keys["user"]
	if !ok {
		s.Write([]byte("AI: " + "couldn't retrieve userdata history from web socket session"))
		return
	}
	userdata := userdataIn.(models.User)
	if userdata.Username != r.Username {
		s.Write([]byte("AI: " + "Wait... Who are you?"))
		return
	}
	userdata.Messages = append(userdata.Messages, models.Message{
		Role:    "user",
		Content: r.Content,
	})
	gptRes, err := api.PromptGPT(h.Client, h.Cfg, &userdata)
	if err != nil {
		s.Write([]byte("AI: " + "Couldn't prompt api"))
		return
	}
	s.Set("user", append(userdata.Messages, models.Message{
		Role:    "user",
		Content: r.Content,
	}))
	s.Write([]byte("AI: " + gptRes.Choices[0].Message.Content))
}

func (h *Handler) HandleClose(s *melody.Session, i int, ss string) error {
	u, ok := s.Get("user")
	if !ok {
		s.Write([]byte("Couldn't get user from ws context"))
		return s.Close()
	}
	userdata := u.(*models.User)
	s.Write([]byte("Closing chat with:" + userdata.Username))
	log.Println("Closing chat with: " + userdata.Username)
	return s.Close()
}

func (h *Handler) HandleConnect(s *melody.Session) {
	log.Println("Starting chat with: " + s.Request.RemoteAddr)
}

func (h *Handler) HandleDisconnect(s *melody.Session) {
	u, _ := s.Get("user")
	userdata := u.(*models.User)
	s.Write([]byte("Disconnected: " + userdata.Username))
	log.Println("Disconnected: " + userdata.Username)
}

func (h *Handler) HandleMessage(s *melody.Session, msg []byte) {
	m := string(msg)
	wsdata := strings.Split(m, ",")
	log.Println("Got:", m)
	typeOfMessage := wsdata[0]
	wsdata = wsdata[1:]
	switch typeOfMessage {
	case "chat":
		r := &models.UserResponse{}
		r.Username = wsdata[0]
		r.Content = wsdata[1]
		h.Chat(s, r)
	case "start":
		userdata := &models.User{
			Username:                    wsdata[0],
			YearsOfCommercialExperience: wsdata[1],
			CurrentPosition:             wsdata[2],
			DesiredPosition:             wsdata[3],
			Stack:                       wsdata[4],
		}
		userdata.Messages = []models.Message{
			{
				Role:    "system",
				Content: "You are a technical interviewer.",
			},
			{
				Role: "user",
				Content: fmt.Sprintf(`I need you to act like a technical interviwer for 
				%s %s that has %s years of commercial exprerience. 
				Stimulate a real-life interview skipping the introductory part focusing only on verifying technical skills. 
				Start right now.",
				`, "Senior", "Golang Developer"),
			},
		}
		h.StartChat(userdata, s)
	}
}
