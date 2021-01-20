package assembly

// TODO add all fields
type TranscribeResponse struct {
	ID            string  `json:"id"`
	Status        string  `json:"status"`
	Text          string  `json:"text"`
	AccouticModel string  `json:"acoustic_model"`
	AudioURL      string  `json:"audio_url"`
	AudioDuration float64 `json:"audio_duration"`
	Confidence    float64 `json:"confidence"`
	//dual_channel
	FormatText    bool   `json:"format_text"`
	LanguageModel string `json:"language_model"`
	Punctuate     bool   `json:"punctuate"`
	// utterances
	WebhookStatusCode int                  `json:"webhook_status_code"`
	WebhookURL        string               `json:"webhook_url"`
	Words             []TranscribeResponse `json:"words"`
}

type TranscribeWords struct {
	Confidence float64 `json:"confidence"`
	End        int     `json:"end"`
	Start      int     `json:"start"`
	Text       string  `json:"text"`
}

type UploadResponse struct {
	AudioURL string `json:"upload_url"`
}

type AssemblyAIClient struct {
	Token      string
	APIVersion string
	Timeout    int
}
