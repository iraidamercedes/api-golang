FROM golang:1.14

# Installl dependencies
RUN apt-get update
RUN apt-get install whois
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valyala/fasthttp

# Use a specific port for the app
EXPOSE 8000