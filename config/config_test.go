package config

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		configBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name:    "NoConfig",
			args:    args{},
			wantErr: true,
		},
		{
			name: "InvalidFormat",
			args: args{
				configBytes: []byte(`#!headstar`),
			},
			wantErr: true,
		},
		{
			name: "SimpleCase",
			args: args{
				configBytes: []byte(`#!headstart

users:
  deedubs:
    name: Dan Williams`),
			},
			want: &Config{
				Users: map[string]User{
					"deedubs": User{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.configBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
