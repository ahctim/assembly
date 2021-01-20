package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchya1/assembly"
)

func main() {
	if os.Getenv("ASSEMBLY_AI_TOKEN") == "" {
		panic("ASSEMBLY_AI_TOKEN is an empty env var")
	}

	aai := assembly.NewDefaultClient(os.Getenv("ASSEMBLY_AI_TOKEN"))

	u := uploadAudioFile(aai)
	id := startTranscription(aai, u)

	for {
		ready, result, err := getTranscriptionStatus(aai, id)
		if err != nil {
			panic(err)
		}
		if !ready {
			fmt.Println("Not ready yet")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("Result!")
			fmt.Println(result)
			break
		}

	}
}

func uploadAudioFile(aai *assembly.AssemblyAIClient) string {
	p, _ := os.Getwd()

	downloadURL, err := aai.UploadFile(fmt.Sprintf("%s/audio/sample_file.wav", p))

	if err != nil {
		panic(err)
	}

	fmt.Println(downloadURL)
	return downloadURL
}

func startTranscription(aai *assembly.AssemblyAIClient, url string) string {
	id, err := aai.SubmitTranscriptionRequest(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	return id
}

func getTranscriptionStatus(aai *assembly.AssemblyAIClient, id string) (bool, string, error) {

	result, ready, err := aai.RetrieveTranscriptionResult(id)

	if !ready {
		if _, ok := err.(*assembly.ErrProcessingNotComplete); ok {
			fmt.Println("Not ready yet!")
			return false, "", nil
		}
		fmt.Println("Error processing audio file")
		return false, "", err
	}

	return true, result.Text, nil
}
