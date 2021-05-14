FROM ubuntu:latest
ADD ./certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
CMD ["/bin/bash","-c","./main"]
