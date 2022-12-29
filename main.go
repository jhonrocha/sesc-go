package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const DAYS = 90
const url = "https://portalturismoapi.sescgo.com.br/api/v1/reservas/disponibilidade/%s/%s/1,1,/null/N/1/1/%%7B%%7D/%%7B%%7D?SescUnidade=1&ReservaId=0"

func main() {
	// start := time.Now()
  start := time.Date(2023, 04, 01, 0, 0, 0, 0, time.UTC)
	for i := 0; i < DAYS; i++ {
		end := start.AddDate(0, 0, 2)
		fullUrl := fmt.Sprintf(url, start.Format("2006-01-02"), end.Format("2006-01-02"))
		resp, err := http.Get(fullUrl)
		if err != nil {
			log.Fatalln(err)
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		body := string(bodyBytes)
		var available string
		if strings.Contains(body, "Não existe disponibilidade") {
			available = "😭 😭"
		} else {
			available = "😎😎😎😎😎😎😎😎😎"
		}
		fmt.Printf("%s - %s: %s\n", start.Format("Mon 2006-01-02"), end.Format("Mon 2006-01-02"), available)
		start = start.AddDate(0, 0, 1)
	}
}
