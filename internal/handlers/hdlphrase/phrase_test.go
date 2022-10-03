package hdlphrase_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luispfcanales/api-phrases/internal/core/domain"
	"github.com/luispfcanales/api-phrases/internal/core/services/srvphrase"
	"github.com/luispfcanales/api-phrases/internal/handlers/hdlphrase"
)

func TestPhrases(t *testing.T) {
	t.Run("GET PHRASES", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/phrase", nil)
		response := httptest.NewRecorder()

		hdlmooc := hdlphrase.New(srvphrase.New())
		hdlmooc.Phrase(response, request)

		var phrases domain.Phrases
		err := json.Unmarshal(response.Body.Bytes(), &phrases)
		if err != nil {
			log.Fatal(err)
		}
		if len(phrases) == 0 {
			t.Errorf("se esperaba obtener una lista de phrases")
		}
	})
}
