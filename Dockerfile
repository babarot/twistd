FROM golang:latest
RUN mkdir /app
ADD ./twist.linux-amd64 /app/twistd
ADD ./config.toml /app
WORKDIR /app
CMD ["/app/twistd", "-c", "/app/config.toml"]
