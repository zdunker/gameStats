package views

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/zdunker/gameStats/helper"
	"github.com/zdunker/webframe"
)

const (
	dotaAPIMatchPrefix = "/matches"
)

type dotaMatch struct {
	MatchID int64 `json:"match_id"`

	FirstBloodTime int64 `json:"first_blood_time"`
	StartTime      int64 `json:"start_time"`

	DireScore    int64 `json:"dire_score"`
	RadiantScore int64 `json:"radiant_score"`
}

func (m dotaMatch) toMap() map[string]interface{} {
	return map[string]interface{}{
		"match_id":         m.MatchID,
		"first_blood_time": m.FirstBloodTime,
		"start_time":       time.Unix(m.StartTime, 0),
		"dire_score":       m.DireScore,
		"radiant_score":    m.RadiantScore,
	}
}

func GetMatch(c *webframe.Context) {
	matchID, exists := c.Query("match_id")
	if !exists {
		c.ErrorResponse(http.StatusBadRequest, errors.New("match_id not provided"))
	}
	match, err := getMatch(matchID)
	if err != nil {
		c.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	c.JSONResponse(http.StatusOK, match.toMap())
}

func getMatch(matchID string) (dotaMatch, error) {
	req, err := helper.MakeDotaRequest(nil, http.MethodGet, dotaAPIMatchPrefix, matchID)
	if err != nil {
		return dotaMatch{}, err
	}
	resp, err := req.Do()
	if err != nil {
		return dotaMatch{}, err
	}
	if !resp.Is200() {
		return dotaMatch{}, errors.New("request error")
	}
	var match dotaMatch
	err = json.Unmarshal(resp.Body(), &match)
	if err != nil {
		return dotaMatch{}, err
	}
	return match, nil
}
