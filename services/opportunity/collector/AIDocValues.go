package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Capture-411/services-opportunity/config"
)

type EntHunterRequestInput struct {
	ObjectKey string `json:"object_key"`
}

type EntHunterRequestBody struct {
	Input EntHunterRequestInput `json:"input"`
}

/*
 * @apiDefine: EntHunterResponseResult
 */
type EntHunterResponseResult struct {
	Title              string `json:"title" openapi:"example:opportunity1;type:string;"`
	Department         string `json:"department" openapi:"example:department1;type:string;"`
	Agency             string `json:"agency" openapi:"example:agency1;type:string;"`
	Office             string `json:"office" openapi:"example:office1;type:string;"`
	Naics              string `json:"naics" openapi:"example:naics1;type:string;"`
	ContractVehicle    string `json:"contract_vehicle" openapi:"example:contract_vehicle1;type:string;"`
	SetAside           string `json:"set_aside" openapi:"example:set_aside1;type:string;"`
	Year               string `json:"year" openapi:"example:2021;type:string;"`
	SolicitationNumber string `json:"solicitation_number" openapi:"example:234234ASDF3424;type:string;"`
	Description        string `json:"description" openapi:"example:description1;type:string;"`
}

/*
 * @apiDefine: EntHunterResponse
 */
type EntHunterResponse struct {
	Status string                  `json:"status" openapi:"example:success;type:string;"`
	IsCui  bool                    `json:"is_cui" openapi:"example:false;type:boolean;"`
	Result EntHunterResponseResult `json:"result" openapi:"$ref:EntHunterResponseResult;type:object"`
}

func AIGetDocValues(fileObjectKey string) (EntHunterResponse, error) {
	// Set your OpenAI API key here
	entHunterURI := config.AppConfig.EntHunterURI
	entHunterInternalApiKey := config.AppConfig.EntHunterInternalApiKey

	nerExtractorURI, err := url.JoinPath(entHunterURI, "ner-extractor")
	if err != nil {
		return EntHunterResponse{}, fmt.Errorf("error joining URL path: %v", err)
	}

	// Construct the request payload
	requestPayload := EntHunterRequestBody{
		Input: EntHunterRequestInput{
			ObjectKey: fileObjectKey,
		},
	}

	// Marshal the request payload into JSON
	jsonRequestData, err := json.Marshal(requestPayload)
	if err != nil {
		return EntHunterResponse{}, fmt.Errorf("error marshaling request payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", nerExtractorURI, bytes.NewBuffer(jsonRequestData))
	if err != nil {
		return EntHunterResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", entHunterInternalApiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return EntHunterResponse{}, fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EntHunterResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	fmt.Println(string(body))

	var response EntHunterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return EntHunterResponse{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return response, nil
}
