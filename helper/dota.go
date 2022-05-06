package helper

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/zdunker/gameStats/config"
	"github.com/zdunker/gameStats/utility"
)

type dotaRequest struct {
	req *http.Request
}

func MakeDotaURL(paths ...string) string {
	dotaConf := config.GetConfig().Dota()
	dotaAPIPrefix := dotaConf.DotaAPIPrefix
	builder := strings.Builder{}
	builder.WriteString(dotaAPIPrefix)
	for _, path := range paths {
		if !strings.HasPrefix(path, "/") {
			builder.WriteString("/")
		}
		if strings.HasSuffix(path, "/") {
			builder.WriteString(path[:len(path)-1])
		} else {
			builder.WriteString(path)
		}
	}
	builder.WriteString("?api_key=")
	builder.WriteString(dotaConf.DotaAPIKey)
	return builder.String()
}

func MakeDotaRequest(data map[string]any, method string, paths ...string) (req *dotaRequest, err error) {
	if method != http.MethodPost && method != http.MethodGet {
		return nil, errors.New("Method not supported")
	}
	var payload []byte
	if method == http.MethodPost && data != nil {
		payload, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	r, err := http.NewRequestWithContext(context.TODO(), method, MakeDotaURL(paths...), strings.NewReader(string(payload)))
	req = &dotaRequest{r}
	return req, err
}

func (req *dotaRequest) Do() (utility.Response, error) {
	return utility.Do(*req.req)
}
