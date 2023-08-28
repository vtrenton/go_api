package api

//import {
//	"encoding/json"
//	"net/http"
//}

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// HTTP Return code - 200 is best!
	Code int
	// Account Balance
	Balance int64
}

// Error handling struct
type error struct {
	// Errior code
	Code int
	// Error message
	Message string
}
