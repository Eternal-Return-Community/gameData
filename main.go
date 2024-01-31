package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type gameInfo struct {
	Url string
}

const BASE = "http://cdn.eternalreturn.io"

var format = fmt.Sprintf

func main() {
	list := gamePath()
	for _, v := range list {
		info := gameInfo{Url: v}
		body := info.Get()
		download(format("%s/%s", strings.Split(v, "/")[0], body), body)
	}
}

func gamePath() [2]string {
	return [2]string{"l10n/l10n-Korean-steam.txt", "gameDb/gamedata-steam.txt"}
}

func (g *gameInfo) Get() string {
	response, err := http.Get(url(g.Url))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func download(endpoint string, fileName string) {

	client := gameInfo{Url: endpoint}
	body := client.Get()

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s foi baixado com sucesso! \n", fileName)
}

func url(endpoint string) string {
	return format("%s/%s", BASE, endpoint)
}
