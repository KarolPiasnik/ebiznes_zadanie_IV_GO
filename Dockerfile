FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /server_app

EXPOSE 1323
EXPOSE 5432

# Run
CMD ["/server_app"]