# go-fib-sequence-generator

### Objective

```
- This is simple RESTful Web serivce with golang.
- API(/api/fib-gen) accepts a number, N and returns N length of Fibonacci sequence
- {"fib_sequence":"5"} => {"payload":["0", "1", "1", "2", "3"]}
```

### Deployment

```
-> source gopath.sh
-> docker build -t go_fib_gen_api:1.0.0-alpine .
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

*   golang
*   docker
