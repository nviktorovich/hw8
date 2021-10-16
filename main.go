package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather/configuration"
)


func main()  {

	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	options := configuration.GetFlags()



	q := req.URL.Query()
	q.Add("q", options.Town)
	q.Add("appid", configuration.ApiKey)
	q.Add("units", options.Units)

	req.URL.RawQuery = q.Encode()
	resp, _ := http.Get(req.URL.String())
	text, _ := io.ReadAll(resp.Body)

	if options.Logger {
		log.Println(string(text))
	} else {
		fmt.Println(string(text))
	}

}