package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)


func getKey(token *jwt.Token, url string) (interface{}, error) {
	keySet, err := jwk.Fetch(context.Background(), url)
	if err != nil {
		return nil, err
	}

	var (
		keyID string
		ok    bool
	)
	// BA IdPの仕様 Header.kidではなく、Claims.jwkを参照する
	keyID, ok = token.Claims.(jwt.MapClaims)["jwk"].(string)
	if !ok {
		// なければ通常通りHeader.kidを参照する
		keyID, ok = token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("expecting JWT header to have a key ID in the kid field")
		}
	}

	key, exist := keySet.LookupKeyID(keyID)
	if !exist {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}
	return pubkey, nil
}
