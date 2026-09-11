package api

import (
	"net/http"

	"github.com/luanrem/GoGiota-API/internal/jsonutils"
	"github.com/luanrem/GoGiota-API/internal/services"
	"github.com/luanrem/GoGiota-API/internal/usecase/conta"
)

func (api *Api) handleCriarConta(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[conta.CriarContaReq](w, r)
	if err != nil {
		if len(problems) > 0 {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
			return
		}
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	var (
		id           int32
		serviceError error
	)

	switch data.Tipo {
	case "fisica":
		id, serviceError = api.ContaService.CriarPessoaFisica(
			r.Context(),
			services.CriarPessoaFisicaParams{
				RendaMensal:  data.RendaMensal,
				Idade:        data.Idade,
				NomeCompleto: data.NomeCompleto,
				Celular:      data.Celular,
				Email:        data.Email,
				Categoria:    data.Categoria,
			},
		)
	case "juridica":
		id, serviceError = api.ContaService.CriarPessoaJuridica(
			r.Context(),
			services.CriarPessoaJuridicaParams{
				Faturamento:      data.Faturamento,
				Idade:            data.Idade,
				NomeFantasia:     data.NomeFantasia,
				Celular:          data.Celular,
				EmailCorporativo: data.EmailCorporativo,
				Categoria:        data.Categoria,
			},
		)
	}

	if serviceError != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "falha ao criar conta, tente novamente",
		})
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"message": "Conta criada com sucesso",
		"id":      id,
	})
}
