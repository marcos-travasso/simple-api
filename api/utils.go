package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http/httptest"
	"testing"
)

func assertBody(t *testing.T, w *httptest.ResponseRecorder, status int, body string) {
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	assert.Equal(t, status, res.StatusCode)
	assert.Equal(t, body, string(data))
}
