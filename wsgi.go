package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"os"

	"github.com/gorilla/mux"
)

var templates *template.Template

type woeidResponse []struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}

type weatherResponse struct {
	ConsolidatedWeather []struct {
		ID                   int64     `json:"id"`
		WeatherStateName     string    `json:"weather_state_name"`
		WeatherStateAbbr     string    `json:"weather_state_abbr"`
		WindDirectionCompass string    `json:"wind_direction_compass"`
		Created              time.Time `json:"created"`
		ApplicableDate       string    `json:"applicable_date"`
		MinTemp              float64   `json:"min_temp"`
		MaxTemp              float64   `json:"max_temp"`
		TheTemp              float64   `json:"the_temp"`
		WindSpeed            float64   `json:"wind_speed"`
		WindDirection        float64   `json:"wind_direction"`
		AirPressure          float64   `json:"air_pressure"`
		Humidity             int       `json:"humidity"`
		Visibility           float64   `json:"visibility"`
		Predictability       int       `json:"predictability"`
	} `json:"consolidated_weather"`
	Time         time.Time `json:"time"`
	SunRise      time.Time `json:"sun_rise"`
	SunSet       time.Time `json:"sun_set"`
	TimezoneName string    `json:"timezone_name"`
	Parent       struct {
		Title        string `json:"title"`
		LocationType string `json:"location_type"`
		Woeid        int    `json:"woeid"`
		LattLong     string `json:"latt_long"`
	} `json:"parent"`
	Sources []struct {
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
		CrawlRate int    `json:"crawl_rate"`
	} `json:"sources"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
	Timezone     string `json:"timezone"`
}

func main() {
	port := os.Getenv("PORT")
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/search", searchHandler)
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.Handle("/search", r)
	http.ListenAndServe(":" + port, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		context := map[string]string{
			"author": "Juan Chacon",
		}
		templates.ExecuteTemplate(w, "index.html", context)

	case "POST":
		http.Redirect(w, r, "/search", 302)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		query := strings.ReplaceAll(r.FormValue("query"), " ", "+") //replace all spaces for + sign
		id, err := getWoeid(query)
		if err != nil {
			fmt.Println(err)
			templates.ExecuteTemplate(w, "error.html", nil)
		} else {
			weather := getWeather(id)
			templates.ExecuteTemplate(w, "search.html", weather)
		}
	}
}

func getWoeid(city string) (int, error) { //calls the api to get the woeid using the string the user provided
	fmt.Println("1. Performing Http Get...")
	url := fmt.Sprintf("https://www.metaweather.com/api/location/search/?query=%s", city)
	fmt.Println("2. Calling URL: " + url)
	res, erro := http.Get(url)
	if erro != nil {
		log.Fatalln(erro)
	}

	defer res.Body.Close()
	var woeid woeidResponse
	fmt.Println("3. Getting woeid")
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &woeid)

	if len(woeid) != 0 {
		fmt.Println(woeid[0].Woeid)
		id := woeid[0].Woeid
		return id, nil
	} else {
		return 0, errors.New("Error: Woeid not found")
	}

}

func getWeather(id int) weatherResponse {
	url := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", id)
	fmt.Println("4. Calling URL: " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Todo struct
	var weather weatherResponse
	json.Unmarshal(bodyBytes, &weather)
	fmt.Println("Success!!")
	return weather
}
