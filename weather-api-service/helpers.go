package main

type WeatherRes struct {
	Address         string           `json:"address"`
	Latitude        float64          `json:"latitude"`
	Longitude       float64          `json:"longitude"`
	ResolvedAddress string           `json:"resolvedAddress"`
	Day             []WeatherEffects `json:"days"`
}

type WeatherEffects struct {
	TempMax           float32 `json:"tempmax"`
	TempMin           float32 `json:"tempmin"`
	Temp              float32 `json:"temp"`
	FeelsLike         float32 `json:"feelslike"`
	Humidity          float32 `json:"humidity"`
	PrecipitationProb float32 `json:"precipcover"`
	SunRise           string  `json:"sunrise"`
	SunSet            string  `json:"sunset"`
	Conditions        string  `json:"conditions"`
	Description       string  `json:"description"`
	DateTime          string  `json:"datetime"`
}
