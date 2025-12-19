package token

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/stretchr/testify/require"
)

func newTestPasetoMaker(t *testing.T) Maker {
	sk := paseto.NewV4AsymmetricSecretKey()
	privateHex := sk.ExportHex()

	maker, err := NewPasetoMaker(privateHex)
	require.NoError(t, err)

	return maker
}

func TestCreateToken(t *testing.T) {
	maker := newTestPasetoMaker(t)

	username := "user123"
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotNil(t, payload)

	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, time.Now(), payload.IssuedAt, time.Second)
	require.WithinDuration(t, time.Now().Add(duration), payload.ExpiredAt, time.Second)
}

func TestVerifyToken(t *testing.T) {
	maker := newTestPasetoMaker(t)

	username := "user123"
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)

	verifiedPayload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, verifiedPayload)

	require.Equal(t, payload.ID, verifiedPayload.ID)
	require.Equal(t, payload.Username, verifiedPayload.Username)
	require.WithinDuration(t, payload.IssuedAt, verifiedPayload.IssuedAt, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, verifiedPayload.ExpiredAt, time.Second)
}

func TestVerifyToken_Expired(t *testing.T) {
	maker := newTestPasetoMaker(t)

	token, _, err := maker.CreateToken("user123", -time.Minute)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestVerifyToken_InvalidToken(t *testing.T) {
	maker := newTestPasetoMaker(t)

	payload, err := maker.VerifyToken("this.is.not.a.paseto.token")

	require.Error(t, err)
	require.Nil(t, payload)
}

func TestVerifyToken_WrongKey(t *testing.T) {
	// Maker 1
	sk1 := paseto.NewV4AsymmetricSecretKey()
	maker1, _ := NewPasetoMaker(sk1.ExportHex())

	// Maker 2 (clave distinta)
	sk2 := paseto.NewV4AsymmetricSecretKey()
	maker2, _ := NewPasetoMaker(sk2.ExportHex())

	token, _, err := maker1.CreateToken("user123", time.Minute)
	require.NoError(t, err)

	payload, err := maker2.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
