FROM golang

WORKDIR /app

COPY ./tests ./tests
COPY ./api ./api

CMD ["go", "test", "-v" ,"./tests/shutdown"]