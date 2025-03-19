package google

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aryanbroy/zap/internal/types"
)

func FetchSheets(sheetId string, accessToken string) (types.SheetResponse, error) {
	url := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%v/values/responses", sheetId)

	log.Println("Making a new request...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.SheetResponse{}, err
	}

	log.Println("Setting authorization token...")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return types.SheetResponse{}, err
	}

	defer res.Body.Close()

	log.Println("Reading response body...")
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return types.SheetResponse{}, err
	}

	var sheetData types.SheetResponse

	log.Println("Unmarshaling data...")
	if err := json.Unmarshal(data, &sheetData); err != nil {
		log.Fatalln("Error unmarshaling json")
		return types.SheetResponse{}, err
	}

	fmt.Println("Sheetdata: ", sheetData)
	return types.SheetResponse{}, nil
}
