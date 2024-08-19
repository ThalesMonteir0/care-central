package request

type CreatePixSendEFIRequest struct {
	Calendario         Calendario `json:"calendario"`
	Devedor            Devedor    `json:"devedor"`
	Valor              Valor      `json:"valor"`
	Chave              string     `json:"chave"`
	SolicitacaoPagador string     `json:"solicitacaoPagador"`
}

type Calendario struct {
	Criacao   string `json:"criacao"`
	Expiracao int    `json:"expiracao"`
}

type Devedor struct {
	Cpf  string `json:"cpf"`
	Nome string `json:"nome"`
}

type Valor struct {
	Original string `json:"original"`
}
