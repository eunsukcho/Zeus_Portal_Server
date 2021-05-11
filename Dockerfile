FROM golang:latest
RUN apt install ca-certificates
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
