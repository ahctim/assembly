package assembly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	baseAPI           = "https://api.assemblyai.com"
	defaultTimeout    = 10
	defaultAPIVersion = "v2"
)

// UploadFile is used to upload a local file to the AssemblyAI API
// This will return either an error or a string
// In the case of a successful upload, the string returned is the URL of the uploaded file
func (c *AssemblyAIClient) UploadFile(path string) (string, error) {

	_, err := os.Stat(path)

	if err != nil {
		return "", err
	}

	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	httpClient := createHttpClient(c.Timeout)

	u := fmt.Sprintf("%s/%s/upload", baseAPI, c.APIVersion)

	req, _ := http.NewRequest("POST", u, f)

	req.Header.Add("authorization", c.Token)
	req.Header.Add("Transfer-Encoding", "chunked")

	resp, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	rb, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if err = checkHTTPResponse(resp.StatusCode, rb); err != nil {
		return string(rb), err
	}

	ur := UploadResponse{}

	err = json.Unmarshal(rb, &ur)

	if err != nil {
		return "", err
	}

	return ur.AudioURL, nil
}

// SubmitTranscriptionRequest provides a new audio file URL for AssemblyAI to download and transcribe
// If the request is successful, you will receive the transciption ID and a nil error
// Use RetrieveTranscriptionResult() to retrieve the transcription
func (c *AssemblyAIClient) SubmitTranscriptionRequest(uploadURL string) (string, error) {

	_, err := url.ParseRequestURI(uploadURL)

	if err != nil {
		return "", &ErrInvalidURL{Message: fmt.Sprintf("'%s' is not a valid URL", uploadURL)}
	}

	httpClient := createHttpClient(c.Timeout)

	b := strings.NewReader(fmt.Sprintf("{\"audio_url\": \"%s\"}", uploadURL))

	url := fmt.Sprintf("%s/%s/transcript", baseAPI, c.APIVersion)

	req, _ := http.NewRequest("POST", url, b)

	req.Header.Add("authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	tr := TranscribeResponse{}

	resp, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	rb, err := ioutil.ReadAll(resp.Body)

	if err = checkHTTPResponse(resp.StatusCode, rb); err != nil {
		return string(rb), err
	}

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(rb, &tr)

	if err != nil {

		return "", err
	}

	return tr.ID, nil
}

// RetrieveTranscriptionResult retrieves the status of a transcription request
// When your transcription is ready, you'll receive a pointer to a TranscribeResponse
// along with `true`, and a nil error
func (c *AssemblyAIClient) RetrieveTranscriptionResult(transcriptionID string) (*TranscribeResponse, bool, error) {
	httpClient := createHttpClient(c.Timeout)

	url := fmt.Sprintf("%s/%s/transcript/%s", baseAPI, c.APIVersion, transcriptionID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", c.Token)

	tr := TranscribeResponse{}

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, false, err
	}

	defer resp.Body.Close()

	rb, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, false, err
	}

	if err = checkHTTPResponse(resp.StatusCode, rb); err != nil {

		fmt.Printf("Response body from errored request: %s", string(rb))
		return nil, false, err
	}

	err = json.Unmarshal(rb, &tr)

	if err != nil {
		return nil, false, err
	}

	// TODO check if the status is queued, completed, or something else (not sure what "we hit an error") status string is
	// Then we can return a different error based on the status
	if tr.Status != "completed" {
		return &tr, false, &ErrProcessingNotComplete{Message: fmt.Sprintf("Processing status is '%s'", tr.Status)}
	} else {
		return &tr, true, nil
	}

}
