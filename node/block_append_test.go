package node

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBlockChain_BlockAppend(t *testing.T) {
	type fields struct {
		Chain           []interface{}
		LastTransaction transaction
		Nodes           []string
	}
	type args struct {
		nonce   int
		hashing string
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		{
			name: "Test Append block",
			fields: fields{
				Chain: nil,
				LastTransaction: transaction{
					Sender:   "a",
					Receiver: "b",
					Amount:   0.25,
				},
				Nodes: nil,
			},
			args: []args{
				{
					nonce:   1,
					hashing: "abz",
				},
				{
					nonce:   2,
					hashing: "ab2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockChain{
				Chain:           tt.fields.Chain,
				LastTransaction: tt.fields.LastTransaction,
				Nodes:           tt.fields.Nodes,
			}
			for _, arg := range tt.args {
				b.BlockAppend(arg.nonce, arg.hashing)
			}

			prettyJSON, err := json.MarshalIndent(b.Chain, "", "\t")
			if err != nil {
				log.Fatal("Failed to generate json", err)
			}
			fmt.Printf("%s\n", string(prettyJSON))

			assert.Equal(t, 2, len(b.Chain))

		})
	}
}
