package data

type Payload struct {
	CurrencyId Currency `json:"currencyId"`
	TokenId    Token    `json:"tokenId"`
	Side       Side     `json:"side"`
	Size       string   `json:"size"`
}

type Response struct {
	RetCode int       `json:"ret_code"`
	RetMsg  string    `json:"ret_msg"`
	Result  AdsResult `json:"result"`
}

type AdsResult struct {
	Count uint16 `json:"count"`
	Items []Item `json:"items"`
}

type Item struct {
	Id         string   `json:"id"`
	AccountId  string   `json:"accountId"`
	UserId     string   `json:"userId"`
	Nickname   string   `json:"nickname"`
	TokenId    Token    `json:"tokenId"`
	CurrencyId Currency `json:"currencyId"`
	Price      string   `json:"price"`
	Quantity   string   `json:"quantity"`
	Remark     string   `json:"remark"`

	RecentExecuteRate uint16 `json:"recentExecuteRate"`
}
