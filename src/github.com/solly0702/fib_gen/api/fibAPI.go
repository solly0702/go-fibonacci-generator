package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// MaxInt64 for compare to FibMaxNum
const MaxInt64 int = 1<<63 - 1

// Fib type with FibMaxNum for user input
type Fib struct {
	FibMaxNum string `json:"fib_max_num"`
}

// FibHandleFunc for fibonacciAPI
func FibHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Something wrong with json body data"}`))
		}

		fib := ReadJSON(body)
		input, isValid, errors := InputValidator(fib)
		fmt.Println("this is input :: ", input, isValid, errors)
		if !isValid {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, errors)
		} else {
			fibSequence := GenFib(input)
			w.WriteHeader(http.StatusOK)
			writeJSON(w, fibSequence)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Unsupported request method."}`))
	}
}

// WriteJSON for convert go to json
func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// ReadJSON for convert json to go
func ReadJSON(data []byte) Fib {
	fib := Fib{}
	err := json.Unmarshal(data, &fib)
	if err != nil {
		panic(err)
	}
	return fib
}

// InputValidator for validating data
func InputValidator(fib Fib) (int64, bool, map[string]string) {
	errors := make(map[string]string)
	intNum, err := strconv.ParseInt(fib.FibMaxNum, 10, 64)

	if err != nil {
		errors["error"] = "invalid syntax."
		return intNum, false, errors
	} else if intNum <= 0 {
		errors["error"] = "invalid value, value must be positive number."
		return intNum, false, errors
	} else if intNum > 92 {
		errors["error"] = "max fibonacci must be lower than 92."
		return intNum, false, errors
	}
	return intNum, true, errors
}

// GenFib for create fib_sequence
func GenFib(userInput int64) map[string][]string {
	fibSequence := make(map[string][]string)

	x, y, temp := int64(0), int64(1), int64(0)

	for i := int64(0); i < userInput; i++ {
		if int64(MaxInt64) > x && temp >= 0 {
			fibSequence["payload"] = append(fibSequence["payload"], strconv.FormatInt(int64(x), 10))
			temp = x + y
			x = y
			y = temp
		}
	}

	return fibSequence
}
