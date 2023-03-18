package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
}

func main() {
	url := "http://t.weather.sojson.com/api/weather/city/101280601"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	weatherData := data["data"].(map[string]interface{})
	weather := Weather{
		Temperature: weatherData["wendu"].(string),
		Weather:     weatherData["forecast"].([]interface{})[0].(map[string]interface{})["type"].(string),
	}

	fmt.Printf("深圳天气：温度%s℃，天气%s\n", weather.Temperature, weather.Weather)
}
