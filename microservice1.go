package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/parnurzeal/gorequest"
)

func uploadRAML(c echo.Context) error {

	fmt.Println("Handler for /v1/uploadRAML called.")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	fileContentsString := string(fileContents[:])

	request := gorequest.New()

	url := fmt.Sprintf("http://ms2:%d/v1/verifyRAML", _configuration.Ms2Port)

	_, body, errs := request.Post(url).
		Send(fileContentsString).
		End()

	if errs != nil {
		return errs[0]
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<html><body>Result for file submitted on %s:<br><br><pre>%s</pre></body></html>", time.Now().Format(time.RFC850), body))
}

func showWelcomePage(c echo.Context) error {
	fmt.Println("Handler for /v1/index called.")

	return c.HTML(
		http.StatusOK,
		`<html>
			<body>
				Welcome to this ugly form. Please upload a RAML file.<br><br>
				<form action="/v1/uploadRAML" method="post" enctype="multipart/form-data">
		    			<div>
		        			<label for="file">File:</label>
		    				<input type="file" name="file"><br><br>
	    					<input type="submit" value="Submit">
					</div>
				</form>
			</body>
		</html>`)
}

type Config struct {
	Ms1Port int `env:"PORT" envDefault:"3000"`
	Ms2Port int `env:"MS2PORT" envDefault:"3000"`
}

// Global configuration
var _configuration = Config{}

func main() {

	env.Parse(&_configuration)

	fmt.Println("Configuration: ", _configuration)

	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Handler for / called.")
		return c.Redirect(302, "/v1/index")
	})

	e.GET("/v1/index", showWelcomePage)

	e.POST("/v1/uploadRAML", uploadRAML)

	fmt.Printf("Microservice 1 is now listening on port %d...", _configuration.Ms1Port)

	e.Run(standard.New(fmt.Sprintf(":%d", _configuration.Ms1Port)))
}
