package main

import ("fmt"; "log"; "net/http"; "io"; "encoding/json")
 
type ISSLocation struct {
	Timestamp   int64 `json:"timestamp"`
	ISSPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message string `json:"message"`
}


func main() {
	resp, err := http.Get("http://api.open-notify.org/iss-now.json")

	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	// Lendo o corpo da resposta
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
	}
	fmt.Println("Corpo da resposta:", string(body))

	// Decodificando o JSON
	var location ISSLocation

	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	// Exibindo as informações úteis
	fmt.Println("Localização atual da ISS:")
	fmt.Printf("Latitude: %s\n", location.ISSPosition.Latitude)
	fmt.Printf("Longitude: %s\n", location.ISSPosition.Longitude)
}