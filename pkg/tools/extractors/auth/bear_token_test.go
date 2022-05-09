package pkg_auth_extractor

import "testing"

func TestExtractBearerToken(t *testing.T) {
	type args struct {
		rawBearerToken string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "tests ExtractBearerToken in a happy path",
			args: args{
				rawBearerToken: "Bearer kjsadhfjsdbciebvxnckjdnfvieodxvl",
			},
			want:    "kjsadhfjsdbciebvxnckjdnfvieodxvl",
			wantErr: false,
		},
		{
			name: "tests ExtractBearerToken when returns an error",
			args: args{
				rawBearerToken: "kjsadhfjsdbciebvxnckjdnfvieodxvl",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractBearerToken(tt.args.rawBearerToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractBearerToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
