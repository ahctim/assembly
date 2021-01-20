# assembly

## Usage

Check `test/main.go` for a working example.

## TODO

Create custom errors
  - Handle cases [here](https://docs.assemblyai.com/overview/errors-and-failed-transcripts)

Write tests
  - Need to mock Assembly AI API

Create complete API response structs

Add support for specifying [accoustic model](https://docs.assemblyai.com/guides/transcribing-with-a-different-acoustic-or-custom-language-model)

Add support for specifying webhooks in `SubmitTranscriptionRequest`

Before uploading, ensure a file is of a [supported media type](https://docs.assemblyai.com/overview/supported-file-formats)

Add support for [deleting a transcription](https://docs.assemblyai.com/all-guides/deleting-a-transcription-from-the-api)

Add support for [stream endpoint](https://docs.assemblyai.com/api-ref/v2-stream)