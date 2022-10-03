package fileload

import (
	"embed"
	"encoding/csv"
	"io"
	"log"
	"sync"

	"github.com/luispfcanales/api-phrases/internal/core/domain"
)

//go:embed storage.csv
var content embed.FS

var once sync.Once
var listPhrases domain.Phrases

// Load function read phrases and save in memory
func Load(phrases chan<- domain.Phrases) {
	once.Do(func() {
		reciver := make(chan []string)

		file, err := content.Open("storage.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		csvReader := csv.NewReader(file)
		csvReader.Read()

		go savePhrases(reciver)

		for {
			rec, err := csvReader.Read()
			if err == io.EOF {
				close(reciver)
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			reciver <- rec
		}
	})

	phrases <- listPhrases
	close(phrases)
}

func savePhrases(reciver chan []string) {
	for {
		select {
		case phrase, ok := <-reciver:
			if !ok {
				return
			}
			p := domain.Phrase{
				Author:  phrase[1],
				Message: phrase[0],
			}
			listPhrases = append(listPhrases, p)
		}
	}
}
