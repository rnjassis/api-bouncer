package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rnjassis/api-bouncer/models"
)

func RunServer(project *models.Project) {
	fmt.Println("Starting server ", project.Name)
	r := gin.Default()
	for _, request := range project.Requests {
		routeFactory(r, request)
	}
	r.Run(project.Port)
}

func routeFactory(ginEngine *gin.Engine, request models.Request) error {
	var err error

	switch request.RequestMethod {
	case models.GET:
		err = getRoute(ginEngine, request)
	default:
		return errors.New("No method found for " + string(request.RequestMethod))
	}

	if err != nil {
		return errors.New("Error creating route - " + err.Error())
	}
	return nil
}

func getRoute(ginEngine *gin.Engine, request models.Request) error {
	var activeResponse *models.Response
	for _, response := range request.Responses {
		if response.Active {
			activeResponse = &response
			break
		}
	}
	if activeResponse == nil {
		return errors.New("active response not found")
	}

	ginEngine.GET(request.Url, func(c *gin.Context) {
		getResponse(c, activeResponse)
	})
	return nil
}

func getResponse(ginContext *gin.Context, response *models.Response) {
	ginContext.Data(response.StatusCode, response.Mime, []byte(response.Body))
}
