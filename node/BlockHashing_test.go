package node

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChain_BlockHashing(t *testing.T) {
	type fields struct {
		Chain           []interface{}
		LastTransaction []transaction
		Nodes           []string
	}
	type args struct {
		trans []transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Hashing",
			fields: fields{
				Chain:           nil,
				LastTransaction: []transaction{},
				Nodes:           nil,
			},
			args: args{trans: []transaction{{
				Sender:   "a1",
				Receiver: "b2",
				Amount:   0.1,
			}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockChain{
				Chain:           tt.fields.Chain,
				LastTransaction: tt.fields.LastTransaction,
				Nodes:           tt.fields.Nodes,
			}
			strHash := b.BlockHashing(tt.args.trans)

			fmt.Println("Hash Result => ", strHash)

			assert.NotEmpty(t, strHash)
		})
	}
}
