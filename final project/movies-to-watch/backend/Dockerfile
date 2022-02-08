FROM golang:alpine
WORKDIR /app
COPY main.go /app
COPY go.mod /app
RUN go build -o backend
CMD ["./backend"]