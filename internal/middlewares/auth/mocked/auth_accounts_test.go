package mocked_auth_middleware

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCtxMockedKey_String(t *testing.T) {
	tests := []struct {
		name string
		cmk  CtxMockedKey
		want string
	}{
		{
			name: "successfully casts into string",
			cmk:  CtxMockedAuthKey,
			want: CtxMockedAuthKey.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cmk.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockedAccountID_ParseForce(t *testing.T) {
	tests := []struct {
		name string
		maid MockedAccountID
		want uuid.UUID
	}{
		{
			name: "successfully parse-forces",
			maid: MainAccountID,
			want: MainAccountID.ParseForce(),
		},
		{
			name: "unsuccessfully parse-force - default zero UUID",
			maid: MockedAccountID("invalid-uuid"),
			want: uuid.UUID{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maid.ParseForce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockedAccountID_String(t *testing.T) {
	tests := []struct {
		name string
		maid MockedAccountID
		want string
	}{
		{
			name: "successfully casts into string",
			maid: MainAccountID,
			want: MainAccountID.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maid.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockedBearerToken_String(t *testing.T) {
	tests := []struct {
		name string
		mbt  MockedBearerToken
		want string
	}{
		{
			name: "successfully casts into string",
			mbt:  MainMockedBearerToken,
			want: MainMockedBearerToken.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mbt.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockedWalletID_String(t *testing.T) {
	tests := []struct {
		name string
		mwid MockedWalletID
		want string
	}{
		{
			name: "happy path",
			mwid: MainBitcoinWalletID,
			want: MainBitcoinWalletID.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mwid.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockedWalletID_ParseForce(t *testing.T) {
	tests := []struct {
		name string
		mwid MockedWalletID
		want func() uuid.UUID
	}{
		{
			name: "successfully parses stringified uuid",
			mwid: MainPietroskiCoinWalletID,
			want: func() uuid.UUID {
				parsedUUID, err := uuid.Parse(MainPietroskiCoinWalletID.String())
				require.NoError(t, err)

				return parsedUUID
			},
		},
		{
			name: "fails to parse stringified uuid",
			mwid: "MainPietroskiCoinWalletID",
			want: func() uuid.UUID {
				return uuid.UUID{}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mwid.ParseForce(); !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("ParseForce() = %v, want %v", got, tt.want())
			}
		})
	}
}
