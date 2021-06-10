FROM ubuntu
RUN apt-get update
RUN apt install -y ca-certificates
RUN update-ca-certificates
ADD ./certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
