package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type PasetoMaker struct {
	privateKey string
}

func NewPasetoMaker(key string) (Maker, error) {
	maker := &PasetoMaker{
		privateKey: key,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", nil, err
	}

	token := paseto.NewToken()
	token.Set("username", payload.Username)
	token.Set("id", payload.ID)
	token.SetIssuedAt(payload.IssuedAt)
	token.SetExpiration(payload.ExpiredAt)

	sk, err := paseto.NewV4AsymmetricSecretKeyFromHex(maker.privateKey)

	if err != nil {
		return "", nil, err
	}

	tokenString := token.V4Sign(sk, nil)
	return tokenString, payload, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(maker.privateKey)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())

	implicit := []byte(nil)

	pToken, err := parser.ParseV4Public(privateKey.Public(), token, implicit)
	if err != nil {
		return nil, err
	}

	username, err := pToken.GetString("username")
	if err != nil {
		return nil, err
	}

	idStr, err := pToken.GetString("id")
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	issuedAt, err := pToken.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	expiredAt, err := pToken.GetExpiration()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	return payload, nil
}
