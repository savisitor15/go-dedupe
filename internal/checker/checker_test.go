package checker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileToSHA256(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "blank.png", args: args{p:"../../sample/blank.png"}, want: "c76cc1140b274e5eb672f0961931ed0779d61c5b371ea9be345da8b9bcd7ed8e", wantErr: false},
		{name: "silly.png", args: args{p:"../../sample/silly.png"}, want: "bf52549cad1260f8ebbb039dc8c5d79dfd9482a1396b980fdd578d38b37c592c", wantErr: false},
		{name: "fail", args: args{p:"cloudsinsky"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileToSHA256(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileToSHA256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr{
				return
			}
			if !reflect.DeepEqual(fmt.Sprintf("%x",got), tt.want) {
				t.Errorf("FileToSHA256() = %v, want = %v", got, tt.want)
			}
		})
	}
}
