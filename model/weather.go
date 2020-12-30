package model

type OpenWeather struct {
	Weather    []weather `json:"weather"`
	Info       main      `json:"main"`
	Visibility uint64    `json:"visibility"`
	Wind       wind      `json:"wind"`
}

type weather struct {
	Id          uint64 `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  uint64  `json:"pressure"`
	Humidity  uint8   `json:"humidity"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   uint16  `json:"deg"`
}
