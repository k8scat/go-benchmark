package main

import (
	"testing"
)

var httpStatus = map[int]string{
	200: "ok",
	404: "not found",
}

var result = make([]string, 3)

func doMap(code int) string {
	status, ok := httpStatus[code]
	if !ok {
		status = ""
	}

	return status
}

func doSwitch(code int) string {
	var message = ""

	switch code {
	case 200:
		message = "ok"
	case 404:
		message = "not found"
	}

	return message
}

func BenchmarkSwitch(b *testing.B) {
	r := make([]string, 3)

	for n := 0; n < b.N; n++ {
		r[0] = doSwitch(200)
		r[1] = doSwitch(404)
		r[2] = doSwitch(999)
	}

	result = r
}

func BenchmarkMap(b *testing.B) {
	r := make([]string, 3)

	for n := 0; n < b.N; n++ {
		r[0] = doMap(200)
		r[1] = doMap(404)
		r[2] = doMap(999)
	}

	result = r
}
