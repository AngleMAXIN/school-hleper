package spider

import "testing"

func TestDecrypt(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test get pass",
			args: args{
				input: "maxin123",
			},
			want: "255ae45deb4ec6e2c79a07d0946ac648",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValueFromCookie(t *testing.T) {
	type args struct {
		cookie string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases
		{
			name: "test get cookie value from source cookie",
			args: args{
				cookie: "JSESSIONID=F87EF45033C29B1326C82A03CE693769; Path=/sso; HttpOnly",
			},
			want: "F87EF45033C29B1326C82A03CE693769",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValueFromCookie(tt.args.cookie); got != tt.want {
				t.Errorf("getValueFromCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}