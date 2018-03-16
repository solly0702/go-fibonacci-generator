# GOOS=linux GOARCH=amd64 go build
# docker build -t go_fib_gen_api:1.0.0-alpine .
# FROM alpine:3.7

# COPY ./fib_gen /app/fib_gen
# RUN chmod +x /app/fib_gen

# ENV PORT 8080
# EXPOSE 8080

# ENTRYPOINT /app/fib_gen


### HEAVY IMANGE? 
# docker build -t golang_fib_gen:1.0.0-alpine .
FROM golang:1.10-alpine
LABEL Sol Lee

ENV SOURCES /go/src/github.com/solly0702/go_fib_gen_api/
ENV PORT 5005

COPY . ${SOURCES}
WORKDIR ${SOURCES}

RUN go get -v github.com/stretchr/testify/assert

RUN cd ${SOURCES} && CGO_ENABLED=0 go install
EXPOSE 5005

ENTRYPOINT fib_gen