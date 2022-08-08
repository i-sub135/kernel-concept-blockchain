package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChain_AppendTransactions(t *testing.T) {
	type fields struct {
		Chain           []interface{}
		LastTransaction []transaction
		Nodes           []string
		HashTarget      string
	}
	type args struct {
		sender   string
		receiver string
		amount   float64
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		expect int
	}{
		{
			name: "Test Append last transactions",
			fields: fields{
				Chain:           nil,
				LastTransaction: []transaction{},
				Nodes:           nil,
				HashTarget:      "",
			},
			args: []args{
				{
					sender:   "a-1",
					receiver: "b-2",
					amount:   0.01,
				},
				{
					sender:   "a-3",
					receiver: "b-4",
					amount:   0.012,
				},
			},
			expect: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockChain{
				Chain:           tt.fields.Chain,
				LastTransaction: tt.fields.LastTransaction,
				Nodes:           tt.fields.Nodes,
				HashTarget:      tt.fields.HashTarget,
			}

			for _, arg := range tt.args {
				b.AppendTransactions(arg.sender, arg.receiver, arg.amount)
			}

			assert.Equal(t, tt.expect, len(b.LastTransaction))
		})
	}
}
