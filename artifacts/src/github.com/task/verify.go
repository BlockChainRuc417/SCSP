
package main


import (
	"fmt"
  "crypto/ecdsa"
  "crypto/sha256"
  "crypto/x509"
  "encoding/pem"
  "math/big"
  "hash"
  "io"
)


// ==================================================
//      getPublicKey - Get publicKey from cert
// ==================================================
func getPublicKey(certPEM string) (*ecdsa.PublicKey) {
	var null *ecdsa.PublicKey

	block, _ := pem.Decode([]byte(certPEM))
  if block == nil {
    return null
  }

  cert, err := x509.ParseCertificate(block.Bytes)
  if err != nil {
    return null
  }

  publicKey := cert.PublicKey.(*ecdsa.PublicKey)
	return publicKey
}


// ====================================================================
//      verifySignature - Verify between signature and publicKey
// ====================================================================
func verifySignature(publicKey *ecdsa.PublicKey, r string, s string, act string) bool {
	bigr := new(big.Int)
	nr, _ := bigr.SetString(r, 10)

	bigs := new(big.Int)
	ns, _ := bigs.SetString(s, 10)

	var h hash.Hash
	h = sha256.New()
	io.WriteString(h, act)
	nsig := h.Sum(nil)

	verify := ecdsa.Verify(publicKey, nsig, nr, ns)
	fmt.Println(verify) // should be true

	return verify
}
