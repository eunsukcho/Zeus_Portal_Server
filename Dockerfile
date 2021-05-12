FROM golang:latest
RUN apt install ca-certificates
ADD ./certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
