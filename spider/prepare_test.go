package spider

import "testing"

func Test_getCookiesFromLoginView(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test for get cookies from login view",
			want: "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCookiesFromLoginView()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCookiesFromLoginView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCookiesFromLoginView() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decryptPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test prepare get password",
			args: args{
				password: "maxin123",
			},
			want: "255ae45deb4ec6e2c79a07d0946ac648",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decryptPassword(tt.args.password); got != tt.want {
				t.Errorf("DecryptPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}