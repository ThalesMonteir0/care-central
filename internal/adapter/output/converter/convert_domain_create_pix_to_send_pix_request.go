package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"strconv"
)

func ConvertDomainCreatePixToSendPixEFI(domain domain.CreatePixDomain) request.CreatePixSendEFIRequest {
	return request.CreatePixSendEFIRequest{
		Calendario: request.Calendario{
			Expiracao: 3600,
		},
		Devedor: request.Devedor{
			Cpf:  domain.DebtorsCPF,
			Nome: domain.DebtorsName,
		},
		Valor: request.Valor{
			Original: strconv.FormatFloat(domain.Value, 'f', 2, 64),
		},
		Chave:              "",
		SolicitacaoPagador: "Pagamento referente a sessão de terapia",
	}
}
