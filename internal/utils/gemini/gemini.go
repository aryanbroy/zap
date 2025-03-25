package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiResponse(cfg *types.Config, prompt string) {

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

	instructions := "i am building a tech startup, i am fetching reviews from early customers. You are a professional who is going to reply them based on their feedbacks. Be professional and dont write too big replies"

	// prompt := "hello there, who are you?"

	instructionRes, err := model.GenerateContent(ctx, genai.Text(instructions))
	if err != nil {
		log.Fatalf("Unable to generate response: %v", err)
	}

	fmt.Println("Gemini instruction response: ", instructionRes)

	if len(instructionRes.Candidates) > 0 && len(instructionRes.Candidates[0].Content.Parts) > 0 {
		fmt.Println(instructionRes.Candidates[0].Content.Parts[0])
	} else {
		fmt.Println("No response (instruction response) from model")
	}

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
