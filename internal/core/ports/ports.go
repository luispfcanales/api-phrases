package ports

import (
	"context"

	"github.com/luispfcanales/api-phrases/internal/core/domain"
)

// SRVPhrase is PORT
type SRVPhrase interface {
	Phrase(ctx context.Context, chanPhrase chan domain.Phrases)
}
