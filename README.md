# Truphone Assessment

## Introduction

This assessment was developed using Golang. I think that with go the assessment could be done faster, because during this week I had my driver's license test, had school work (cybersecurity master's degree), and my day to day normal job. I would be able to do this
with Java, this decision was only because of the time.

I hope you enjoy reviewing this work. 

### Contents
* Sorting
* Building
* Running
* Testing
* Postman
* Swagger

## Sorting

The sorting criteria was the number of members of a family.

## Building

Running the service requires to build the docker image locally. 
To do this just run the command:

```shell script
make build_app
```

There you go, you now have truphone_go:latest docker image.

## Running

To run the service, as well its database just run:

```shell script
make run
```

This will run one instance of the postgres and one of the truphone_go.

## Testing

### Unit Testing

To run the unit tests and see the coverage (76.2% of files and 85.3% of statements with 485 tests) just run:

```shell script
make coverage
```

This will run the unit tests and will open in the browser the coverage report.

## Postman

A postman collection included in the file: `Truphone.postman_collection.json`

## Swagger

A swagger documentation of the API included in the file: `swagger.yml`

