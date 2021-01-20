package assembly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultClient(t *testing.T) {

	token := "Hello world"

	c := NewDefaultClient(token)

	if !assert.Equal(t, defaultAPIVersion, c.APIVersion) {
		t.Fail()
	}

	if !assert.Equal(t, defaultTimeout, c.Timeout) {
		t.Fail()
	}

	if !assert.Equal(t, token, c.Token) {
		t.Fail()
	}
}

func TestCreateHttpClient(t *testing.T) {
	timeout := 5
	createHttpClient(timeout)
}

func TestCheckHTTPResponse(t *testing.T) {
	var err error
	err = checkHTTPResponse(403, []byte("403!"))

	if _, ok := err.(*ErrUnauthorized); !ok {
		t.Fail()
	}

	err = checkHTTPResponse(400, []byte("400!"))

	if _, ok := err.(*ErrBadRequest); !ok {
		t.Fail()
	}

	err = checkHTTPResponse(200, []byte("200!"))

	if err != nil {
		t.Fail()
	}

}
