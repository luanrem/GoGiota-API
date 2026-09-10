-- name: CriarPessoaJuridica :one
INSERT INTO pessoa_juridica (
  faturamento,
  idade,
  nome_fantasia,
  celular,
  email_corporativo,
  categoria
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: ConsultarSaldoPessoaJuridica :one
SELECT saldo FROM pessoa_juridica WHERE id = $1;