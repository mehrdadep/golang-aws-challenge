package functions

import (
	"reflect"
	"testing"
)

func TestReturnResponse(t *testing.T) {
	type args struct {
		body string
		code int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			args: args{
				body: "String of test",
				code: 400,
			},
			want: 400,
		},
		{
			args: args{
				body: "String of test",
				code: 200,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReturnResponse(tt.args.body, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want) {
				t.Errorf("ReturnResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			want: "https://dynamodb.eu-west-3.amazonaws.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Endpoint, tt.want) {
				t.Errorf("ConnectDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDeviceByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Device
		wantErr bool
	}{
		{
			args: args{
				id: "afasfas-asdasd-",
			},
			want: nil,
		},
		{
			args: args{
				id: "624f653c-3787-11e9-b5f3-3a32138d7a6f",
			},
			want: &Device{
				ID:          "624f653c-3787-11e9-b5f3-3a32138d7a6f",
				Name:        "Test name6",
				Serial:      "A205ad00565000",
				Note:        "Test note",
				DeviceModel: "ee230a7b-3615-11e9-88d9-2288fa4453c9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDeviceByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDeviceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDeviceByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDeviceBySerial(t *testing.T) {
	type args struct {
		serial string
	}
	tests := []struct {
		name    string
		args    args
		want    *Device
		wantErr bool
	}{
		{
			args: args{
				serial: "afasfas-asdasd-",
			},
			want: nil,
		},
		{
			args: args{
				serial: "A205ad00565000",
			},
			want: &Device{
				ID:          "624f653c-3787-11e9-b5f3-3a32138d7a6f",
				Name:        "Test name6",
				Serial:      "A205ad00565000",
				Note:        "Test note",
				DeviceModel: "ee230a7b-3615-11e9-88d9-2288fa4453c9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDeviceBySerial(tt.args.serial)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDeviceBySerial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDeviceBySerial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindModelByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Model
		wantErr bool
	}{
		{
			args: args{
				name: "afasfas-asdasd",
			},
			want: nil,
		},
		{
			args: args{
				name: "Model Three",
			},
			want: &Model{
				ID:   "ee230a7b-3615-11e9-88d9-2288fa4453c9",
				Name: "Model Three",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindModelByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindModelByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindModelByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindModelByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Model
		wantErr bool
	}{
		{
			args: args{
				id: "afasfas-asdasd",
			},
			want: nil,
		},
		{
			args: args{
				id: "ee230a7b-3615-11e9-88d9-2288fa4453c9",
			},
			want: &Model{
				ID:   "ee230a7b-3615-11e9-88d9-2288fa4453c9",
				Name: "Model Three",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindModelByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindModelByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindModelByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
