package mocked_auth_middleware

import (
	"github.com/google/uuid"
)

type (
	CtxMockedKey      string
	MockedBearerToken string
	MockedAccountID   string
	MockedWalletID    string
)

const (
	MainMockedBearerToken       MockedBearerToken = "c554dcf3120527507902ca423c53dd10111cfaf6647c4218d3b557471b1fe33e"
	SecondaryMockedBearerToken  MockedBearerToken = "c322b964506e4b858c2cd1a56855c14ee5e5374e1459fbcca5996454fd42539d"
	TertiaryMockedBearerToken   MockedBearerToken = "2d866b2342ab4fa5d350063b8897bc71e7c602d7ebb55bfbffb0fb92b57ada5c"
	QuaternaryMockedBearerToken MockedBearerToken = "3860a90970ab7ea90507c350619044ac2f41a973471dd8ef84cdec00ce7079bc"

	MainAccountID       MockedAccountID = "410ae8d3-1573-4f37-8b38-1bdece117db5"
	SecondaryAccountID  MockedAccountID = "dcf5baae-ce7e-426e-8c75-6f7212aace94"
	TertiaryAccountID   MockedAccountID = "e0c3dc07-abc9-4ec2-8994-3215b11ee935"
	QuaternaryAccountID MockedAccountID = "44adb9a0-6575-42f4-8fce-c9f90ddd71f1"

	MainBitcoinWalletID       MockedWalletID = "81f84772-cf13-4a98-9ff7-48eabadfb5ef"
	SecondaryBitcoinWallerID  MockedWalletID = "f9561336-0a1d-44d6-9c85-cbd4b19204d9"
	TertiaryBitcoinWalletID   MockedWalletID = "798af056-9f34-4a6e-aacc-d721560838f4"
	QuaternaryBitcoinWalletID MockedWalletID = "573ffe3d-9f27-4f0a-a0df-d275c861181f"

	MainDodgeCoinWalletID       MockedWalletID = "30313dad-2ba7-4572-ad6a-56379109d928"
	SecondaryDodgeCoinWalletID  MockedWalletID = "6406ba42-0442-4945-ac06-21f282f50c03"
	TertiaryDodgeCoinWalletID   MockedWalletID = "65322b22-e040-4362-b5b3-1e27a9d8d241"
	QuaternaryDodgeCoinWalletID MockedWalletID = "e189bf34-f26f-4ce8-9ede-c25d05c96328"

	MainEthereumWalletID       MockedWalletID = "0735ffae-f124-4d79-a614-b28452e8f14c"
	SecondaryEthereumWalletID  MockedWalletID = "d70c69b3-f901-4d35-9b67-71b9a83fe253"
	TertiaryEthereumWalletID   MockedWalletID = "bb73bc98-e0db-474f-8644-ed0e3789e521"
	QuaternaryEthereumWalletID MockedWalletID = "562fdc01-fe18-41db-8d18-04290d7048c3"

	MainPietroskiCoinWalletID       MockedWalletID = "73342efd-23f2-4c89-9d3c-7668e2a28925"
	SecondaryPietroskiCoinWalletID  MockedWalletID = "1ebc3c98-cd2b-4e15-932d-dd2b9a599bed"
	TertiaryPietroskiCoinWalletID   MockedWalletID = "7b56f874-3f50-414f-8660-72b949e00f98"
	QuaternaryPietroskiCoinWalletID MockedWalletID = "f148e59d-fcce-413c-a1cb-a00a968737d8"

	FailureMockedBearerToken MockedBearerToken = "force-to-fail-uuid"
	FailureMockedAccountID   MockedAccountID   = "will-fail-uuid-parsing"

	CtxMockedAuthKey CtxMockedKey = "Mocked-Auth-Layer"
)

func (cmk CtxMockedKey) String() string {
	return string(cmk)
}

func (mbt MockedBearerToken) String() string {
	return string(mbt)
}

func (maid MockedAccountID) String() string {
	return string(maid)
}

func (maid MockedAccountID) ParseForce() uuid.UUID {
	parsedUUID, err := uuid.Parse(maid.String())
	if err != nil {
		return uuid.UUID{}
	}
	return parsedUUID
}

func (mwid MockedWalletID) String() string {
	return string(mwid)
}

func (mwid MockedWalletID) ParseForce() uuid.UUID {
	parsedUUID, err := uuid.Parse(mwid.String())
	if err != nil {
		return uuid.UUID{}
	}
	return parsedUUID
}

type MockedAuthValues struct {
	BearerToken MockedBearerToken
	AccountID   MockedAccountID
}

var (
	MockedAuthMap = map[MockedBearerToken]MockedAuthValues{
		MainMockedBearerToken: {
			BearerToken: MainMockedBearerToken,
			AccountID:   MainAccountID,
		},
		SecondaryMockedBearerToken: {
			BearerToken: SecondaryMockedBearerToken,
			AccountID:   SecondaryAccountID,
		},
		TertiaryMockedBearerToken: {
			BearerToken: TertiaryMockedBearerToken,
			AccountID:   TertiaryAccountID,
		},
		QuaternaryMockedBearerToken: {
			BearerToken: QuaternaryMockedBearerToken,
			AccountID:   QuaternaryAccountID,
		},

		FailureMockedBearerToken: {
			BearerToken: FailureMockedBearerToken,
			AccountID:   FailureMockedAccountID,
		},
	}
)
