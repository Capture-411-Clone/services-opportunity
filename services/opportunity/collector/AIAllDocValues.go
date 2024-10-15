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

type EntHunterAllRequestInput struct {
	ObjectKey string `json:"object_key"`
}

type EntHunterAllRequestBody struct {
	Input EntHunterAllRequestInput `json:"input"`
}

/*
 * @apiDefine: ResultItem
 */
type ResultItem struct {
	Value    string `json:"value" openapi:"example:opportunity1;type:string;"`
	Sentence string `json:"sentence" openapi:"example:sentence1;type:string;"`
}

/*
 * @apiDefine: EntHunterAllResponseResult
 */
type EntHunterAllResponseResult struct {
	Title              []ResultItem `json:"TITLE" openapi:"$ref:ResultItem;type:object;"`
	Department         []ResultItem `json:"DEPARTMENT" openapi:"$ref:ResultItem;type:object;"`
	Agency             []ResultItem `json:"AGENCY"  openapi:"$ref:ResultItem;type:object;"`
	Office             []ResultItem `json:"OFFICE"  openapi:"$ref:ResultItem;type:object;"`
	Naics              []ResultItem `json:"NAICS"  openapi:"$ref:ResultItem;type:object;"`
	ContractVehicle    []ResultItem `json:"CONTRACT_VEHICLE"  openapi:"$ref:ResultItem;type:object;"`
	SetAside           []ResultItem `json:"SET_ASIDE"  openapi:"$ref:ResultItem;type:object;"`
	FiscalYear         []ResultItem `json:"FISCAL_YEAR"  openapi:"$ref:ResultItem;type:object;"`
	SolicitationNumber []ResultItem `json:"SOLICITATION_NUMBER"  openapi:"$ref:ResultItem;type:object;"`
	Description        []ResultItem `json:"DESCRIPTION" openapi:"$ref:ResultItem;type:object;"`
}

/*
 * @apiDefine: EntHunterAllResponse
 */
type EntHunterAllResponse struct {
	Status string                     `json:"status" openapi:"example:success;type:string;"`
	IsCui  bool                       `json:"is_cui" openapi:"example:false;type:boolean;"`
	Result EntHunterAllResponseResult `json:"result" openapi:"$ref:EntHunterAllResponseResult;type:object"`
}

func AIGetAllDocValues(fileObjectKey string) (EntHunterAllResponse, error) {
	// Set your OpenAI API key here
	entHunterURI := config.AppConfig.EntHunterURI
	entHunterInternalApiKey := config.AppConfig.EntHunterInternalApiKey

	nerExtractorAllURI, err := url.JoinPath(entHunterURI, "ner-extractor-all")
	if err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error joining URL path: %v", err)
	}

	// Construct the request payload
	requestPayload := EntHunterAllRequestBody{
		Input: EntHunterAllRequestInput{
			ObjectKey: fileObjectKey,
		},
	}

	// Marshal the request payload into JSON
	jsonRequestData, err := json.Marshal(requestPayload)
	if err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error marshaling request payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", nerExtractorAllURI, bytes.NewBuffer(jsonRequestData))
	if err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", entHunterInternalApiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	fmt.Println(string(body))

	var response EntHunterAllResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return EntHunterAllResponse{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return response, nil
}
