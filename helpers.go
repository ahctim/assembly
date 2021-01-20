package assembly

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func NewDefaultClient(token string) *AssemblyAIClient {
	c := &AssemblyAIClient{
		Token:      token,
		APIVersion: defaultAPIVersion,
		Timeout:    defaultTimeout,
	}

	return c
}

func createHttpClient(timeout int) *http.Client {

	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func verifyAudioFileType(path string) error {

	// https://mimesniff.spec.whatwg.org/#matching-an-audio-or-video-type-pattern
	// TODO figure this out
	allowedTypes := map[string]string{
		"audio/wave": "ok",
		"audio/mpeg": "ok",
		"video/mp4":  "ok",
	}

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	buffer := make([]byte, 512)

	_, err = f.Read(buffer)

	if err != nil {
		return err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	if _, ok := allowedTypes[contentType]; ok {
		return nil
	} else {
		return &ErrUnsupportedFileType{Message: fmt.Sprintf("'%s' is not a supported file type", contentType)}
	}

}

func checkHTTPResponse(code int, respBody []byte) error {
	switch code {
	case 403:
		return &ErrUnauthorized{HTTPCode: 403, Message: string(respBody)}
	case 400:
		return &ErrBadRequest{HTTPCode: 400, Message: string(respBody)}
	default:
		return nil
	}

}
