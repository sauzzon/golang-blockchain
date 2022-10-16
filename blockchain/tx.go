package blockchain

import (
	"bytes"

	"github.com/sauzzon/golang-blockchain/wallet"
)

// input contain pub key not hashed along with signature of the transaction
// out contain hash of the pub key

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID        []byte // references older output
	Out       int    //index of output if there are many outputs
	Signature []byte
	PubKey    []byte
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// to check if the account i.e data owns the info inside the output which is referenced by the input
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// to check if the account i.e data owns info inside output
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
