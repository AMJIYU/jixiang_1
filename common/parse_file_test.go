package common

import (
	"reflect"
	"testing"
)

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
