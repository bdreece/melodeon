package rooms

import (
	"encoding/json"
	"time"

	"github.com/bdreece/melodeon/pkg/spotify"
)

type Room struct {
	name         string
	accessToken  string
	refreshToken string
	expiration   time.Time
}

type roomJson struct {
	Name         string    `json:"name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiration   time.Time `json:"expiration"`
}

func (r Room) Name() string { return r.name }

func (r Room) Token() *spotify.Token {
	return &spotify.Token{
		AccessToken:  r.accessToken,
		RefreshToken: r.refreshToken,
		ExpiresIn:    r.expiration,
	}
}

func (r *Room) UnmarshalJSON(data []byte) error {
	var value roomJson
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	r.name = value.Name
	r.accessToken = value.AccessToken
	r.refreshToken = value.RefreshToken
	r.expiration = value.Expiration

	return nil
}

func (r Room) MarshalJSON() ([]byte, error) {
	return json.Marshal(roomJson{
		Name:         r.name,
		AccessToken:  r.accessToken,
		RefreshToken: r.refreshToken,
		Expiration:   r.expiration,
	})
}

func NewRoom(name string, token *spotify.Token) *Room {
	return &Room{
		name:         name,
		accessToken:  token.AccessToken,
		refreshToken: token.RefreshToken,
		expiration:   token.ExpiresIn,
	}
}
