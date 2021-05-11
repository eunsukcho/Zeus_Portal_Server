FROM golang:latest
RUN apk --no-cache add ca-certificates 
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go-wrapper install -ldflags '-extldflags "-static'
FROM scratch
COPY --from-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from-build /go/bin/app /app
WORKDIR /tmp
ADD . /tmp
EXPOSE 3000
