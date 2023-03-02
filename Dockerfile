# your-repo/Dockerfile
FROM golang:1.19
RUN go build -o /bin/app .
ENTRYPOINT ["/bin/app"]
