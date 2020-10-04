// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type woeidResponse []struct {
// 	Title        string `json:"title"`
// 	LocationType string `json:"location_type"`
// 	Woeid        int    `json:"woeid"`
// 	LattLong     string `json:"latt_long"`
// }

// type weatherResponse struct {
// 	ConsolidatedWeather []struct {
// 		ID                   int64     `json:"id"`
// 		WeatherStateName     string    `json:"weather_state_name"`
// 		WeatherStateAbbr     string    `json:"weather_state_abbr"`
// 		WindDirectionCompass string    `json:"wind_direction_compass"`
// 		Created              time.Time `json:"created"`
// 		ApplicableDate       string    `json:"applicable_date"`
// 		MinTemp              float64   `json:"min_temp"`
// 		MaxTemp              float64   `json:"max_temp"`
// 		TheTemp              float64   `json:"the_temp"`
// 		WindSpeed            float64   `json:"wind_speed"`
// 		WindDirection        float64   `json:"wind_direction"`
// 		AirPressure          float64   `json:"air_pressure"`
// 		Humidity             int       `json:"humidity"`
// 		Visibility           float64   `json:"visibility"`
// 		Predictability       int       `json:"predictability"`
// 	} `json:"consolidated_weather"`
// 	Time         time.Time `json:"time"`
// 	SunRise      time.Time `json:"sun_rise"`
// 	SunSet       time.Time `json:"sun_set"`
// 	TimezoneName string    `json:"timezone_name"`
// 	Parent       struct {
// 		Title        string `json:"title"`
// 		LocationType string `json:"location_type"`
// 		Woeid        int    `json:"woeid"`
// 		LattLong     string `json:"latt_long"`
// 	} `json:"parent"`
// 	Sources []struct {
// 		Title     string `json:"title"`
// 		Slug      string `json:"slug"`
// 		URL       string `json:"url"`
// 		CrawlRate int    `json:"crawl_rate"`
// 	} `json:"sources"`
// 	Title        string `json:"title"`
// 	LocationType string `json:"location_type"`
// 	Woeid        int    `json:"woeid"`
// 	LattLong     string `json:"latt_long"`
// 	Timezone     string `json:"timezone"`
// }

// func main() {
// 	fmt.Println("1. Performing Http Get...")
// 	fmt.Println("2. Getting woeid")

// 	var woeid woeidResponse
// 	city := "rio"
// 	url := fmt.Sprintf("https://www.metaweather.com/api/location/search/?query=%s", city)
// 	res, erro := http.Get(url)

// 	if erro != nil {
// 		log.Fatalln(erro)
// 	}

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	json.Unmarshal(body, &woeid)
// 	fmt.Println(woeid[0].Woeid)
// 	id := woeid[0].Woeid
// 	fmt.Println(id)

// 	url = fmt.Sprintf("https://www.metaweather.com/api/location/%d/", id)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to Todo struct
// 	var weather weatherResponse
// 	json.Unmarshal(bodyBytes, &weather)
// 	fmt.Printf("API Response as struct %+v\n", weather)
// }
