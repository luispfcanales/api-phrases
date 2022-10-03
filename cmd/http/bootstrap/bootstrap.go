package bootstrap

import (
	"fmt"
	"net/http"
	"os"

	"github.com/luispfcanales/api-phrases/internal/core/services/srvphrase"
	"github.com/luispfcanales/api-phrases/internal/handlers/hdlphrase"
)

// RunHTTP server up
func RunHTTP() error {
	mux := http.NewServeMux()

	hdlPhrase := hdlphrase.New(srvphrase.New())
	hdlPhrase.Setup(mux)

	return http.ListenAndServe(getPort(), mux)
}

func getPort() string {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	return fmt.Sprintf(":%s", PORT)
}
