package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

// func stubWallet() ([]Wallet, error){
// 	var wallets []Wallet
// 	return wallets, nil
// }

type Stub struct {
	wallets          []Wallet
	findWalletByType []Wallet
	findWalletByUser []Wallet
	createWallet     Wallet
	updateWallet     Wallet
	deleteWallet     error
	error            error
}

func (s *Stub) Wallets() ([]Wallet, error) {
	return s.wallets, s.error
}

func (s *Stub) FindWalletByType(walletType string) ([]Wallet, error) {
	return s.findWalletByType, s.error
}

func (s *Stub) FindWalletByUser(user_id string) ([]Wallet, error) {
	return s.findWalletByUser, s.error
}

func (s *Stub) CreateWallet(wallet Wallet) (Wallet, error) {
	return s.createWallet, s.error
}

func (s *Stub) UpdateWallet(wallet Wallet) (Wallet, error) {
	return s.updateWallet, s.error
}

func (s *Stub) DeleteWallet(id string) error {
	return s.deleteWallet
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := &Stub{error: echo.ErrInternalServerError}
		h := New(stubError)

		h.WalletHandler(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}

		if rec.Code == http.StatusInternalServerError {

			//log the status code
			t.Logf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)

		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubWallet := &Stub{wallets: []Wallet{
			{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John Doe Wallet", WalletType: "Personal", Balance: 1000.00},
		}, error: nil}

		h := New(stubWallet)
		h.WalletHandler(c)

		want := Wallet{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John Doe Wallet", WalletType: "Personal", Balance: 1000.00}
		gotJson := rec.Body.Bytes()

		var got []Wallet

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}

		if !reflect.DeepEqual(want, got[0]) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
