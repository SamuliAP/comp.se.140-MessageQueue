FROM golang
FROM golang

WORKDIR /app

COPY . .
WORKDIR /app/orig
RUN go get -d ./...
RUN go build -o out .

CMD ["/app/orig/out"]