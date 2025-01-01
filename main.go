package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Estruturas para armazenar os dados da API
type ISSLocation struct {
	Timestamp   int64 `json:"timestamp"`
	ISSPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message string `json:"message"`
}

type Astronauts struct {
	Number  int `json:"number"`
	People  []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
	Message string `json:"message"`
}

func main() {
	// Buscar localização da ISS
	location := fetchISSLocation()
	fmt.Println("Localização atual da ISS:")
	fmt.Printf("Latitude: %s\n", location.ISSPosition.Latitude)
	fmt.Printf("Longitude: %s\n", location.ISSPosition.Longitude)
	fmt.Printf("Detalhes: %s\n\n", fetchLocationDetails(location.ISSPosition.Latitude, location.ISSPosition.Longitude))

	// Buscar astronautas no espaço
	astronauts := fetchAstronauts()
	fmt.Println("Astronautas no espaço:")
	fmt.Printf("Total: %d\n", astronauts.Number)
	for _, person := range astronauts.People {
		fmt.Printf("Nome: %s | Nave: %s\n", person.Name, person.Craft)
	}
}

// Função para buscar a localização da ISS
func fetchISSLocation() ISSLocation {
	url := "http://api.open-notify.org/iss-now.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição para localização da ISS: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta da API de localização: %v", err)
	}

	var location ISSLocation
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON de localização: %v", err)
	}

	return location
}

// Função para buscar astronautas no espaço
func fetchAstronauts() Astronauts {
	url := "http://api.open-notify.org/astros.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição para astronautas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta da API de astronautas: %v", err)
	}

	var astronauts Astronauts
	err = json.Unmarshal(body, &astronauts)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON de astronautas: %v", err)
	}

	return astronauts
}

type LocationDetails struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	LocalityInfo struct {
		Informative []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"informative"`
	} `json:"localityInfo"`
}

func fetchLocationDetails(latitude, longitude string) string {
	url := fmt.Sprintf("https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%s&longitude=%s&localityLanguage=pt", latitude, longitude)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição para localização: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta da API: %v", err)
	}

	// fmt.Println("Resposta da API:", string(body)) // Debug da resposta

	var location LocationDetails
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON da localização: %v", err)
	}

	// Verificar se há informações disponíveis
	if len(location.LocalityInfo.Informative) > 0 {
		// Usar o primeiro elemento como exemplo
		info := location.LocalityInfo.Informative[0]
		return fmt.Sprintf("Nome: %s, Descrição: %s", info.Name, info.Description)
	}

	return "Informações não disponíveis para esta localização"
}
