package collector

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"sync"
)

var (
	commaRe       = regexp.MustCompile(`,`)
	centsRe       = regexp.MustCompile(`\.\d{2}`)
	dollarSignRe  = regexp.MustCompile(`\$`)
	criteriaValRe = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

// Define the data structures
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
	Links   []Link   `xml:"link"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Entry struct {
	Title   string  `xml:"title"`
	Content Content `xml:"content"`
}

type Content struct {
	IDV   *IDV   `xml:"IDV"`
	Award *Award `xml:"award"`
}

type Award struct {
	DollarValues      DollarValues      `xml:"dollarValues"`
	TotalDollarValues TotalDollarValues `xml:"totalDollarValues"`
	Vendor            Vendor            `xml:"vendor"`
}

type IDV struct {
	DollarValues      DollarValues      `xml:"dollarValues"`
	TotalDollarValues TotalDollarValues `xml:"totalDollarValues"`
	Vendor            Vendor            `xml:"vendor"`
}

type Vendor struct {
	VendorHeader VendorHeader `xml:"vendorHeader"`
}

type VendorHeader struct {
	VendorName string `xml:"vendorName"`
}

type DollarValues struct {
	ObligatedAmount        string `xml:"obligatedAmount"`
	BaseAndAllOptionsValue string `xml:"baseAndAllOptionsValue"`
}

type TotalDollarValues struct {
	TotalObligatedAmount        string `xml:"totalObligatedAmount"`
	TotalBaseAndAllOptionsValue string `xml:"totalBaseAndAllOptionsValue"`
}

// PrepareFeedURL cleans the input value and constructs the URL
func PrepareFeedURL(searchCriteria, value string) string {
	v := criteriaValRe.ReplaceAllString(value, "")
	baseURL := "https://www.fpds.gov/ezsearch/FEEDS/ATOM?FEEDNAME=PUBLIC&q=%s:\"%s\"&feed=atom0.3"
	return fmt.Sprintf(baseURL, searchCriteria, v)
}

func fetchDataFromURL(url string) ([]byte, error) {
	// Make the HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the HTTP status is in the 200-299 range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return data, nil
}

// fetches and decodes the XML from the given URL
func unmarshalXml(data []byte) (Feed, error) {

	var feed Feed
	if err := xml.Unmarshal(data, &feed); err != nil {
		log.Printf("Error unmarshalling XML: %v", err)
		return Feed{}, fmt.Errorf("error unmarshalling XML: %v", err)
	}

	return feed, nil
}

// FindLastPageLink looks for a link with rel="next" and returns the href
func findLastPageLink(links []Link) string {
	for _, link := range links {
		if link.Rel == "last" {
			return link.Href
		}
	}
	return "" // Return empty if no next link is found
}

// getPageURLs takes the URL from the <link rel="last"> and generates all page URLs.
func getPageURLs(lastURL string) ([]string, error) {
	// Parse the last URL to get the query parameters
	var links []string
	parsedURL, err := url.Parse(lastURL)
	if err != nil {
		return links, err
	}

	// Extract query parameters
	queryParams := parsedURL.Query()

	// Extract and print the original 'start' parameter value
	startParam := queryParams.Get("start")
	if startParam == "" {
		return links, err
	}
	startValue, err := strconv.Atoi(startParam)
	if err != nil {
		return links, err
	}
	for i := 1; i <= (startValue / 10); i++ {

		newStartValue := i * 10
		queryParams.Set("start", strconv.Itoa(newStartValue))

		// Reconstruct the URL with the updated query parameters
		parsedURL.RawQuery = queryParams.Encode()
		newURL := parsedURL.String()
		links = append(links, newURL)
	}
	return links, nil
}

// processes multiple pages of feed data, logging entry details
func processAllPages(feedURL string) ([]Feed, error) {
	var lastPageLink string
	pagesFeed := []Feed{}

	log.Println("Fetching URL:", feedURL)

	data, err := fetchDataFromURL(feedURL)
	if err != nil {
		log.Printf("Error fetching data from URL %s: %v", feedURL, err)
		return []Feed{}, err
	}

	log.Println("Data fetched successfully, processing...")
	firstPageFeed, err := unmarshalXml(data)
	if err != nil {
		return []Feed{}, err
	}

	pagesFeed = append(pagesFeed, firstPageFeed)

	lastPageLink = findLastPageLink(firstPageFeed.Links) // Initialize nextPageLink
	urls, err := getPageURLs(lastPageLink)
	if err != nil {
		return []Feed{}, err
	}

	// For pagination
	var wg sync.WaitGroup
	resultChan := make(chan Feed)
	errorChan := make(chan error, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			data, err := fetchDataFromURL(url)
			if err != nil {
				errorChan <- fmt.Errorf("error fetching data from URL %s: %v", url, err)
				return
			}
			pageFeed, err := unmarshalXml(data)
			if err != nil {
				errorChan <- fmt.Errorf("error processing data from URL %s: %v", url, err)
				return
			}
			resultChan <- pageFeed
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	for feed := range resultChan {
		pagesFeed = append(pagesFeed, feed)
	}

	if len(errorChan) > 0 {
		for err := range errorChan {
			log.Println(err)
		}
		return []Feed{}, fmt.Errorf("errors occurred during page processing")
	}

	log.Println("Atom feed processed successfully.")

	return pagesFeed, nil
}

func isMultiAward(feeds []Feed) bool {
	seenAwardVendorNames := make(map[string]bool)
	uniqueVendorCount := 0

	for _, feed := range feeds {
		for _, entry := range feed.Entries {
			var vendorName string
			if entry.Content.Award != nil {
				vendorName = entry.Content.Award.Vendor.VendorHeader.VendorName
			} else if entry.Content.IDV != nil {
				vendorName = entry.Content.IDV.Vendor.VendorHeader.VendorName
			}

			if vendorName != "" {
				if !seenAwardVendorNames[vendorName] {
					seenAwardVendorNames[vendorName] = true
					uniqueVendorCount++
				}
			}
		}
	}
	result := uniqueVendorCount > 1
	if result {
		log.Println("This is a multi-award solicitation.")
	}
	return result // Returns true if more than one unique vendor name is found
}

// CleanContractValue takes a contract value string as input
// and returns a cleaned version of the string without commas,
// cents, and dollar signs.
func CleanContractValue(contractValue string) string {
	contractValue = commaRe.ReplaceAllString(contractValue, "")
	contractValue = centsRe.ReplaceAllString(contractValue, "")
	contractValue = dollarSignRe.ReplaceAllString(contractValue, "")
	return contractValue
}

func ExtractContractValueFromAtomFeeds(solNumber string) (string, bool, error) {
	// Initialize contract value as empty
	var contVal string
	criteria := "SOLICITATION_ID"
	url := PrepareFeedURL(criteria, solNumber)
	feeds, err := processAllPages(url)
	if err != nil {
		return "", false, err // Return on error
	}
	multiAward := isMultiAward(feeds)
	if multiAward {
		return "", multiAward, nil
	}

	firstEntry := feeds[0].Entries[0].Content
	if firstEntry.IDV != nil {
		if firstEntry.IDV.DollarValues.BaseAndAllOptionsValue != "" {
			contVal = firstEntry.IDV.DollarValues.BaseAndAllOptionsValue
		} else {
			contVal = firstEntry.IDV.TotalDollarValues.TotalBaseAndAllOptionsValue
		}
	} else if firstEntry.Award != nil {
		if firstEntry.Award.DollarValues.BaseAndAllOptionsValue != "" {
			contVal = firstEntry.Award.DollarValues.BaseAndAllOptionsValue
		} else {
			contVal = firstEntry.Award.TotalDollarValues.TotalBaseAndAllOptionsValue
		}
	}

	// Check if there are multiple awards
	return CleanContractValue(contVal), false, nil // Return the contract value and the multi-award status
}
