FROM golang:alpine
RUN mkdir /app
ADD .. /app/
WORKDIR /app/
COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download && go mod verify
COPY .. ./
RUN go build -o main ./cmd/
EXPOSE 9000
CMD ["./main"]
