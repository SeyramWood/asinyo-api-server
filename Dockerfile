FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go build -o main ./cmd/
COPY . ./
EXPOSE 9000
CMD ["./main"]