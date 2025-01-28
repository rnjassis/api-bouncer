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
	// TODO error when len(project.Requests) == 0
	for _, request := range project.Requests {
		err := routeFactory(r, request)
		if err != nil {
			fmt.Println(err, ", url will return 404")
		}
	}
	r.Run(project.Port)
}

func routeFactory(ginEngine *gin.Engine, request models.Request) error {
	var err error

	switch request.RequestMethod {
	case models.GET:
		err = getRoute(ginEngine, request)
	case models.POST:
		err = postRoute(ginEngine, request)
	default:
		return errors.New("No method found for " + string(request.RequestMethod))
	}

	if err != nil {
		return errors.New("Error creating route - " + err.Error())
	}
	return nil
}

func validateOneResponse(request models.Request) (*models.Response, error) {
	if len(request.Responses) > 1 {
		return nil, fmt.Errorf("too many responses for the request %s", request.Url)
	} else if len(request.Responses) == 0 {
		return nil, fmt.Errorf("no responses for the request %s", request.Url)
	} else {
		return &request.Responses[0], nil
	}
}
func getRoute(ginEngine *gin.Engine, request models.Request) error {
	response, err := validateOneResponse(request)
	if err != nil {
		return err
	}
	ginEngine.GET(request.Url, func(c *gin.Context) {
		getResponse(c, response)
	})

	return nil
}

func postRoute(ginEngine *gin.Engine, request models.Request) error {
	response, err := validateOneResponse(request)
	if err != nil {
		return err
	}
	ginEngine.POST(request.Url, func(c *gin.Context) {
		getResponse(c, response)
	})
	return nil
}

func getResponse(ginContext *gin.Context, response *models.Response) {
	ginContext.Data(response.StatusCode, response.Mime, []byte(response.Body))
}
