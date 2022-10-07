FROM golang:1.19.1-bullseye
WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
CMD ["go", "run", "main.go"]