package main

import (
	"fmt"

	"gopkg.in/olahol/melody.v1"
	_ "gopkg.in/olahol/melody.v1"
)

var (
	name                        = "Ильяс"
	YearsOfCommercialExperience = 1
	CurrentPosition             = "Junior"
	DesiredPosition             = "Middle"
	AmountOfQuestions           = 1
	messsageContent             = "I want you to act as an interviewer. I will be the candidate and you will ask me coding interview questions for a Junior Software Engineer position. I want you to only reply as the interviewer. Do not write all the conversation at once. I want you to only do the interview with me. Ask me the questions and wait for my answers. Do not write explanations. Ask me the questions one by one like an interviewer does and wait for my answers. Ask me random questions from one of the following topics and ask follow-up questions: Data Structure, Algorithm, Operating System, System Design, Network and Security"

	// messsageContent = fmt.Sprintf("I want you to act as a career consultant. "+
	// 	"My name is %s. "+
	// 	"My current position is %s with %v year of experience. "+
	// 	"I want to grow from %s level and you will as me questions as for %s"+
	// 	"I want you to ask me questions one at a time, wait for an answer and only then ask the next question. "+
	// 	"Count every question and once you will recieve answers for %v questions give your assessment and recommendation! "+
	// 	"I warn you, you will continue the dialoge with the client. "+
	// 	"Start the interview with "+
	// 	"'Привет, %s. Ты бы хотел подняться с уровня %s до %s, верно?'"+
	// 	"Strictly ask only one question at a time!"+
	// 	"answer in russian",
	// 	name, CurrentPosition, YearsOfCommercialExperience, DesiredPosition, DesiredPosition, AmountOfQuestions, name, CurrentPosition, DesiredPosition,
	// )
)

func startChat(userdata *Userdata, s *melody.Session) {
	logINFO.Println("starting chat for", userdata.username)

	err := s.Write([]byte(fmt.Sprintf(`
	{
		"username": "%s",
		"content": "%s"
	}`, userdata.username, "Hi from server")))
	panic(err)
	// client := &http.Client{}
	// 	go func(userdata *Userdata) {
	// 		for {
	// 			requestBody, err := GetRequestBodyString(turbo, messsageContent)
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			PrintDebug("request body:" + requestBody)

	// 			resp, err := SendRequest(context.TODO(), url, *client, requestBody)
	// 			if err != nil {
	// 				logERROR.Println(err)
	// 				return
	// 			}
	// 			defer resp.Body.Close()

	// 			body, err := ioutil.ReadAll(resp.Body)
	// 			if err != nil {
	// 				log.Println("Something went wrong")
	// 				fmt.Println("Response status:", resp.Body)
	// 			}

	// 			PrintDebug("responce body:" + string(body))

	// 			res := gptRes{}
	// 			err = json.Unmarshal(body, &res)
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			if resp.Status != "200 OK" {
	// 				logINFO.Println("Status:", resp.Status)
	// 				fmt.Println("Body:", string(body))
	// 				return
	// 			}

	// 			// send msg to client
	// 			if len(res.Choices) > 0 {
	// 				fmt.Println(res.Choices[0].Message.Content)
	// 			}

	//			// recieve message from
	//		}
	//	}(userdata)
}
