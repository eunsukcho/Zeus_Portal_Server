FROM ubuntu:latest
ADD main /tmp/
WORKDIR /tmp/main
EXPOSE 3000
CMD ['main']
