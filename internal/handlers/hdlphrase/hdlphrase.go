package hdlphrase

import (
	"net/http"

	"github.com/luispfcanales/api-phrases/internal/core/domain"
	"github.com/luispfcanales/api-phrases/internal/core/ports"
	"github.com/luispfcanales/api-phrases/pkg/middleware"
	"github.com/luispfcanales/api-phrases/pkg/resmsg"
)

type phraseHDL struct {
	srv ports.SRVPhrase
}

// New return handler to phrases
func New(servicePhrase ports.SRVPhrase) *phraseHDL {
	return &phraseHDL{
		srv: servicePhrase,
	}
}

// Setup set handlers to managment http
func (hdl *phraseHDL) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/phrase", middleware.Cors(hdl.Phrase))
}

// Phrase return list phrases
func (hdl *phraseHDL) Phrase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	phrasesChan := make(chan domain.Phrases)

	go hdl.srv.Phrase(ctx, phrasesChan)

	select {
	case phrases := <-phrasesChan:
		res := resmsg.New()
		res.Data = phrases
		res.SendReponse(w)
		return
	}
}
