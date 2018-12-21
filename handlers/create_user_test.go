package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/stretchr/testify/assert"
)

// TODO this needs to be written using assert
func TestCreateUserHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	var body = []byte(`{
		"FirstName": "Acme Supplier",
		"LastName": "Goob",
		"Email": "ndavis+29@nav.com",
		"Strategy": {
			"type": "password",
			"Username": "test",
			"Password": "password"
		}
		}`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `a user was created`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
