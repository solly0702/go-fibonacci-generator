# fib-sequence-generator

### Objective

```
- This is simple RESTful Web serivce with golang.
- API(/api/fib-gen) accepts a number, N and returns N length of Fibonacci sequence
- {"fib_max_num":"5"} => {"payload":["0", "1", "1", "2", "3"]}
```

### Deployment

```
-> source gopath.sh
-> docker build -t fib_gen:1.0.1-alpine .
-> docker-compose up -d
```

### Test

```
-> go test ./test -v
```

#### api only provide POST method

```
-> api entry: localhost:5005/api/fib-gen
-> api field: fib_max_num
```

#### Deactive Server and Remove Container

```
-> docker-compose kill
-> docker-compose rm
```

#### Requirement

*   go 1.10
*   docker
