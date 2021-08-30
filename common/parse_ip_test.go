package common

import (
	"reflect"
	"testing"
)



var ips = []string{"192.168.1.1","192.168.1.2","192.168.1.3"}

func TestParseIP(t *testing.T) {
	type args struct {
		ip       string
		filename string
	}
	tests := []struct {
		name      string
		args      args
		wantHosts []string
		wantErr   bool
	}{

		{"ceshi", args{"192.168.1.1-192.168.1.3",""}, ips,false},
		{"ceshi", args{"","/Users/eleven/GolandProjects/GoScan/ips.txt"}, ips,false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHosts, err := ParseIP(tt.args.ip, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHosts, tt.wantHosts) {
				t.Errorf("ParseIP() gotHosts = %v, want %v", gotHosts, tt.wantHosts)
			}
		})
	}
}