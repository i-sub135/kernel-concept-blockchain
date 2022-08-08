package node

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChain_PoW(t *testing.T) {
	type fields struct {
		Chain           []interface{}
		LastTransaction transaction
		Nodes           []string
		HashTarget      string
	}
	type args struct {
		idx     int
		hashing string
		tran    transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Test Prof of Work",
			fields: fields{
				Chain:           nil,
				LastTransaction: transaction{},
				Nodes:           nil,
				HashTarget:      "0000",
			},
			args: args{
				idx:     1,
				hashing: "abz",
				tran: transaction{
					Sender:   "a-1",
					Receiver: "b-2",
					Amount:   0.01,
				},
			},
			want: 156674,
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
			content := fmt.Sprintf("%v-%v-%v-%v", tt.args.idx, tt.args.hashing, tt.args.tran, tt.want)
			strHash := createHash(content)
			fmt.Println("string hash => ", strHash)

			assert.Equalf(t, tt.want, b.PoW(tt.args.idx, tt.args.hashing, tt.args.tran), "PoW(%v, %v, %v)", tt.args.idx, tt.args.hashing, tt.args.tran)
		})
	}
}
