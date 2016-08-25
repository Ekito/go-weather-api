package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"os"
	"fmt"
	"io/ioutil"
	"net/url"
	"github.com/labstack/echo/middleware"
	"time"
)

const SEPERATOR string = "_"

var GEOCODE_KEY string
var WEATHER_KEY string

var dynCache map[string]string
var staticCache map[string]string

func main() {
	// get environment keys
	GEOCODE_KEY = os.Getenv("GEOCODE_KEY")
	if GEOCODE_KEY == "" {
		fmt.Println("ERROR :: can't find GEOCODE_KEY variable")
		return
	}
	WEATHER_KEY = os.Getenv("WEATHER_KEY")
	if WEATHER_KEY == "" {
		fmt.Println("ERROR :: can't find WEATHER_KEY variable")
		return
	}
	// set cache
	staticCache = make(map[string]string)
	dynCache = make(map[string]string)
	// run clean cache routine
	go cacheClean()

	// run http engine
	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is working ! :)")
	})
	e.GET("/geocode", geocodeHandler)
	e.GET("/weather", weatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("using default port ...")
		port = "8080"
	} else {
		fmt.Printf("using port %s\n", port)
	}
	fmt.Println("Running ...")
	e.Run(standard.New(":" + port))
}

func cacheClean() {
	for {
		time.Sleep(time.Hour * 1)
		dynCache = make(map[string]string)
		fmt.Println("clean dyn cache !")
	}
}

func geocodeHandler(c echo.Context) error {
	address := c.QueryParam("address")
	var result string
	if staticCache[address] == "" {
		resp, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + url.QueryEscape(address) + "&key=" + GEOCODE_KEY)
		if err != nil {
			fmt.Printf("got error : %s", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		result = string(body)
		staticCache[address] = result
	} else {
		result = staticCache[address]
	}
	return c.String(http.StatusOK, result)
}

func weatherHandler(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	lang := c.QueryParam("lang")
	key := make_key(lat, lon, lang)

	var result string
	if dynCache[key] == "" {
		resp, err := http.Get("http://api.wunderground.com/api/" + WEATHER_KEY + "/forecast/lang:" + lang + "/pws:0/q/" + lat + "," + lon + ".json")
		if err != nil {
			fmt.Printf("got error : %s", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		result = string(body)
		dynCache[key] = result
	} else {
		result = dynCache[key]
	}
	return c.String(http.StatusOK, result)
}

func make_key(lat, lon, lang string) string {
	return lat + SEPERATOR + lon + SEPERATOR + lang
}