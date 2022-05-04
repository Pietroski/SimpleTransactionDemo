package mocked_auth_middleware

type (
	CtxMockedKey    string
	MockedBearToken string
	MockedAccountID string
)

const (
	MainMockedBearToken      MockedBearToken = "c554dcf3120527507902ca423c53dd10111cfaf6647c4218d3b557471b1fe33e"
	SecondaryMockedBearToken MockedBearToken = "d88b4b1e77c70ba780b56032db1c259b"

	MainAccountID      MockedAccountID = "410ae8d3-1573-4f37-8b38-1bdece117db5"
	SecondaryAccountID MockedAccountID = "dcf5baae-ce7e-426e-8c75-6f7212aace94"

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
	}
)
