package postgres

import (
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) FindWalletByType(walletType string) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var wallets []wallet.Wallet

	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) FindWalletByUser(user string) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", user)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var wallets []wallet.Wallet

	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) CreateWallet(w wallet.Wallet) (wallet.Wallet, error) {

	var createWallet wallet.Wallet

	err := p.Db.QueryRow("INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance) VALUES ($1, $2, $3, $4, $5) RETURNING *",
		w.UserID,
		w.UserName,
		w.WalletName,
		w.WalletType,
		w.Balance).Scan(&createWallet.ID, &createWallet.UserID, &createWallet.UserName, &createWallet.WalletName, &createWallet.WalletType, &createWallet.Balance, &createWallet.CreatedAt)

	if err != nil {
		return wallet.Wallet{}, err
	}

	return createWallet, nil
}

func (p *Postgres) UpdateWallet(w wallet.Wallet) (wallet.Wallet, error) {

	var updateWallet wallet.Wallet

	err := p.Db.QueryRow("UPDATE user_wallet SET user_id = $1, user_name = $2, wallet_name = $3, wallet_type = $4, balance = $5 WHERE id = $6 RETURNING *",
		w.UserID,
		w.UserName,
		w.WalletName,
		w.WalletType,
		w.Balance,
		w.ID).Scan(&updateWallet.ID, &updateWallet.UserID, &updateWallet.UserName, &updateWallet.WalletName, &updateWallet.WalletType, &updateWallet.Balance, &updateWallet.CreatedAt)

	if err != nil {
		return wallet.Wallet{}, err
	}

	return updateWallet, nil
}

func (p *Postgres) DeleteWallet(id string) error {
	_, err := p.Db.Exec("DELETE FROM user_wallet WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
