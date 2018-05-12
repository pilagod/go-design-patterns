package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type WeatherDataFetcher interface {
	GetByCityAndCountryCode(city, countryCode string) (WeatherData, error)
	GetByGeoCoordinates(lat, lon float32) (WeatherData, error)
}

type WeatherData struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type OpenWeatherMap struct {
	APIKey string
}

func (o *OpenWeatherMap) GetByCityAndCountryCode(city, countryCode string) (*WeatherData, error) {
	return o.doRequest(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", city, countryCode, o.APIKey))
}

func (o *OpenWeatherMap) GetByGeoCoordinates(lat, lon float32) (*WeatherData, error) {
	return o.doRequest(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, o.APIKey))
}

func (o *OpenWeatherMap) doRequest(uri string) (*WeatherData, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		bytes, errmsg := ioutil.ReadAll(res.Body)

		if errmsg == nil {
			errmsg = fmt.Errorf("%s", string(bytes))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was: %s", res.StatusCode, errmsg)
		return nil, errmsg
	}
	weather, err := o.parseResponse(res.Body)
	res.Body.Close()
	return weather, err
}

func (w *OpenWeatherMap) parseResponse(body io.Reader) (*WeatherData, error) {
	weatherData := &WeatherData{}
	err := json.NewDecoder(body).Decode(weatherData)

	if err != nil {
		return nil, err
	}
	return weatherData, nil
}
