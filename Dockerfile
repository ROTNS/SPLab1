#Lab1 SP
#9 var
#Debian
#Golang

FROM debian 

MAINTAINER KuznetsovDmitriy

WORKDIR /home

COPY main.go /home/main.go

RUN apt-get update
RUN apt-get install -y golang

CMD go run main.go