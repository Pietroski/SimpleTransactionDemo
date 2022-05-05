package mocked_auth_middleware

type (
	CtxMockedKey    string
	MockedBearToken string
	MockedAccountID string
)

const (
	MainMockedBearToken       MockedBearToken = "c554dcf3120527507902ca423c53dd10111cfaf6647c4218d3b557471b1fe33e"
	SecondaryMockedBearToken  MockedBearToken = "c322b964506e4b858c2cd1a56855c14ee5e5374e1459fbcca5996454fd42539d"
	TertiaryMockedBearToken   MockedBearToken = "2d866b2342ab4fa5d350063b8897bc71e7c602d7ebb55bfbffb0fb92b57ada5c"
	QuaternaryMockedBearToken MockedBearToken = "3860a90970ab7ea90507c350619044ac2f41a973471dd8ef84cdec00ce7079bc"

	MainAccountID       MockedAccountID = "410ae8d3-1573-4f37-8b38-1bdece117db5"
	SecondaryAccountID  MockedAccountID = "dcf5baae-ce7e-426e-8c75-6f7212aace94"
	TertiaryAccountID   MockedAccountID = "e0c3dc07-abc9-4ec2-8994-3215b11ee935"
	QuaternaryAccountID MockedAccountID = "44adb9a0-6575-42f4-8fce-c9f90ddd71f1"

	CtxMockedAuthKey CtxMockedKey = "Mocked-Auth-Layer"
)

func (cmk CtxMockedKey) String() string {
	return string(cmk)
}

func (mbt MockedBearToken) String() string {
	return string(mbt)
}

func (maid MockedAccountID) String() string {
	return string(maid)
}

type MockedAuthValues struct {
	BearToken MockedBearToken
	AccountID MockedAccountID
}

var (
	MockedAuthMap = map[MockedBearToken]MockedAuthValues{
		MainMockedBearToken: {
			BearToken: MainMockedBearToken,
			AccountID: MainAccountID,
		},
		SecondaryMockedBearToken: {
			BearToken: SecondaryMockedBearToken,
			AccountID: SecondaryAccountID,
		},
		TertiaryMockedBearToken: {
			BearToken: TertiaryMockedBearToken,
			AccountID: TertiaryAccountID,
		},
		QuaternaryMockedBearToken: {
			BearToken: QuaternaryMockedBearToken,
			AccountID: QuaternaryAccountID,
		},
	}
)
