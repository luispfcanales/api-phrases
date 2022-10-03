package srvphrase

import (
	"context"

	"github.com/luispfcanales/api-phrases/internal/core/domain"
	"github.com/luispfcanales/api-phrases/pkg/fileload"
)

type phraseSRV struct{}

func New() *phraseSRV {
	return &phraseSRV{}
}

func (s *phraseSRV) Phrase(ctx context.Context, chanPhrase chan domain.Phrases) {
	go fileload.Load(chanPhrase)
	select {
	case <-ctx.Done():
		return
	}
}
