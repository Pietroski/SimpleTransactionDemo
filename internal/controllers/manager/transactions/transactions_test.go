package transaction_controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	mockedTransactionStore "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc/mock"
	manager_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager"
	mocked_auth_middleware "github.com/Pietroski/SimpleTransactionDemo/internal/middlewares/auth/mocked"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
