package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	cloudflareAPIBase = "https://api.cloudflare.com/client/v4/"
)

type DNSRecord struct {
	ID string `json:"id"`
}

type DNSRecordsResponse struct {
	Result []DNSRecord `json:"result"`
}

type Zone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ZonesResponse struct {
	Result []Zone `json:"result"`
}

func getZoneID(apiToken, domainName string) (string, error) {
	fmt.Println("getZoneID() called...")
	client := &http.Client{}

	req, err := http.NewRequest("GET", cloudflareAPIBase+"zones?name="+domainName, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var zonesResponse ZonesResponse
	err = json.Unmarshal(body, &zonesResponse)
	if err != nil {
		return "", err
	}

	if len(zonesResponse.Result) == 0 {
		return "", fmt.Errorf("no zone found with domain name %s", domainName)
	}

	return zonesResponse.Result[0].ID, nil
}

func deleteAllDNSRecords(apiToken, domainName string) error {
	fmt.Println("deleteAllDNSRecords() called...")
	zoneID, err := getZoneID(apiToken, domainName)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", cloudflareAPIBase+"zones/"+zoneID+"/dns_records", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dnsRecords DNSRecordsResponse
	err = json.Unmarshal(body, &dnsRecords)
	if err != nil {
		return err
	}

	for _, record := range dnsRecords.Result {
		req, err := http.NewRequest("DELETE", cloudflareAPIBase+"zones/"+zoneID+"/dns_records/"+record.ID, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+apiToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Deleted DNS record:", record.ID)
		} else {
			return fmt.Errorf("failed to delete DNS record %s, status code: %d", record.ID, resp.StatusCode)
		}
	}

	return nil
}
func main() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		fmt.Println("API_TOKEN environment variable is not set")
		return
	}

	domainName := os.Getenv("DOMAIN")
	if domainName == "" {
		fmt.Println("DOMAIN environment variable is not set")
		return
	}

	err := deleteAllDNSRecords(apiToken, domainName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
