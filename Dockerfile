FROM docker.io/golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init sandbox-go
RUN go build -o  .
EXPOSE 8080
CMD ["/app/sandbox-go"]