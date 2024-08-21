package request

type CreatePixRequest struct {
	CpfDebtor    string  `json:"cpf_debtor" biding:"required"`
	NomeDebtor   string  `json:"nome_debtor" biding:"required"`
	ValueSession float64 `json:"value_session" biding:"required"`
	SessionID    int     `json:"session_id" biding:"required"`
}
