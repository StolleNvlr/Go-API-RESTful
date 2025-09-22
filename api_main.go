package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TemperatureRequest struct{
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}


type TemperatureResponse struct {
    Celsius    float64 `json:"celsius"`
    Fahrenheit float64 `json:"fahrenheit"`
    Kelvin     float64 `json:"kelvin"`
}


func TemperatureHandler(value float64, unit string) (float64, float64, float64, error) {
	switch unit {
	case "celsius":
		c := value
		f := (value * 9.0 / 5.0) + 32.0
        k := value + 273.15
        return c, f, k, nil
	case "fahrenheit":
		c := (value - 32.0) * 5.0 / 9.0
		f := value
		k := c + 273.15
		return c, f, k, nil
	case "kelvin":
		c := value - 273.15
		f := (c * 9.0 / 5.0) + 32.0
		k := value


		return c, f, k, nil
	default:
		return 0,0,0, fmt.Errorf("invalid unit")
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}
	var req TemperatureRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error":"Invalid request data."}`, http.StatusBadRequest)
		return
	}
	c,f,k, convErr := TemperatureHandler(req.Value, req.Unit)
	if convErr != nil {
		http.Error(w, `{"error":"Invalid unit. Use celsius, fahrenheit or kelvin."}`, http.StatusBadRequest)
		return
	}

	resp := TemperatureResponse{Celsius: c, Fahrenheit: f, Kelvin: k}
	json.NewEncoder(w).Encode(resp)
}

func main() {
    http.HandleFunc("/convert", requestHandler)
    log.Println("Servidor rodando em http://localhost:8080/convert")
    log.Fatal(http.ListenAndServe(":8080", nil))
}