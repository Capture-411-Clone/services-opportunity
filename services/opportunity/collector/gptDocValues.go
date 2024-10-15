package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/exp"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

/*
 * @apiDefine: GptChoiceValues
 */
type GptChoiceValues struct {
	Title              string `json:"title" openapi:"example:opportunity1;type:string;"`
	Department         string `json:"department" openapi:"example:department1;type:string;"`
	Agency             string `json:"agency" openapi:"example:agency1;type:string;"`
	Office             string `json:"office" openapi:"example:office1;type:string;"`
	Naics              string `json:"naics" openapi:"example:naics1;type:string;"`
	ContractVehicle    string `json:"contract_vehicle" openapi:"example:contract_vehicle1;type:string;"`
	SetAside           string `json:"set_aside" openapi:"example:set_aside1;type:string;"`
	Year               string `json:"year" openapi:"example:2021;type:string;"`
	SolicitationNumber string `json:"solicitation_number" openapi:"example:234234ASDF3424;type:string;"`
}

func GptGetDocValues(text string, prompt string) (GptChoiceValues, error) {
	// Set your OpenAI API key here
	apiKey := config.AppConfig.ChatGptApiKey

	systemPromt := "What is the title of this government solicitation document? " +
		"Also tell me the government department, agency, office, NAICS code, " +
		"contract vehicle, set aside, year this was issued, and solicitation " +
		"number and return each item if exist in json format. the format is " +
		"{'title': 'value', 'department': 'value', 'agency': 'value', " +
		"'office': 'value', 'naics': 'value', 'contract_vehicle': 'value', " +
		"'set_aside': 'value', 'year': 'value', 'solicitation_number': 'value'}"

	content := exp.TerIf(prompt != "", prompt, systemPromt)

	// Construct the request payload
	requestPayload := ChatCompletionRequest{
		// Model: "gpt-3.5-turbo-16k-0613",
		Model: "gpt-4",
		Messages: []Message{
			{
				Role:    "system",
				Content: content,
			},
			{
				Role:    "user",
				Content: text,
			},
		},
		Temperature: 0.7,
		TopP:        1.0,
	}

	// Marshal the request payload into JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		return GptChoiceValues{}, fmt.Errorf("error marshaling request payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return GptChoiceValues{}, fmt.Errorf("error creating request: %v", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GptChoiceValues{}, fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GptChoiceValues{}, fmt.Errorf("error reading response: %v", err)
	}

	fmt.Println(string(body))

	var response ChatCompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return GptChoiceValues{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(response.Choices) == 0 {
		return GptChoiceValues{}, fmt.Errorf("no choices in response. means no data found")
	}

	// Iterate through the choices (if any) and print the text content
	for k, choice := range response.Choices {
		fmt.Printf("Response %d: %s\n", k, choice.Message.Content)
	}

	fixedBody := strings.ReplaceAll(response.Choices[0].Message.Content, `\"`, `"`)
	fixedBody = strings.ReplaceAll(fixedBody, "None", "")
	fixedBody = strings.ReplaceAll(fixedBody, "N/A", "")
	fixedBody = strings.ReplaceAll(fixedBody, "True", "true")
	fixedBody = strings.ReplaceAll(fixedBody, "False", "false")
	fixedBody = strings.ReplaceAll(fixedBody, "Unknown", "")
	fixedBody = strings.ReplaceAll(fixedBody, `\n`, "")
	fixedBody = strings.ReplaceAll(fixedBody, `'`, `"`)

	fmt.Println("fixedBody:", fixedBody)

	// Extract the values from the response
	gptChoiceValues := GptChoiceValues{}
	if err := json.Unmarshal([]byte(fixedBody), &gptChoiceValues); err != nil {
		return GptChoiceValues{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return gptChoiceValues, nil
}
