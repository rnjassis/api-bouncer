package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/rnjassis/api-bouncer/models"
)

func RunServer(project *models.Project) {
	fmt.Println("Starting server ", project.Name)
	r := gin.Default()
	for _, request := range project.Requests {
		err := routeFactory(r, request)
		if err != nil {
			fmt.Println(err, ", url will return 404")
		}
	}
	r.Run(project.Port)
}

func routeFactory(ginEngine *gin.Engine, request models.Request) error {
	response, err := getOneResponse(request)

	if err != nil {
		ginEngine.Any(request.Url, func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error while creating non specified method"})
		})
		return errors.New("Error creating route - " + err.Error())
	}

	routeByMethod(&request, response, ginEngine)

	return nil
}

func routeByMethod(request *models.Request, response *models.Response, ginEngine *gin.Engine) error {
	switch request.RequestMethod {
	case models.GET:
		ginEngine.GET(request.Url, getMethodHandlerFunc(request, response))
	case models.POST:
		ginEngine.POST(request.Url, getMethodHandlerFunc(request, response))
	case models.PUT:
		ginEngine.PUT(request.Url, getMethodHandlerFunc(request, response))
	case models.DELETE:
		ginEngine.DELETE(request.Url, getMethodHandlerFunc(request, response))
	case models.OPTIONS:
		ginEngine.OPTIONS(request.Url, getMethodHandlerFunc(request, response))
	case models.PATCH:
		ginEngine.PATCH(request.Url, getMethodHandlerFunc(request, response))
	case models.ANY:
		ginEngine.Any(request.Url, getMethodHandlerFunc(request, response))
	}
	return nil
}

func getMethodHandlerFunc(request *models.Request, response *models.Response) gin.HandlerFunc {
	if response.Proxy {
		return proxyRoute(response.Body, string(request.RequestMethod))
	} else if response.Redirect {
		return func(c *gin.Context) {
			addHeaders(c, response.Headers)
			c.Redirect(http.StatusPermanentRedirect, response.Body) //response.Body has the new url
		}
	} else {
		return func(c *gin.Context) {
			addHeaders(c, response.Headers)
			c.Data(response.StatusCode, response.Mime, []byte(response.Body)) // response.Body has the new url
		}
	}
}

func getOneResponse(request models.Request) (*models.Response, error) {
	if len(request.Responses) > 1 {
		return nil, fmt.Errorf("too many responses for the request %s", request.Url)
	} else if len(request.Responses) == 0 {
		return nil, fmt.Errorf("no responses for the request %s", request.Url)
	} else {
		return &request.Responses[0], nil
	}
}

func proxyRoute(target string, method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error extracting the body"})
			return
		}

		req, err := http.NewRequest(method, target, bytes.NewReader(body))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error creating the redirect request"})
			return
		}

		// headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// query params
		params := c.Request.URL.Query()
		targetParsed, _ := url.Parse(target)
		targetQuery := targetParsed.Query()
		for key, values := range params {
			for _, value := range values {
				targetQuery.Add(key, value)
			}
		}
		targetParsed.RawQuery = targetQuery.Encode()
		req.URL = targetParsed

		// Send to the target
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error to send the request"})
			return
		}
		defer resp.Body.Close()

		// Read response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error to parse the response"})
			return
		}

		// Read response headers
		for key, values := range resp.Header {
			for _, value := range values {
				c.Writer.Header().Add(key, value)
			}
		}

		// Return
		c.Status(resp.StatusCode)
		c.Writer.Write(responseBody)
	}
}

func addHeaders(c *gin.Context, headers string) {
	if len(headers) > 0 {
		json_raw := make(map[string]json.RawMessage)
		err := json.Unmarshal([]byte(headers), &json_raw)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error extracting the headers"})
		}

		for key, value := range json_raw {
			c.Header(key, string(value))
		}

	}
}
