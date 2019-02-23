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
					Body:    "{\"name\": \"Test name5\",\"serial\": \"A205ad05065605000\",\"deviceModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c9\",\"note\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 201,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Test name5\",\"serial\": \"A205ad0565000\",\"deviceModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c9\",\"note\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 200,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Unit Test n\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Test name5\",\"serial\": \"A205ad0565000\",\"deviceModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c9\",\"noste\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Test name5\",\"serial\": \"A205ad0565000\",\"devicedModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c9\",\"note\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Test name5\",\"sersial\": \"A205ad0565000\",\"deviceModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c9\",\"note\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name\": \"Test name5\",\"serial\": \"A205ad0565000000\",\"deviceModel\":\"ee230a7b-3615-11e9-88d9-2288fa4453c90\",\"note\": \"Test note\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "{\"name5\": \"Unit Test n\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
				},
			},
			want: 400,
		},
		{
			args: args{
				functions.Request{
					Body:    "\"name5\": \"Unit Test n\"}",
					Headers: map[string]string{"x-api-key": "3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr"},
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
