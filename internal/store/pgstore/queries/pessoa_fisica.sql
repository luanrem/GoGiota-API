-- name: CriarPessoaFisica :one
INSERT INTO pessoa_fisica (
  renda_mensal,
  idade,
  nome_completo,
  celular,
  email,
  categoria
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id;

-- name: ConsultarSaldoPessoaFisica :one
SELECT saldo FROM pessoa_fisica WHERE id = $1;