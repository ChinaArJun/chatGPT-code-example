package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func main() {
	// Set up the HTTP client
	client := &http.Client{}

	// Set up the request parameters
	values := url.Values{}
	values.Set("app_key", "YOUR_APP_KEY")
	values.Set("method", "taobao.tbk.tpwd.create")
	values.Set("format", "json")
	values.Set("v", "2.0")
	values.Set("sign_method", "md5")
	values.Set("timestamp", "YOUR_TIMESTAMP")
	values.Set("text", "YOUR_TEXT")
	values.Set("url", "YOUR_URL")
	values.Set("user_id", "YOUR_USER_ID")
	values.Set("adzone_id", "YOUR_ADZONE_ID")
	values.Set("site_id", "YOUR_SITE_ID")

	// Generate the sign
	sign := generateSign(values, "YOUR_APP_SECRET")

	// Add the sign to the request parameters
	values.Set("sign", sign)

	// Set up the request URL
	url := "https://eco.taobao.com/router/rest?" + values.Encode()

	// Send the request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}

func generateSign(values url.Values, appSecret string) string {
	// Sort the keys
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Concatenate the key-value pairs
	var buffer strings.Builder
	for _, k := range keys {
		buffer.WriteString(k)
		buffer.WriteString(values.Get(k))
	}

	// Add the app secret
	buffer.WriteString(appSecret)

	// Calculate the MD5 hash
	hash := md5.Sum([]byte(buffer.String()))

	// Convert the hash to a hex string
	return hex.EncodeToString(hash[:])
}
