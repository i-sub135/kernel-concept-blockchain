package node

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type (
	BlockChain struct {
		Chain           []interface{}
		LastTransaction transaction
		Nodes           []string
		HashTarget      string
	}

	Block struct {
		Index       int
		Time        time.Time
		Transaction transaction
		Nonce       int
		HashBlock   string
	}

	transaction struct {
		Sender   string
		Receiver string
		Amount   float64
	}
)

func (b *BlockChain) New() {

}

func (b *BlockChain) BlockHashing(trans transaction) string {
	return createHash(fmt.Sprintf("%v", trans))
}

func (b *BlockChain) BlockAppend(nonce int, hashing string) {
	block := Block{
		Index:       len(b.Chain) + 1,
		Time:        time.Now(),
		Transaction: b.LastTransaction,
		Nonce:       nonce,
		HashBlock:   hashing,
	}
	b.Chain = append(b.Chain, block)
}

func (b *BlockChain) PoW(idx int, hashing string, tran transaction) /* Proof of work */ int {
	nonce := 0
	for !b.ValPoW(idx, hashing, tran, nonce) {
		nonce = nonce + 1
	}
	return nonce
}

func (b *BlockChain) ValPoW(idx int, hashing string, trans transaction, nonce int) /* Validation Proof of work */ bool {
	content := fmt.Sprintf("%v-%v-%v-%v", idx, hashing, trans, nonce)
	strHash := createHash(content)
	return strHash[0:4] == b.HashTarget
}

//private function

func createHash(input string) string {
	data := []byte(input)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash[:])
}
