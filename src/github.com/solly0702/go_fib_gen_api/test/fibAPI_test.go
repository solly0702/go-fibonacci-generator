// go test ./test -v
package test

import (
	"testing"

	"github.com/solly0702/go_fib_gen_api/api"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	json := []byte(`{"fib_sequence":"5"}`)
	fib := api.ReadJSON(json)

	assert.Equal(t, api.Fib{FibSequence: "5"}, fib, "Fib JSON Unmarshalling wrong")
}

func TestInputValidatorSuccess(t *testing.T) {
	fib := "5"
	input, isValid, errors := api.InputValidator(fib)
	errs := make(map[string]string)

	assert.Equal(t, int64(5), input, "Input Validator expect to return int64 of 5 in input, but failed")
	assert.True(t, isValid, "Input validator expect to return true in isValid, but failed")
	assert.Equal(t, errs, errors, "Input validator expect to return empty map[] in errors, but failed")
}

func TestInputValidatorFailWithEmptyVal(t *testing.T) {
	fib := ""
	input, isValid, errors := api.InputValidator(fib)
	errs := make(map[string]string)
	errs["error"] = "invalid syntax."

	assert.Equal(t, int64(0), input, "Input Validator expect to return int64 of 0 in input, but failed")
	assert.True(t, !isValid, "Input validator expect to return false in isValid, but failed")
	assert.Equal(t, errs, errors, "Input validator expect to return map[] with error in it in errors, but failed")
}

func TestInputValidatorFailWithNegVal(t *testing.T) {
	fib := "-5"
	input, isValid, errors := api.InputValidator(fib)
	errs := make(map[string]string)
	errs["error"] = "field must be positive number."

	assert.Equal(t, int64(-5), input, "Input Validator expect to return int64 of -5 in input, but failed")
	assert.True(t, !isValid, "Input validator expect to return false in isValid, but failed")
	assert.Equal(t, errs, errors, "Input validator expect to return empty map[] with error in it in errors, but failed")
}

func TestInputValidatorFailWithNonNumeric(t *testing.T) {
	fib := "abcde"
	input, isValid, errors := api.InputValidator(fib)
	errs := make(map[string]string)
	errs["error"] = "invalid syntax."

	assert.Equal(t, int64(0), input, "Input Validator expect to return int64 of 0 in input, but failed")
	assert.True(t, !isValid, "Input validator expect to return false in isValid, but failed")
	assert.Equal(t, errs, errors, "Input validator expect to return empty map[] with error in it in errors, but failed")
}

func TestGenFibSuccess(t *testing.T) {
	fibSequence := api.GenFib(int64(5))
	result := make(map[string][]string)
	result["payload"] = append(result["payload"], "0", "1", "1", "2", "3")
	result["error"] = append(result["error"], "0")
	assert.Equal(t, result, fibSequence, "GenFib expect to return map[payload: data], but failed")
}

func TestGenFibOutofInt(t *testing.T) {
	fibSequence := api.GenFib(int64(100))
	result := make(map[string][]string)
	result["error"] = append(result["error"], "1")

	assert.Equal(t, result, fibSequence, `GenFib expect to return map[error: ["0"]], but failed`)
}
