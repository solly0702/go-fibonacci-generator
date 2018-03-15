export GOPATH=$(pwd) && export PATH=${PATH}:${GOPATH} && cd ${GOPATH}/src/github.com/solly0702/fib_gen && go get -v github.com/stretchr/testify/assert &&GOOS=linux GOARCH=amd64 go build
