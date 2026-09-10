package conta

import (
	"context"

	"github.com/luanrem/GoGiota-API/internal/validator"
)

type CriarContaReq struct {
	Tipo string `json:"tipo"` // "fisica" | "juridica"

	// comuns
	Idade     int32  `json:"idade"`
	Celular   string `json:"celular"`
	Categoria string `json:"categoria"`

	// pessoa fisica
	NomeCompleto string `json:"nome_completo"`
	Email        string `json:"email"`
	RendaMensal  string `json:"renda_mensal"`

	// pessoa juridica
	NomeFantasia     string `json:"nome_fantasia"`
	EmailCorporativo string `json:"email_corporativo"`
	Faturamento      string `json:"faturamento"`
}

func (req CriarContaReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	// ---------- campos comuns aos dois tipos ----------
	eval.CheckField(validator.NotBlank(req.Celular), "celular", "esse campo nao pode ficar em branco")
	eval.CheckField(validator.MaxChars(req.Celular, 20), "celular", "esse campo deve ter no maximo 20 caracteres")

	eval.CheckField(validator.NotBlank(req.Categoria), "categoria", "esse campo nao pode ficar em branco")
	eval.CheckField(validator.MaxChars(req.Categoria, 50), "categoria", "esse campo deve ter no maximo 50 caracteres")

	// ---------- campos especificos de cada tipo ----------
	switch req.Tipo {
	case "fisica":
		eval.CheckField(req.Idade >= 0 && req.Idade <= 150, "idade", "esse campo deve estar entre 0 e 150")

		eval.CheckField(validator.NotBlank(req.NomeCompleto), "nome_completo", "esse campo nao pode ficar em branco")
		eval.CheckField(validator.MaxChars(req.NomeCompleto, 255), "nome_completo", "esse campo deve ter no maximo 255 caracteres")

		eval.CheckField(validator.NotBlank(req.Email), "email", "esse campo nao pode ficar em branco")
		eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "esse campo deve ser um email valido")
		eval.CheckField(validator.MaxChars(req.Email, 255), "email", "esse campo deve ter no maximo 255 caracteres")

		eval.CheckField(validator.Matches(req.RendaMensal, validator.MoneyRX), "renda_mensal", "esse campo deve ser um valor positivo com ate duas casas decimais")

	case "juridica":
		eval.CheckField(req.Idade >= 0, "idade", "esse campo deve ser um numero positivo")

		eval.CheckField(validator.NotBlank(req.NomeFantasia), "nome_fantasia", "esse campo nao pode ficar em branco")
		eval.CheckField(validator.MaxChars(req.NomeFantasia, 255), "nome_fantasia", "esse campo deve ter no maximo 255 caracteres")

		eval.CheckField(validator.NotBlank(req.EmailCorporativo), "email_corporativo", "esse campo nao pode ficar em branco")
		eval.CheckField(validator.Matches(req.EmailCorporativo, validator.EmailRX), "email_corporativo", "esse campo deve ser um email valido")
		eval.CheckField(validator.MaxChars(req.EmailCorporativo, 255), "email_corporativo", "esse campo deve ter no maximo 255 caracteres")

		eval.CheckField(validator.Matches(req.Faturamento, validator.MoneyRX), "faturamento", "esse campo deve ser um valor positivo com ate duas casas decimais")

	default:
		eval.AddFieldError("tipo", "esse campo deve ser 'fisica' ou 'juridica'")
	}

	return eval
}
