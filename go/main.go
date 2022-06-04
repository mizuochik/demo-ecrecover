package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func recoverSignature(signatrue string, msg []byte) (string, error) {
	sig, err := hexutil.Decode(signatrue)
	if err != nil {
		return "", err
	}
	msg = accounts.TextHash(msg)
	sig[crypto.RecoveryIDOffset] -= 27
	pub, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return "", err
	}
	addr := crypto.PubkeyToAddress(*pub)
	return strings.ToLower(addr.Hex()), nil
}

func main() {
	pa, err := recoverSignature("0x317af8ad8c3b69b6bf9ca9e193565105726e49ff3c4b2db55bd178d1b26b367126e751f10888319824f001295912a316fd3d565d765acaf3a695846b0420742f1c", []byte("hello"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("public addr=%s\n", pa)
}
