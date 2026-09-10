-- Write your migrate up statements here

create sequence conta_id_seq as integer;

create table pessoa_fisica (
  id            integer       primary key default nextval('conta_id_seq'),
  renda_mensal  numeric(15,2) not null,
  idade         integer       not null,
  nome_completo varchar(255)  not null,
  celular       varchar(20)   not null,
  email         varchar(255)  not null,
  categoria     varchar(50)   not null,
  saldo         numeric(15,2) not null default 0,

  constraint pessoa_fisica_saldo_nao_negativo check (saldo >= 0),
  constraint pessoa_fisica_renda_nao_negativa check (renda_mensal >= 0),
  constraint pessoa_fisica_com_idade_dentro_do_limite check (idade >= 0 and idade <= 150)
);

---- create above / drop below ----

drop table pessoa_fisica;
drop sequence conta_id_seq;
