FROM golang:latest

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]