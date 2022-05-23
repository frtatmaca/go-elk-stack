# Go ELK Stack
Go ELK Stack

## About

Simply, i designed, coded and implemented ELK stack in the project.


This project was developed by Go Language.

## Usage

You can directly run server.go file to use program.

Two endpoints are defined to use, these are;

```Shell
To write log to Elasticsearch via RabbitMQ

GET http://localhost:8000/writinglogtorabbitmq

```

```Shell
To write log to Elasticsearch via Filebeat and log file

GET http://localhost:8000/writinglogtofile

```

## Run

You can run the project via command:

```Shell
go run server.go
```

## Run services

Every service on ELK stack (Elasticsearch, Logstash, Kibana, RabbitMQ and Beats) needs a container, so on docker-compose.yml you can check every config.

You can dockerize the project via command:

```Shell
docker-compose up -d
```

## About me

I am Fırat Atmaca.

I have been working on software projects since 2013.

You can contact with me via [Linkedin](https://www.linkedin.com/in/firat-atmaca-469b2769/)

