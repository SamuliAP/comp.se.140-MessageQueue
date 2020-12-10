# comp.se.140-MessageQueue

### Running the application services
`docker-compose up -d --build`

### Running the linter
`docker-compose -f docker-compose-lint.yml up --build `

### Running the tests
`docker-compose -f docker-compose-tests.yml up --build`

`docker-compose -f docker-compose-tests-shutdown.yml up --build`

Note that running the shutdown test suite will shut down all of the services (excluding RabbitMQ).