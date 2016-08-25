# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
#FROM golang
FROM golang_glide

# Copy the local package files to the container's workspace.
ADD /src/weather_api/main.go /go/src/app/main.go
ADD /src/weather_api/glide.yaml /go/src/app/glide.yaml

WORKDIR /go/src/app

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN glide install

ENV GOBIN /go/bin
RUN go install main.go

# Run the outyet command by default when the container starts.
RUN useradd -m myuser
USER myuser

# Document that the service listens on port 8080.
# EXPOSE 8080

CMD /go/bin/main