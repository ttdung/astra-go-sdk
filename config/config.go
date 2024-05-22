package config

type Config struct {
	ChainId       string   `json:"chain_id,omitempty"`
	Endpoint      string   `json:"endpoint,omitempty"`
	CoinType      uint32   `json:"coin_type,omitempty"`
	PrefixAddress string   `json:"prefix_address,omitempty"`
	TokenSymbol   []string `json:"token_symbol,omitempty"`
}
