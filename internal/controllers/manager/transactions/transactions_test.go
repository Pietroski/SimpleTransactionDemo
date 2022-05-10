package transaction_controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	mockedTransactionStore "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc/mock"
	manager_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager"
	mocked_auth_middleware "github.com/Pietroski/SimpleTransactionDemo/internal/middlewares/auth/mocked"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	AnyErr = errors.New("any-error")
)

func TestTransactionController_Transfer(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully transacts between two accounts",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     2,
						WalletID:  mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   0,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					TransferTx(gomock.Any(), sqlc_bank_account_store.TransferTxParams{
						FromAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						FromWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						ToAccountID:   mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						ToWalletID:    mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
						Amount:        10,
						Coin:          sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.TransferTxResult{
						TransactionRecord: sqlc_bank_account_store.TransactionRecord{
							RowID:         1,
							FromAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							FromWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							ToAccountID:   mocked_auth_middleware.SecondaryAccountID.ParseForce(),
							ToWalletID:    mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
							Coin:          sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:        10,
							CreatedAt:     time.Time{},
						},
						FromEntry: sqlc_bank_account_store.EntryRecord{
							RowID:     1,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    -10,
							CreatedAt: time.Time{},
						},
						ToEntry: sqlc_bank_account_store.EntryRecord{
							RowID:     2,
							AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						FromWallet: sqlc_bank_account_store.Wallet{
							RowID:     1,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Balance:   90,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						ToWallet: sqlc_bank_account_store.Wallet{
							RowID:     2,
							WalletID:  mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Balance:   10,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						TransferredAmount: 10,
						TransferredCoin:   sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong account id",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
		{
			name: "fails - payload request validation error",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails - not able to retrieve fromWallet",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails - not able to retrieve toWallet",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails - not able to retrieve fromWallet - no rows",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, sql.ErrNoRows)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusNotFound),
				)
			},
		},
		{
			name: "fails - not able to process transaction",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					ToAccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
					Amount:      10,
					Coin:        manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/transfer", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     2,
						WalletID:  mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   0,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					TransferTx(gomock.Any(), sqlc_bank_account_store.TransferTxParams{
						FromAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						FromWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						ToAccountID:   mocked_auth_middleware.SecondaryAccountID.ParseForce(),
						ToWalletID:    mocked_auth_middleware.SecondaryPietroskiCoinWalletID.ParseForce(),
						Amount:        10,
						Coin:          sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.TransferTxResult{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_Deposit(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully deposits on a user's account",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/deposit", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					DepositTx(gomock.Any(), sqlc_bank_account_store.DepositTxParams{
						ToAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						ToWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						Amount:      10,
						Coin:        sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.DepositTxResult{
						ToEntry: sqlc_bank_account_store.EntryRecord{
							RowID:     2,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						ToWallet: sqlc_bank_account_store.Wallet{
							RowID:     2,
							WalletID:  mocked_auth_middleware.MainAccountID.ParseForce(),
							AccountID: mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Balance:   110,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						TransferredAmount: 10,
						TransferredCoin:   sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong account id",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/deposit", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
		{
			name: "fails - payload request validation error",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/deposit", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to retrieve the user's wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/deposit", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to deposit into the user's wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/deposit", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					DepositTx(gomock.Any(), sqlc_bank_account_store.DepositTxParams{
						ToAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						ToWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						Amount:      10,
						Coin:        sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.DepositTxResult{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_Withdraw(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully withdraws on a user's account",
			prepareReq: func() *http.Request {
				body := &manager_models.WithdrawRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/withdraw", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					WithdrawTx(gomock.Any(), sqlc_bank_account_store.WithdrawTxParams{
						FromAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						FromWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						Amount:        10,
						Coin:          sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.WithdrawTxResult{
						FromEntry: sqlc_bank_account_store.EntryRecord{
							RowID:     2,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    -10,
							CreatedAt: time.Time{},
						},
						FromWallet: sqlc_bank_account_store.Wallet{
							RowID:     2,
							WalletID:  mocked_auth_middleware.MainAccountID.ParseForce(),
							AccountID: mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Balance:   90,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						TransferredAmount: 10,
						TransferredCoin:   sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong account id",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/withdraw", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
		{
			name: "fails - payload request validation error",
			prepareReq: func() *http.Request {
				body := &manager_models.WithdrawRequest{
					Amount: 10,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/withdraw", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to retrieve the user's wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.TransactionRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/withdraw", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to deposit into the user's wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.DepositRequest{
					Amount: 10,
					Coin:   manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost, "/v1/manager/transactions/withdraw", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)

				txStore.
					EXPECT().
					WithdrawTx(gomock.Any(), sqlc_bank_account_store.WithdrawTxParams{
						FromAccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						FromWalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						Amount:        10,
						Coin:          sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.WithdrawTxResult{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_GetWalletBalance(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully get the user balance of a specific wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.BalanceRequest{
					Coin: manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/balance", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{
						RowID:     1,
						WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						Balance:   100,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong account id",
			prepareReq: func() *http.Request {
				body := &manager_models.BalanceRequest{
					Coin: manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/balance", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
		{
			name: "fails - payload validation error",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/balance", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "successfully get the user balance of a specific wallet",
			prepareReq: func() *http.Request {
				body := &manager_models.BalanceRequest{
					Coin: manager_models.CryptoCurrenciesPIETROSKICOIN,
				}
				reqBody, err := json.Marshal(body)
				require.NoError(t, err)

				payload := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/balance", payload)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetTxWallet(gomock.Any(), sqlc_bank_account_store.GetTxWalletParams{
						AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
						Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
					}).
					Times(1).
					Return(sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_GetWallets(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully get the user's wallets",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/wallets", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetAccountWallets(gomock.Any(), mocked_auth_middleware.MainAccountID.ParseForce()).
					Times(1).
					Return([]sqlc_bank_account_store.Wallet{
						{
							RowID:     1,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						{
							RowID:     4,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesBITCOIN,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						{
							RowID:     7,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesETHEREUM,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						{
							RowID:     8,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesDODGECOIN,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong account id",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/wallets", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
		{
			name: "fails - pagination error",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/wallets?page_size=test", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to get the user's wallets",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/wallets", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetAccountWallets(gomock.Any(), mocked_auth_middleware.MainAccountID.ParseForce()).
					Times(1).
					Return([]sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "successfully get the paginated user's wallets",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/wallets"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetPaginatedWalletsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.GetPaginatedWalletsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.Wallet{
						{
							RowID:     7,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesETHEREUM,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						{
							RowID:     8,
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesDODGECOIN,
							Balance:   100,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails to get the paginated user's wallets",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/wallets"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetPaginatedWalletsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.GetPaginatedWalletsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.Wallet{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to get the paginated user's wallets - sql.NoRowErr",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/wallets"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					GetPaginatedWalletsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.GetPaginatedWalletsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.Wallet{}, sql.ErrNoRows)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusNotFound),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_GetHistory(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully gets the user's entries",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/history", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListEntryLogsByAccountID(
						gomock.Any(), mocked_auth_middleware.MainAccountID.ParseForce(),
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     1,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     2,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     7,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesBITCOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     12,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesETHEREUM,
							Amount:    -10,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails to get the user's entries",
			prepareReq: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/manager/transactions/history", nil)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListEntryLogsByAccountID(
						gomock.Any(), mocked_auth_middleware.MainAccountID.ParseForce(),
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "successfully gets the paginated user's entries",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     7,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesBITCOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     12,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesETHEREUM,
							Amount:    -10,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails to get the paginated user's entries",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to get the paginated user's entries - sql.ErrNoRows",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, sql.ErrNoRows)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusNotFound),
				)
			},
		},
		{
			name: "fails - incorrect pagination",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", "test", 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails - invalid bear token",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.FailureMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusBadRequest),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}

func TestTransactionController_GetCoinHistory(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *http.Request
		stubs      func(txStore *mockedTransactionStore.MockStore)
		assertResp func(resp *http.Response)
	}{
		{
			name: "successfully gets the user's coin entries",
			prepareReq: func() *http.Request {
				query := fmt.Sprintf("?coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+query,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     1,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     2,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     5,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    5,
							CreatedAt: time.Time{},
						},
						{
							RowID:     11,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    105,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails - wrong query",
			prepareReq: func() *http.Request {
				uri := fmt.Sprintf("?coin=%v", "invalid-uri")
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+uri,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "empty query only",
			prepareReq: func() *http.Request {
				uri := fmt.Sprintf("?coin=%v", "")
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+uri,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListEntryLogsByAccountID(
						gomock.Any(), mocked_auth_middleware.MainAccountID.ParseForce(),
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     1,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     2,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     7,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesBITCOIN,
							Amount:    10,
							CreatedAt: time.Time{},
						},
						{
							RowID:     12,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesETHEREUM,
							Amount:    -10,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "successfully gets paginated user's coin entries",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("&page_size=%v&page_id=%v", 2, 2)
				query := fmt.Sprintf("?coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+query+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     5,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    5,
							CreatedAt: time.Time{},
						},
						{
							RowID:     11,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    105,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "successfully gets paginated user's coin entries - swapped order",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("?page_size=%v&page_id=%v", 2, 2)
				query := fmt.Sprintf("&coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+pagination+query,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{
						{
							RowID:     5,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    5,
							CreatedAt: time.Time{},
						},
						{
							RowID:     11,
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							WalletID:  mocked_auth_middleware.MainPietroskiCoinWalletID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Amount:    105,
							CreatedAt: time.Time{},
						},
					}, nil)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusOK),
				)
			},
		},
		{
			name: "fails to get paginated user's coin entries",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("&page_size=%v&page_id=%v", 2, 2)
				query := fmt.Sprintf("?coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+query+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
		{
			name: "fails to get paginated user's coin entries - sql.ErrNoRows",
			prepareReq: func() *http.Request {
				pagination := fmt.Sprintf("&page_size=%v&page_id=%v", 2, 2)
				query := fmt.Sprintf("?coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+query+pagination,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListPaginatedCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListPaginatedCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
							Limit:     2,
							Offset:    2,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, sql.ErrNoRows)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusNotFound),
				)
			},
		},
		{
			name: "fails to get the user's coin entries",
			prepareReq: func() *http.Request {
				query := fmt.Sprintf("?coin=%v", manager_models.CryptoCurrenciesPIETROSKICOIN.String())
				req, err := http.NewRequest(
					http.MethodGet,
					"/v1/manager/transactions/history"+query,
					nil,
				)
				require.NoError(t, err)
				req.Header.Set(
					mocked_auth_middleware.AuthorizationKey,
					"Bearer "+mocked_auth_middleware.MainMockedBearerToken.String(),
				)

				return req
			},
			stubs: func(txStore *mockedTransactionStore.MockStore) {
				txStore.
					EXPECT().
					ListCoinEntryLogsByAccountID(
						gomock.Any(),
						sqlc_bank_account_store.ListCoinEntryLogsByAccountIDParams{
							AccountID: mocked_auth_middleware.MainAccountID.ParseForce(),
							Coin:      sqlc_bank_account_store.CryptoCurrenciesPIETROSKICOIN,
						},
					).
					Times(1).
					Return([]sqlc_bank_account_store.EntryRecord{}, AnyErr)
			},
			assertResp: func(resp *http.Response) {
				require.Contains(
					t,
					resp.Status,
					fmt.Sprintf("%d", http.StatusInternalServerError),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			txStore := mockedTransactionStore.NewMockStore(ctrl)
			server := manager_factory.NewManagerServer(manager_models.Stores{
				DeviceStore: nil,
				TxStore:     txStore,
			})
			recorder := httptest.NewRecorder()

			tt.stubs(txStore)
			req := tt.prepareReq()
			server.Router.ServeHTTP(recorder, req)
			t.Log(recorder)
			tt.assertResp(recorder.Result())
		})
	}
}
