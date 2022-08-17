# The base go-image
FROM golang:1.14-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o crawler ./main.go

# Run the server executable
CMD [ "/app/crawler" ]