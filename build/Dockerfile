FROM golang:alpine
LABEL maintainer="mfathoor.23@gmail.com" \
      name="mfathoor/pemilu-fraud-2024" \
      github="https://github.com/fathoor/pemilu-fraud-2024" \
      dockerhub="https://hub.docker.com/r/mfathoor/pemilu-fraud-2024"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o fraud-checker ./cmd/main.go

EXPOSE 2024

ENTRYPOINT ["/app/fraud-checker"]
