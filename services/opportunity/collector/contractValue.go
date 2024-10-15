package collector

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func ContractValue(solNumber string) (string, bool, error) {
	baseURL := "https://www.fpds.gov"

	// remove any special chars and spaces from the solicitation number
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	solNumber = re.ReplaceAllString(solNumber, "")

	log.Printf("Starting to scrape contract value for solicitation: %s\n", solNumber)

	// Create HTTP client
	client := &http.Client{}
	log.Println("HTTP client created.")

	// Generate the URL for the search
	searchURL := fmt.Sprintf("%s/ezsearch/search.do?indexName=awardfull&templateName=1.5.3&s=FPDS.GOV&q=%s", baseURL, solNumber)
	log.Printf("Search URL generated: %s\n", searchURL)

	// Fetch the search page
	resp, err := fetchPage(client, searchURL, baseURL)
	if err != nil {
		return "", false, fmt.Errorf("error fetching search page: %v", err)
	}
	defer resp.Body.Close()
	log.Println("Search page fetched successfully.")

	// Parse the search page with goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", false, fmt.Errorf("error loading http response body: %v", err)
	}
	log.Println("Search page loaded into goquery document.")

	// Check if this is multi award
	// /html/body/table[4]/tbody/tr/td[1]/table[3]/tbody/tr[2]/td/font
	awardCount := doc.Find("html > body > table:nth-of-type(4) > tbody >" +
		" tr > td:nth-of-type(1) > table:nth-of-type(3) >" +
		" tbody > tr:nth-of-type(2) > td > font").Find("p").Length()

	fmt.Println("awardCount:", awardCount)

	if awardCount > 1 {
		log.Println("This is a multi-award solicitation.")
		return "", true, nil
	}

	// Find the specific link on the first page
	var contractValue string
	var valueFindErr error
	doc.Find("html > body > table:nth-of-type(4) >" +
		" tbody > tr > td:nth-of-type(2) > table >" +
		" tbody > tr:nth-of-type(2) > td > table >" +
		" tbody > tr > td > table:nth-of-type(1) >" +
		" tbody > tr:nth-of-type(1) > td:nth-of-type(2) >" +
		" span > a:nth-of-type(3)").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if !exists {
			valueFindErr = fmt.Errorf("no href attribute found")
			return
		}

		data := href
		re := regexp.MustCompile(`javascript:getParentURL\('([^']+)'\)`)
		matches := re.FindStringSubmatch(data)
		if len(matches) > 1 {
			fmt.Println("Extracted URL:", matches[1])
		} else {
			log.Fatal("No URL found")
		}

		link := joinURL(baseURL, matches[1])
		log.Printf("Found link: %s\n", link)

		// Fetch the details page
		detailResp, err := fetchPage(client, link, baseURL)
		if err != nil {
			valueFindErr = fmt.Errorf("error fetching detail page: %v", err)
			return
		}
		defer detailResp.Body.Close()
		log.Println("Detail page fetched successfully.")

		// Parse the details page with goquery
		detailDoc, err := goquery.NewDocumentFromReader(detailResp.Body)
		if err != nil {
			valueFindErr = fmt.Errorf("error loading detail response body: %v", err)
			return
		}
		log.Println("Detail page loaded into goquery document.")

		// Extract contract value
		contractValue = detailDoc.Find("#totalUltimateContractValue").AttrOr("value", "")
		log.Printf("Contract Value: %s\n", contractValue)

		if contractValue == "" {
			contractValue = detailDoc.Find("#ultimateContractValue").AttrOr("value", "")
		}
	})

	if valueFindErr != nil {
		return "", false, valueFindErr
	}

	// remove , from string number
	re = regexp.MustCompile(`,`)
	contractValue = re.ReplaceAllString(contractValue, "")

	// remove the cent from the whole dollor value
	re = regexp.MustCompile(`\.\d{2}`)
	contractValue = re.ReplaceAllString(contractValue, "")

	// remove dollor sign
	re = regexp.MustCompile(`\$`)
	contractValue = re.ReplaceAllString(contractValue, "")

	return contractValue, false, nil
}

// fetchPage performs an HTTP GET request for a given URL
func fetchPage(client *http.Client, pageURL string, baseUrl string) (*http.Response, error) {
	log.Printf("Fetching page: %s\n", pageURL)

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		log.Printf("Error creating new request: %v\n", err)
		return nil, err
	}

	// Mimic a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Cache-Control", "max-age=0")
	log.Println("HTTP request created successfully.")

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v\n", err)
		return nil, err
	}
	log.Println("HTTP request completed successfully.")

	return response, nil
}

// joinURL safely joins a base URL and a path without encoding
func joinURL(base, path string) string {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatal("Parsing Base Path", err)
	}

	// Concatenate the base URL and the path directly
	fullURL := u.String() + path

	return fullURL
}
