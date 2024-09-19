package main

import (
	"net/http"
	"testing"

	"snippetbox.ayonchakroborty.net/internal/assert"
)

// end to end test
func TestPing(t *testing.T){
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")
	
	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}