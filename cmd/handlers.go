package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "gopkg.in/olahol/melody.v1"
)

func handleChat() {

	client := &http.Client{}
	fmt.Println("Enter message:")
	for {
		requestBody, err := GetRequestBodyString(turbo, messsageContent)
		if err != nil {
			panic(err)
		}

		PrintDebug("request body:" + requestBody)

		resp, err := SendRequest(context.TODO(), url, *client, requestBody)
		if err != nil {
			logERROR.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Something went wrong")
			fmt.Println("Response status:", resp.Body)
		}
		PrintDebug("responce body:" + string(body))

		res := gptRes{}
		err = json.Unmarshal(body, &res)
		if err != nil {
			panic(err)
		}

		if resp.Status != "200 OK" {
			logINFO.Println("Status:", resp.Status)
			fmt.Println("Body:", string(body))
			return
		}

		if len(res.Choices) > 0 {
			fmt.Println(res.Choices[0].Message.Content)
		}

	}
}
