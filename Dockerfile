FROM golang:1.20.5-alpine as base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

FROM base as dev
RUN go install github.com/cosmtrek/air@latest;
COPY . .
EXPOSE 9000
CMD ["sh", "-c", "air"]

FROM base as prod
RUN go build -o /cmd/build ./cmd/
EXPOSE 9000
CMD ["sh", "-c", "/cmd/build"]
