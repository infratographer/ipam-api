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
		errMss  string
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
			errMss:  "error provided IP Address is not part of the IP Block - Prefix: 192.168.1.0/28; IP Address: 192.168.1.25",
		},
		{
			name: "far from block",
			args: args{
				ipBl:   "108.1.80.128/30",
				ipAdrr: "192.168.10.12"},
			wantErr: true,
			errMss:  "error provided IP Address is not part of the IP Block - Prefix: 108.1.80.128/30; IP Address: 192.168.10.12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PartOfBlock(tt.args.ipBl, tt.args.ipAdrr)
			if tt.wantErr {
				assert.Error(t, err)
				t.Logf("error: %+v", err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
