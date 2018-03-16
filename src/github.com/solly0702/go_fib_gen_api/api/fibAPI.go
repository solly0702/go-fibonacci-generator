package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Fib type with FibSequence for user input
type Fib struct {
	FibSequence string `json:"fib_sequence"`
}

// FibHandleFunc for fibonacciAPI
func FibHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		fibSequence := r.FormValue("fib_sequence")

		if len(fibSequence) == 0 {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"payload":"generates fib sequence of up to param of fib_sequence"}`))
		} else {
			input, isValid, errors := InputValidator(fibSequence)
			if !isValid {
				writeJSON(w, errors, http.StatusBadRequest)
			} else {
				fibObj := GenFib(input)
				if fibObj["error"][0] == "1" {
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"error":"Out of bounds of maximum positive number possible in fibonacci sequence"}`))
				} else {
					delete(fibObj, "error")
					writeJSON(w, fibObj, http.StatusOK)
				}
			}
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Something wrong with json body data"}`))
		}

		fib := ReadJSON(body)
		input, isValid, errors := InputValidator(fib.FibSequence)

		if !isValid {
			writeJSON(w, errors, http.StatusBadRequest)
		} else {
			fibObj := GenFib(input)
			if fibObj["error"][0] == "1" {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"error":"Out of bounds of maximum positive number possible in fibonacci sequence"}`))
			} else {
				delete(fibObj, "error")
				writeJSON(w, fibObj, http.StatusOK)
			}
		}
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Unsupported request method."}`))
	}
}

// WriteJSON for convert go to json
func writeJSON(w http.ResponseWriter, i interface{}, statusCode int) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
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
func InputValidator(fibSequence string) (int64, bool, map[string]string) {
	errors := make(map[string]string)
	intNum, err := strconv.ParseInt(fibSequence, 10, 64)

	if err != nil {
		errors["error"] = "invalid syntax."
		return intNum, false, errors
	} else if intNum <= 0 {
		errors["error"] = "field must be positive number."
		return intNum, false, errors
	}

	return intNum, true, errors
}

// GenFib for create fib_sequence
func GenFib(userInput int64) map[string][]string {
	const MaxInt64 int = 1<<63 - 1

	fibSequence := make(map[string][]string)

	x, y, temp := int64(0), int64(1), int64(0)

	fibSequence["error"] = append(fibSequence["error"], strconv.FormatInt(0, 10))

	for i := int64(0); i < userInput; i++ {
		if int64(MaxInt64) > x && temp >= 0 {
			fibSequence["payload"] = append(fibSequence["payload"], strconv.FormatInt(int64(x), 10))
			temp = x + y
			x = y
			y = temp
		} else {
			delete(fibSequence, "payload")
			fibSequence["error"][0] = "1"
			break
		}
	}

	return fibSequence
}
