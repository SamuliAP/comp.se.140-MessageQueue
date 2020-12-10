FROM golang

WORKDIR /app

COPY . .

RUN go get -u golang.org/x/lint/golint

CMD ["golint", "./..."]