package data

type Token string

const (
	USDT Token = "USDT"
	USDC Token = "USDC"
	BTC  Token = "BTC"
	ETH  Token = "ETH"
)

var validTokens = map[Token]bool{
	USDT: true,
	USDC: true,
	BTC:  true,
	ETH:  true,
}

func IsValidToken(token Token) bool {
	return validTokens[token]
}
