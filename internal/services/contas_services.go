package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luanrem/GoGiota-API/internal/store/pgstore"
)

type ContaService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewContaService(pool *pgxpool.Pool) ContaService {
	return ContaService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

type CriarPessoaFisicaParams struct {
	RendaMensal  string
	Idade        int32
	NomeCompleto string
	Celular      string
	Email        string
	Categoria    string
}

func (cs *ContaService) CriarPessoaFisica(
	ctx context.Context,
	arg CriarPessoaFisicaParams,
) (int32, error) {

	pessoaFisicaArgs := pgstore.CriarPessoaFisicaParams{
		RendaMensal:  arg.RendaMensal,
		Idade:        arg.Idade,
		NomeCompleto: arg.NomeCompleto,
		Celular:      arg.Celular,
		Email:        arg.Email,
		Categoria:    arg.Categoria,
	}

	id, err := cs.queries.CriarPessoaFisica(ctx, pessoaFisicaArgs)
	if err != nil {
		return 0, fmt.Errorf("criar pessoa fisica: %w", err)
	}

	return id, nil
}

type CriarPessoaJuridicaParams struct {
	Faturamento      string
	Idade            int32
	NomeFantasia     string
	Celular          string
	EmailCorporativo string
	Categoria        string
}

func (cs *ContaService) CriarPessoaJuridica(
	ctx context.Context,
	arg CriarPessoaJuridicaParams,
) (int32, error) {

	pessoaJuridicaArgs := pgstore.CriarPessoaJuridicaParams{
		Faturamento:      arg.Faturamento,
		Idade:            arg.Idade,
		NomeFantasia:     arg.NomeFantasia,
		Celular:          arg.Celular,
		EmailCorporativo: arg.EmailCorporativo,
		Categoria:        arg.Categoria,
	}

	id, err := cs.queries.CriarPessoaJuridica(ctx, pessoaJuridicaArgs)
	if err != nil {
		return 0, fmt.Errorf("criar pessoa juridica: %w", err)
	}

	return id, nil
}
