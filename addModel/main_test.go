package main

import (
	"golang-aws-challenge/functions"
	"reflect"
	"testing"
)

//TestHandler tests addModel function and Status code responses
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
					Body: "{\"name\": \"Unit Test new2\"}",
				},
			},
			want: 201,
		},
		{
			args: args{
				functions.Request{
					Body: "{\"name\": \"Unit Test n\"}",
				},
			},
			want: 200,
		},
		{
			args: args{
				functions.Request{
					Body: "{\"name5\": \"Unit Test n\"}",
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body: "\"name5\": \"Unit Test n\"}",
				},
			},
			want: 500,
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
