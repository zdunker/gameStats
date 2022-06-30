package data

import (
	"context"

	"github.com/uptrace/bun"
)

type match struct {
	bun.BaseModel `bun:"table:dota_match"`

	MatchID   int64          `bun:"match_id"`
	MatchData map[string]any `bun:"match_data"`
}

type dotaMatchService interface {
	dotaMatchReadOnlyService
	Write(ctx context.Context, matchID int64, data map[string]any) error
}

type dotaMatchReadOnlyService interface { // read-only interface
	Read(ctx context.Context, matchID int64) (match, error)
}

var _ dotaMatchReadOnlyService = (*dotaMatchReadOnlyHelper)(nil)

type dotaMatchReadOnlyHelper struct { // a postgres helper for reading dota matches data
	client *bun.DB // better a read-only client
}

func (h dotaMatchReadOnlyHelper) Read(ctx context.Context, matchID int64) (match, error) {
	m := new(match)
	_, err := h.client.NewSelect().
		Model(m).
		Where("match_id = ?", matchID).
		Exec(ctx)
	return *m, err
}

var _ dotaMatchService = (*dotaMatchRWHelper)(nil)

type dotaMatchRWHelper struct { // a postgres helper for writing dota matches data
	client *bun.DB
}

func (h dotaMatchRWHelper) Write(ctx context.Context, matchID int64, data map[string]any) error {
	m := &match{
		MatchID:   matchID,
		MatchData: data,
	}
	_, err := h.client.NewInsert().
		Model(m).
		On("CONFLICT (match_id) DO UPDATE").
		Set("match_data = EXCLUDED.match_data").
		Exec(ctx)
	return err
}

func (h dotaMatchRWHelper) Read(ctx context.Context, matchID int64) (match, error) {
	m := new(match)
	_, err := h.client.NewSelect().
		Model(m).
		Where("match_id = ?", matchID).
		Exec(ctx)
	return *m, err
}

func NewDotaMatchReadOnlyService(client *bun.DB) dotaMatchReadOnlyService {
	return &dotaMatchReadOnlyHelper{client: client}
}

func NewDotaMatchRWService(client *bun.DB) dotaMatchService {
	return &dotaMatchRWHelper{client: client}
}
