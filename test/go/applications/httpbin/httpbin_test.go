package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIGetStatus200(t *testing.T) {
	resp, err := http.Get("http://172.18.255.200:8080/status/200")
	assert.NoError(t, err, "no response")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "status != 200")
}
