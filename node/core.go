package node

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type (
	BlockChain struct {
		Chain           []interface{}
		LastTransaction []transaction
		Nodes           []string
		HashTarget      string
	}

	Block struct {
		Index       int
		Time        time.Time
		Transaction []transaction
		Nonce       int
		HashBlock   string
	}

	transaction struct {
		Sender   string
		Receiver string
		Amount   float64
	}
)

func New(chain *BlockChain) *BlockChain {

	trans := transaction{
		Sender:   "",
		Receiver: "",
		Amount:   0,
	}
	initialHash := chain.BlockHashing([]transaction{trans})

	chain.BlockAppend(chain.PoW(0, initialHash, trans), initialHash)
	return chain
}

func (b *BlockChain) BlockHashing(trans []transaction) string {
	return createHash(fmt.Sprintf("%v", trans))
}

func (b *BlockChain) BlockAppend(nonce int, hashing string) {
	block := Block{
		Index:       len(b.Chain),
		Time:        time.Now(),
		Transaction: b.LastTransaction,
		Nonce:       nonce,
		HashBlock:   hashing,
	}
	b.Chain = append(b.Chain, block)
	b.LastTransaction = []transaction{}
}

func (b *BlockChain) PoW(idx int, hashing string, tran interface{}) /* Proof of work */ int {
	nonce := 0
	for !b.ValPoW(idx, hashing, tran, nonce) {
		nonce++
	}
	return nonce
}

func (b *BlockChain) ValPoW(idx int, hashing string, trans interface{}, nonce int) /* Validation Proof of work */ bool {
	content := fmt.Sprintf("%v-%v-%v-%v", idx, hashing, trans, nonce)
	strHash := createHash(content)
	return strHash[0:4] == b.HashTarget
}

func (b *BlockChain) AppendTransactions(sender, receiver string, amount float64) {
	trans := transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
	b.LastTransaction = append(b.LastTransaction, trans)
}

//private function
func createHash(input string) string {
	data := []byte(input)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash[:])
}
