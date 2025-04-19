// Package token provides the implementation of the token service interface for the authentication service.
package token

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/service/token"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
)

type jwtTokenService struct {
	priv       *rsa.PrivateKey
	pubJwk     string
	accessTTL  time.Duration
	refreshTTL time.Duration
	issuer     string
}

func NewTokenService(pemKey string, accessTTL, refreshTTL time.Duration, issuer string) (
	token.Service,
	error,
) {
	priv, err := parseRsaPrivateKey(pemKey)
	if err != nil {
		return nil, err
	}
	pubJwk, err := buildJWK(priv)
	if err != nil {
		return nil, err
	}
	return &jwtTokenService{
		priv:       priv,
		pubJwk:     pubJwk,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		issuer:     issuer,
	}, nil
}

// -------------------------------------------------------------------
// Issue: Access + RefreshToken
// -------------------------------------------------------------------
func (s *jwtTokenService) Issue(id *model.AuthIdentityModel, deviceID string) (
	*model.TokensModel,
	error,
) {
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"sub":   id.ID,
		"tid":   id.OrganizationID,
		"email": id.Email,
		"iat":   now.Unix(),
		"exp":   now.Add(s.accessTTL).Unix(),
		"iss":   s.issuer,
	}
	j := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	j.Header["kid"] = "key‑1" // rota claves → cambia kid
	access, err := j.SignedString(s.priv)
	if err != nil {
		return nil, err
	}

	refreshUUID := uuid.NewString() // 128‑bit random opaque
	hash := sha256.Sum256([]byte(refreshUUID))

	return &model.TokensModel{
		Access:       access,
		AccessExp:    now.Add(s.accessTTL),
		RefreshToken: refreshUUID,
		RefreshHash:  hash[:],
	}, nil
}

// -------------------------------------------------------------------
// VerifyRefresh: devuelve el SHA‑256 del token si es formato válido
// -------------------------------------------------------------------
func (s *jwtTokenService) VerifyRefresh(refresh string) ([]byte, error) {
	// Opcional: valida que sea UUID
	if _, err := uuid.Parse(refresh); err != nil {
		return nil, errors.New("refresh token format invalid")
	}
	hash := sha256.Sum256([]byte(refresh))
	return hash[:], nil
}

// -------------------------------------------------------------------
// PublicJWK expone la clave pública en formato JWK (para /.well‑known)
// -------------------------------------------------------------------
func (s *jwtTokenService) PublicJWK() string { return s.pubJwk }

// -------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------
func parseRsaPrivateKey(pemKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("invalid pem block")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func buildJWK(priv *rsa.PrivateKey) (string, error) {
	pub := priv.Public().(*rsa.PublicKey)
	jwk := map[string]string{
		"kty": "RSA",
		"alg": "RS256",
		"use": "sig",
		"kid": "key-1",
		"n":   base64.RawURLEncoding.EncodeToString(pub.N.Bytes()),
		"e":   base64.RawURLEncoding.EncodeToString(bigIntToBytes(pub.E)),
	}
	b, err := json.Marshal(jwk)
	return string(b), err
}

func bigIntToBytes(x int) []byte {
	return big.NewInt(int64(x)).Bytes()
}
