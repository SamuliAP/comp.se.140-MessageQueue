FROM golang

WORKDIR /app

COPY . .
WORKDIR /app/server

RUN go build -o out .

CMD ["/app/server/out"]