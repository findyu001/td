// Code generated by rsagen, DO NOT EDIT.

package example

import (
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotd/td/crypto"
)

func TestFingerprintEmptyPK(t *testing.T) {
	fingerprint := func(pubkey *rsa.PublicKey) string {
		return fmt.Sprintf("%08x", uint64(crypto.RSAFingerprint(pubkey)))
	}
	expected := []string{}
	assert.Len(t, EmptyPK, len(expected))
	for i, pubkey := range EmptyPK {
		assert.Equal(t, expected[i], fingerprint(pubkey))
	}
}
