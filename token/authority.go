package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"
)

type Authority struct {
	secret []byte
}

func New(secret string) Authority {
	return Authority{
		secret: []byte(secret),
	}
}

func (c *Authority) GetToken(guildID, channelID string) (string, error) {
	var h, err = c.encodeToken(guildID, channelID)
	if err != nil {
		return "", err
	}

	sha := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s.%s.%s", guildID, channelID, sha), nil
}

func (c *Authority) CheckToken(token string) error {
	var parts = strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("token parts invalid length")
	}

	msgHmac, err := c.encodeToken(parts[0], parts[1])
	if err != nil {
		return err
	}

	expectedHmac, err := hex.DecodeString(parts[2])
	if err != nil {
		return err
	}

	if !hmac.Equal(msgHmac.Sum(nil), expectedHmac) {
		return errors.New("signature invalid")
	}

	return nil
}

func (c *Authority) encodeToken(guildID, channelID string) (hash.Hash, error) {
	h := hmac.New(sha256.New, c.secret)

	_, err := h.Write([]byte(guildID + "." + channelID))
	if err != nil {
		return nil, err
	}

	return h, nil
}
