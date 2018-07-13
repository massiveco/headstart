package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestFetchConfig(t *testing.T) {
	type args struct {
		filename string
	}
	userData, err := ioutil.ReadFile("../../sample_config.yml")
	if err != nil {
		t.Error(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "basic", args: args{filename: "../../sample_config.yml"}, want: userData},
		{name: "basic", args: args{filename: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
