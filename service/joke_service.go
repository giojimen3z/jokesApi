package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"awesomeProject1/model"
)

func GeJoke() model.ObjectJokes {

	joke := model.ObjectJokes{}

	response, err := http.Get("https://api.chucknorris.io/jokes/random")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		log.Fatal(err)
	}

	json.Unmarshal(body, &joke)

	return joke
}
