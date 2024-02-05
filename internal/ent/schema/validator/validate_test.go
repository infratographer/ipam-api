package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOfBlock(t *testing.T) {
	type args struct {
		ipBl   string
		ipAdrr string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{

			name: "happy path",
			args: args{
				ipBl:   "192.168.1.0/28",
				ipAdrr: "192.168.1.13"},
			wantErr: false,
		},
		{
			name: "outside block",
			args: args{
				ipBl:   "192.168.1.0/28",
				ipAdrr: "192.168.1.25"},
			wantErr: true,
		},
		{
			name: "far from block",
			args: args{
				ipBl:   "108.1.80.128/30",
				ipAdrr: "192.168.10.12"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PartOfBlock(tt.args.ipBl, tt.args.ipAdrr)
			if tt.wantErr {
				assert.Error(t, err)
				assert.ErrorContainsf(t, err, tt.args.ipAdrr, tt.args.ipBl)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
