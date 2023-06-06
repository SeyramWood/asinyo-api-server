package jwt

import (
	"crypto"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateToken(ttl time.Duration, session any) (string, error) {
	now := time.Now().UTC()
	// Create the Claims
	claims := jwt.MapClaims{
		"data": session,
		"exp":  now.Add(ttl).Unix(),
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString(j.privateKey())
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %v", err)
	}
	return tokenString, nil
}
func (j *JWT) ValidateToken(token string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(
		token, func(jwtToken *jwt.Token) (any, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}
			return j.publicKey(), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not cast claims")
	}
	return claims, nil
}

func (j *JWT) Command() *JWT {
	if j.privateKey() != nil && j.publicKey() != nil {
		return j
	}
	prikeyArgs := []string{"genrsa", "-out", "mnt/cert/jwt_rsa", "4096"}
	pubkeyArgs := []string{"rsa", "-in", "mnt/cert/jwt_rsa", "-pubout", "-out", "mnt/cert/jwt_rsa.pub"}
	cmd := exec.Command("openssl", prikeyArgs...) // pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout
	// Run still runs the command and waits for completion
	// but the output is instantly piped to Stdout
	if err := cmd.Run(); err != nil {
		log.Panicln("could not run command private key: ", err)
	}
	cmd = exec.Command("openssl", pubkeyArgs...) // pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Panicln("could not run command public key: ", err)
	}
	return j
}
func (j *JWT) privateKey() crypto.PrivateKey {
	keyByte, err := os.ReadFile("mnt/cert/jwt_rsa")
	if err != nil {
		log.Panicln("could not read private key: ", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyByte)
	if err != nil {
		log.Panicln("could not parse private key: ", err)
	}
	return key
}
func (j *JWT) publicKey() crypto.PublicKey {
	keyByte, err := os.ReadFile("mnt/cert/jwt_rsa.pub")
	if err != nil {
		log.Panicln("could not read public key: ", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyByte)
	if err != nil {
		log.Panicln("could not parse public key: ", err)
	}
	return key
}
