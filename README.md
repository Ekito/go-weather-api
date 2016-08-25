# go-weather-api

A small Go app for serving geocoding & weather

Based on echo http stack : https://echo.labstack.com/
Make reaquest on google maps geocoding service and wunderground weather service

- geocode service : /geocode?address=<> i.e : /geocode?address=Toulouse,France
- weather service : /weather?lat= & lon= &lang= i.e: /weather?lat=43.604652&lon=1.444209&lang=FR

## requirements
Go SDK installed (tested in version 1.7 here)
Glide installed : see https://github.com/Masterminds/glide

## setup environment
GOPATH must be set to the root project path.

Go to the /src/weather_api
run ```glide install```(it will install dependencies in vendor folder)

## make / run
from project folder, run ```go build ./src/weather_api/main.go```
or ```go run ./src/weather_api/main.go```

## environment keys
- PORT : your http port (needed for heroku)
- GEOCODE_KEY : your google map api key
- WEATHER_KEY : your wunderground api key

## docker image
make your docker image : ```docker build -t weather-api .```
(run it with the needed environment variables - https://devcenter.heroku.com/articles/container-registry-and-runtime)


## deploy on heroku
Edit .env local file in your root project folder, and write your environment keys (https://devcenter.heroku.com/articles/heroku-local#set-up-your-local-environment-variables)

Follow the commands below:
- heroku login
- heroku git:remote -a <your-app> (link local project to your app)
- heroku labs:enable log-runtime-metrics (enable metrics)
- heroku container:push web (push the docker image)
- heroku open / heroku logs (see what's going on)

Feedback and suggestion are welcome ! :)
