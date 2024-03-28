package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func containsSubstring(str, substr string) bool {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	// Create request for services
	endpoint := "https://console.runzero.com/api/v1.0/export/org/services.json?search=alive%3At%20protocol%3Atls"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	req.Header.Set("Authorization", "Bearer UPDATE_ME")

	// Send the request and get the response
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse response and loop the services
	var serviceJson []map[string]interface{}
	json.Unmarshal(body, &serviceJson)
	for _, service := range serviceJson {
		fmt.Println("asset_id", "-->", service["id"])
		fmt.Println("service_id", "-->", service["service_id"])
		loopValues := reflect.ValueOf(service["service_data"]).MapKeys()
		for _, service_attribute := range loopValues {
			if containsSubstring(service_attribute.String(), "tls") {
				fmt.Println(service_attribute.String(), "-->", service["service_data"].(map[string]interface{})[service_attribute.String()])
			}
		}

		// This can be used to get a full list of keys for reference
		// for k, _ := range v {
		// 	fmt.Println(k)
		// }

		fmt.Println("---")
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

}
