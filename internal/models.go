package internal

type Data struct {
	Location Location
	Current Current
}

type Location struct {
	Name            string  `json:"name"`
	Region          string  `json:"region"`
	Country         string  `json:"country"`
	Lat             float64 `json:"lat"`
	Lon             float64 `json:"lon"`
	TZ_ID           string  `json:"tz_id"`
	Localtime_epoch float64 `json:"localtime_epoch"`
	Localtime       string  `json:"localtime"`
}

type Current struct {
	Last_updated_epoch int     `json:"last_updated_epoch"`
	Last_updated       string  `json:"last_updated"`
	Temp_c             float64 `json:"temp_c"`
	Temp_f             float64 `json:"temp_f"`
	Is_day             int     `json:"is_day"`
	Condition          Condition
	Wind_mph           float64 `json:"wind_mph"`
	Wind_kph           float64 `json:"wind_kph"`
	Wind_degree        int     `json:"wind_degree"`
	Wind_dir           string  `json:"wind_dir"`
	Pressure_mb        float64 `json:"pressure_mb"`
	Pressure_in        float64 `json:"pressure_in"`
	Precip_mm          float64 `json:"precip_mm"`
	Precip_in          float64 `json:"precip_in"`
	Humidity           int     `json:"humidity"`
	Cloud              int     `json:"cloud"`
	Feelslike_c        float64 `json:"feelslike_c"`
	Feelslike_f        float64 `json:"feelslike_f"`
	Vis_km             float64 `json:"vis_km"`
	Vis_miles          float64 `json:"vis_miles"`
	Uv                 float64 `json:"uv"`
	Gust_mph           float64 `json:"gust_mph"`
	Gust_kph           float64 `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
