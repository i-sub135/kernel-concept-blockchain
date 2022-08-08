package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChain_ValPoW(t *testing.T) {
	type fields struct {
		Chain           []interface{}
		LastTransaction transaction
		Nodes           []string
		HashTarget      string
	}
	type args struct {
		idx     int
		hashing string
		trans   transaction
		nonce   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test Validations POW",
			fields: fields{
				Chain:           nil,
				LastTransaction: transaction{},
				Nodes:           nil,
				HashTarget:      "3c3a",
			},
			args: args{
				idx:     1,
				hashing: "xyz",
				trans: transaction{
					Sender:   "a-1",
					Receiver: "b-2",
					Amount:   0.2,
				},
				nonce: 1,
			},
			want: true,
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
			assert.Equalf(t, tt.want, b.ValPoW(tt.args.idx, tt.args.hashing, tt.args.trans, tt.args.nonce), "ValPoW(%v, %v, %v, %v)", tt.args.idx, tt.args.hashing, tt.args.trans, tt.args.nonce)
		})
	}
}
