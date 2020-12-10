FROM golang

WORKDIR /app

COPY . .
WORKDIR /app/imed

RUN go get -d ./...
RUN go build -o out .

CMD ["/app/imed/out"]