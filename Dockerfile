FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go build -o main .
EXPOSE 2000

CMD [ "/app/main" ]

