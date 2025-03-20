package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiResponse(cfg *types.Config) {

	ctx := context.Background()
	apiKey := cfg.GEMINI_API
	if apiKey == "" {
		log.Fatalln("Missing gemini api key")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create client: %v", err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	prompt := "hello there, who are you?"

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Unable to generate response: %v", err)
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	} else {
		fmt.Println("No response from model")
	}
}
