package common

import (

	"reflect"
	"testing"
)
var ipfile = []string{"192.168.1.1","192.168.1.2","192.168.1.3"}
var filename11  = "ips.txt"

func TestReadfile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"fileread",args{filename11},ipfile,false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Readfile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Readfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Readfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
