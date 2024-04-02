package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	FindWalletByType(walletType string) ([]Wallet, error)
	FindWalletByUser(user_id string) ([]Wallet, error)
	CreateWallet(wallet Wallet) (Wallet, error)
	UpdateWallet(wallet Wallet) (Wallet, error)
	DeleteWallet(id string) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet_type	query	string	false	"name search by wallet_type"
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		400	{object}	Err
//
// @Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {

	wallet_type := c.QueryParam("wallet_type")
	if wallet_type == "Savings" {
		wallets, err := h.store.FindWalletByType(wallet_type)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, wallets)
	}

	wallets, err := h.store.Wallets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, wallets)
}

// WalletUserHandler
//
//	@Summary		Get wallets by user ID
//	@Description	Get wallets by user ID
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/{id}/wallets [get]
//	@Failure		400	{object}	Err
//
// @Failure		500	{object}	Err
func (h *Handler) WalletUserHandler(c echo.Context) error {
	user := c.Param("id")

	wallets, err := h.store.FindWalletByUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	if len(wallets) == 0 {
		return c.JSON(http.StatusBadRequest, Err{Message: "Wallet not found by user ID"})
	}

	return c.JSON(http.StatusOK, wallets)
}

// CreateWalletHandler
//
//	@Summary		Create a new wallet
//	@Description	Create a new wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body	Wallet	true	"Wallet object that needs to be created"
//	@Success		201	{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		400	{object}	Err
//
// @Failure		500	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	var wallet Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	_wallet, err := h.store.CreateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, _wallet)
}

// UpdateWalletHandler
//
//	@Summary		Update a wallet
//	@Description	Update a wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body	Wallet	true	"Wallet object that needs to be updated"
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [put]
//	@Failure		400	{object}	Err
//
// @Failure		500	{object}	Err
func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	var wallet Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	_wallet, err := h.store.UpdateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, _wallet)
}

// DeleteWalletHandler
//
//	@Summary		Delete a wallet
//	@Description	Delete a wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body	Wallet	true	"Wallet object that needs to be updated"
//
// @Example			ExampleBody { "id":6 }
//
//	@Success		204	{object}	Wallet
//	@Router			/api/v1/users/{id}/wallets [delete]
//
// @Failure		500	{object}	Err
func (h *Handler) DeleteWalletHandler(c echo.Context) error {

	var wallet Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	err := h.store.DeleteWallet(strconv.Itoa(wallet.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, "Wallet deleted")
}
