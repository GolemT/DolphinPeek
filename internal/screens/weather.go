package screens

import "math/rand"

func GetWeatherInfo() string {
	// Random dolphin weather puns
	weatherPuns := []string{
		"Dolphinitely sunny today!",
		"Whale, it's looking cloudy...",
		"Sea-riously nice weather!",
		"Having a whale of a time with this weather!",
		"Pod-fect conditions outside!",
	}

	pun := weatherPuns[rand.Intn(len(weatherPuns))]

	return `
    \   /     ` + pun + `
     .-.      22°C
  ― (   ) ―   
     '-'      ↗ 15km/h
    /   \     10km visibility
              
Location: Tim's Ocean 🌊
Humidity: 65%
Pressure: 1013hPa

☀️ Today: Sunny, 18-24°C
🌤️ Tomorrow: Partly Cloudy

🐬 Dolphin Fact: Weather perfect 
   for jumping out of water!`
}
