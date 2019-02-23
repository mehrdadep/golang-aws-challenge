package main

import (
	"golang-aws-challenge/functions"
	"reflect"
	"testing"
)

//TestHandler test whether returned status code is expected or not
func TestHandler(t *testing.T) {
	type args struct {
		request functions.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			args: args{
				functions.Request{
					PathParameters: map[string]string{"id": "ee230a7b-3615-11e9-88d9-2288fa4453c9"},
				},
			},
			want: 200,
		},
		{
			args: args{
				functions.Request{
					PathParameters: map[string]string{"id": "ee230a7b-3615-11e9-88d9-2288fa4453c98"},
				},
			},
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
