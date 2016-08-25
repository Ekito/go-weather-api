# Start from a Debian image with the latest version of Go installed
#FROM golang
# extend version with glide
FROM arnaudgiuliani/golang-glide

# Copy the local package files to the container's workspace.
ADD /src/weather_api/main.go /go/src/app/main.go
ADD /src/weather_api/glide.yaml /go/src/app/glide.yaml

WORKDIR /go/src/app

# get dependencies
RUN glide install

ENV GOBIN /go/bin
# make bin
RUN go install main.go

# Add user (for heroku)
RUN useradd -m myuser
USER myuser

# Heroku provide PORT environment variable 
# EXPOSE 8080

# GEOCODE_KEY : google maps api key
# WEATHER_KEY : wundeground api key

CMD /go/bin/main