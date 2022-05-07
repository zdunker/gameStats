package data

import "github.com/go-pg/pg/v10"

type match struct {
	MatchID   int64          `pg:"match_id"`
	MatchData map[string]any `pg:"match_data"`
}

type dotaMatchService interface {
	dotaMatchReadOnlyService
	Write(matchID int64, data map[string]any) error
}

type dotaMatchReadOnlyService interface { // read-only interface
	Read(matchID int64) (match, error)
}

var _ dotaMatchReadOnlyService = (*dotaMatchReadOnlyPGHelper)(nil)

type dotaMatchReadOnlyPGHelper struct { // a postgres helper for reading dota matches data
	client *pg.DB // better a read-only client
}

func (h dotaMatchReadOnlyPGHelper) Read(matchID int64) (match, error) {
	m := new(match)
	err := h.client.Model(m).Where("match_id = ?", matchID).Select()
	return *m, err
}

var _ dotaMatchService = (*dotaMatchRWPGHelper)(nil)

type dotaMatchRWPGHelper struct { // a postgres helper for writing dota matches data
	client *pg.DB
}

func (h dotaMatchRWPGHelper) Write(matchID int64, data map[string]any) error {
	m := &match{
		MatchID:   matchID,
		MatchData: data,
	}
	_, err := h.client.Model(m).WherePK().OnConflict("DO NOTHING").Insert()
	return err
}

func (h dotaMatchRWPGHelper) Read(matchID int64) (match, error) {
	m := new(match)
	err := h.client.Model(m).Where("match_id = ?", matchID).Select()
	return *m, err
}

func NewDotaMatchReadOnlyService(client *pg.DB) dotaMatchReadOnlyService {
	return &dotaMatchReadOnlyPGHelper{client: client}
}

func NewDotaMatchRWService(client *pg.DB) dotaMatchService {
	return &dotaMatchRWPGHelper{client: client}
}
