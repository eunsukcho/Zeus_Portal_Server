FROM ubuntu
RUN apt install -y ca-certificates
ADD ./certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
