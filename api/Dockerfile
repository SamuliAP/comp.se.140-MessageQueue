FROM golang

WORKDIR /app

COPY ./api ./src

RUN go mod init api
RUN go build -o out ./src

CMD ["/app/out"]