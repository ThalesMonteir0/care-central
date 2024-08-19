package response

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/request"
)

type PixOauthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type CreatePixEFIResponse struct {
	Calendario         request.Calendario `json:"calendario"`
	Txid               string             `json:"txid"`
	Revisao            int                `json:"revisao"`
	Loc                Loc                `json:"loc"`
	Location           string             `json:"location"`
	Status             string             `json:"status"`
	Devedor            request.Devedor    `json:"devedor"`
	Valor              request.Valor      `json:"valor"`
	Chave              string             `json:"chave"`
	SolicitacaoPagador string             `json:"solicitacaoPagador"`
	PixCopiaECola      string             `json:"pixCopiaECola"`
}

type Loc struct {
	Id       int    `json:"id"`
	Location string `json:"location"`
	TipoCob  string `json:"tipoCob"`
}

type CreatePixEFIError struct {
	Nome      string `json:"nome"`
	Menssagem string `json:"menssagem"`
}
