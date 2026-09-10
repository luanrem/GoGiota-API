package jsonutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/luanrem/GoGiota-API/internal/validator"
)

// maxBodyBytes limita o tamanho do corpo aceito numa requisicao. Sem esse
// limite, um cliente pode enviar um corpo enorme e consumir recursos do
// servidor indefinidamente.
const maxBodyBytes = 1 << 20 // 1 MB

var (
	ErrEmptyBody     = errors.New("corpo da requisicao vazio")
	ErrBodyTooLarge  = errors.New("corpo da requisicao excede o limite permitido")
	ErrMalformedJSON = errors.New("json malformado")
)

// EncodeJson serializa data e escreve a resposta com o status informado.
//
// A serializacao acontece antes de qualquer byte ir para o cliente: se ela
// falhar, o status ainda nao foi enviado e quem chamou consegue responder um
// erro no lugar.
func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(body); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

// DecodeJson le o corpo da requisicao para T, sem validar.
func DecodeJson[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var data T

	if err := decodeBody(w, r, &data); err != nil {
		return data, err
	}

	return data, nil
}

// DecodeValidJson le o corpo da requisicao para T e roda a validacao do
// proprio tipo.
//
// Quando a validacao falha, o mapa retornado descreve o problema de cada
// campo e o erro nao e nil. Quando a falha e de decodificacao, o mapa vem
// vazio — e assim que quem chama distingue 400 de 422.
func DecodeValidJson[T validator.Validator](w http.ResponseWriter, r *http.Request) (T, map[string]string, error) {
	var data T

	if err := decodeBody(w, r, &data); err != nil {
		return data, nil, err
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

func decodeBody(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return decodeError(err)
	}

	// O decoder para no primeiro valor JSON. Uma segunda leitura precisa
	// encontrar o fim do corpo, senao veio mais de um objeto.
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return fmt.Errorf("%w: o corpo deve conter um unico objeto", ErrMalformedJSON)
	}

	return nil
}

// decodeError traduz os erros do encoding/json em mensagens que ajudam quem
// esta chamando a API a corrigir o request.
func decodeError(err error) error {
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError
	var maxBytesErr *http.MaxBytesError

	switch {
	case errors.Is(err, io.EOF):
		return ErrEmptyBody

	case errors.As(err, &maxBytesErr):
		return ErrBodyTooLarge

	case errors.As(err, &syntaxErr):
		return fmt.Errorf("%w na posicao %d", ErrMalformedJSON, syntaxErr.Offset)

	case errors.As(err, &typeErr):
		return fmt.Errorf("%w: o campo %q esperava um valor do tipo %s", ErrMalformedJSON, typeErr.Field, typeErr.Type)

	case errors.Is(err, io.ErrUnexpectedEOF):
		return fmt.Errorf("%w: json incompleto", ErrMalformedJSON)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		campo := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return fmt.Errorf("%w: campo desconhecido %s", ErrMalformedJSON, campo)

	default:
		return fmt.Errorf("failed to decode json: %w", err)
	}
}
