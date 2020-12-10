FROM golang

WORKDIR /app

COPY . .

RUN go get -d ./...
RUN go build -o out ./obse

CMD ["/app/out"]