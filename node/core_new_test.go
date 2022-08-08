package node

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		chain *BlockChain
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test new block",
			args: args{chain: &BlockChain{
				Chain:           []interface{}{},
				LastTransaction: []transaction{},
				Nodes:           []string{},
				HashTarget:      "0000",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := New(tt.args.chain)

			//create new transaction
			for i := 1; i < 10; i++ {
				block.AppendTransactions(fmt.Sprintf("a-0-%v", i), fmt.Sprintf("b-0-%v", i), 0.22*float64(i))
			}

			prettyJSON, _ := json.MarshalIndent(block.Chain, "", "\t")
			fmt.Println("len chain", len(block.Chain))
			fmt.Printf("Before mine => %s\n", string(prettyJSON))

			//miner process
			lastHast := block.BlockHashing(block.LastTransaction)
			idx := len(block.Chain)
			nonce := block.PoW(idx, lastHast, block.LastTransaction)
			block.BlockAppend(nonce, lastHast)

			fmt.Println("len chain", len(block.Chain))
			prettyJSON, _ = json.MarshalIndent(block.Chain, "", "\t")
			fmt.Printf("after mine => %s\n", string(prettyJSON))

		})
	}
}
