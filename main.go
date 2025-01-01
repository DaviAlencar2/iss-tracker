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
	fmt.Printf("Longitude: %s\n\n", location.ISSPosition.Longitude)

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