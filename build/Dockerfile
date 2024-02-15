FROM golang:alpine
LABEL maintainer="mfathoor.23@gmail.com" \
      name="mfathoor/fraud24" \
      github="https://github.com/fathoor/fraud24" \
      dockerhub="https://hub.docker.com/r/mfathoor/fraud24"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o fraud24 ./cmd/fraud24/main.go

EXPOSE 2024

ENTRYPOINT ["/app/fraud24"]
