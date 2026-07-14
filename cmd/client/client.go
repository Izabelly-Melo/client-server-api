package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Print("Timeout: ", err)
			return
		}
		log.Print(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	var cotacao CotacaoResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	_, err = f.WriteString(fmt.Sprintf("Dólar: %v", cotacao.Bid))
	if err != nil {
		log.Println("Error:", err)
		return
	}
	f.Close()

}
