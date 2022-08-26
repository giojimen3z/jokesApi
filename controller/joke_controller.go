package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"

	"awesomeProject1/model"
	"awesomeProject1/service"
)

const (
	totalJokes = 25
)

func GetJokes() gin.HandlerFunc {
	return func(context *gin.Context) {
		jokes := make([]model.ObjectJokes, 0)

		getJokes(&jokes)

		context.JSONP(http.StatusOK, jokes)
	}
}

func getJokes(jokes *[]model.ObjectJokes) {
	wg := sync.WaitGroup{}

	for len(*jokes) < totalJokes {

		distinctJokes := totalJokes - len(*jokes)

		wg.Add(distinctJokes)

		for i := 0; i < distinctJokes; i++ {

			go getJoke(&wg, jokes)
		}
		wg.Wait()
	}

}

func getJoke(wg *sync.WaitGroup, jokes *[]model.ObjectJokes) {

	joke := service.GeJoke()

	if !isDifferent(joke, *jokes) {
		*jokes = append(*jokes, joke)
	}

	wg.Done()

}

func isDifferent(joke model.ObjectJokes, jokes []model.ObjectJokes) bool {

	for _, value := range jokes {

		if value.ID == joke.ID {

			return true
		}

	}
	return false
}
