package generate_token

import "time"

type GenerateTokenWorkflowOutput struct {
	Token string `json:"token"`
	ExpiresIn time.Time `json:"expires_in"`
}
