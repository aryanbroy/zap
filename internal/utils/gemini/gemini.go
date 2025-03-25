package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GeminiResponse(cfg *types.Config, prompt string) (ApiResponse, error) {

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

	model := client.GenerativeModel("gemini-1.5-flash")

	instructions := "i am a individual who is fetching reviews from customers. You are a professional who is going to reply them based on their feedbacks. Be professional and dont write too big replies"

	_, err = model.GenerateContent(ctx, genai.Text(instructions))
	if err != nil {
		log.Fatalf("Unable to generate response: %v", err)
	}

	// if len(instructionRes.Candidates) > 0 && len(instructionRes.Candidates[0].Content.Parts) > 0 {
	// 	fmt.Println(instructionRes.Candidates[0].Content.Parts[0])
	// } else {
	// 	fmt.Println("No response (instruction response) from model")
	// }

	jsonPrompt := fmt.Sprintf(`
        Generate ONLY a JSON response with the following structure:
        {
          "status": "success" or "error",
          "message": "A descriptive message"
        }

        For example:
        {
          "status": "success",
          "message": "The operation completed successfully."
        }

        Now, generate a JSON response and provide an appropriate and professional reply for the user's feedback "%v".
        `, prompt)
	resp, err := model.GenerateContent(ctx, genai.Text(jsonPrompt))
	if err != nil {
		log.Fatalf("Unable to generate response: %v", err)
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		fmt.Println(resp.Candidates[0].Content.Parts[0])

		rawText, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
		if !ok {
			log.Fatalf("Error fetching raw text")
		}

		jsonResponse := string(rawText)
		jsonResponse = strings.TrimSpace(jsonResponse)
		jsonResponse = strings.Trim(jsonResponse, "`")
		jsonResponse = strings.Trim(jsonResponse, "json")

		var apiResponse ApiResponse
		err := json.Unmarshal([]byte(jsonResponse), &apiResponse)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}
		return apiResponse, nil
	} else {
		fmt.Println("No response from model")
	}
	return ApiResponse{}, fmt.Errorf("Failed to fetch information from ai")
}
