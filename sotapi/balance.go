package sotapi

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/wneessen/sotbot/httpclient"
	"net/http"
)

type UserBalance struct {
	GamerTag     string `json:"gamertag"`
	Title        string `json:"title"`
	Doubloons    int    `json:"doubloons"`
	Gold         int    `json:"gold"`
	AncientCoins int    `json:"ancientCoins"`
}

func GetBalance(hc *http.Client, rc string) (UserBalance, error) {
	l := log.WithFields(log.Fields{
		"action": "sotapi.SetRatCookie",
	})
	var userBalance UserBalance
	apiUrl := "https://www.seaofthieves.com/api/profilev2/balance"

	l.Debugf("Fetching balance from API...")
	httpResp, err := httpclient.HttpReqGet(apiUrl, hc, rc, "")
	if err != nil {
		return userBalance, err
	}
	if err := json.Unmarshal(httpResp, &userBalance); err != nil {
		l.Errorf("Failed to unmarshal API response: %v", err)
		return userBalance, err
	}

	return userBalance, nil
}
