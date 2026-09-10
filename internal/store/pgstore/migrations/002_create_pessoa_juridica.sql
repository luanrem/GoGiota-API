-- Write your migrate up statements here

create table pessoa_juridica (
  id                integer       generated always as identity primary key,
  faturamento       numeric(15,2) not null,
  idade             integer       not null,
  nome_fantasia     varchar(255)  not null,
  celular           varchar(20)   not null,
  email_corporativo varchar(255)  not null,
  categoria         varchar(50)   not null,
  saldo             numeric(15,2) not null default 0,

  constraint pessoa_juridica_saldo_nao_negativo       check (saldo >= 0),
  constraint pessoa_juridica_faturamento_nao_negativo check (faturamento >= 0),
  constraint pessoa_juridica_idade_nao_negativa       check (idade >= 0)
);

---- create above / drop below ----

drop table pessoa_juridica;
