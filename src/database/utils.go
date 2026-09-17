package database

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
)

func jsonErr(struct_err any) []byte {
	var indent_err_json bytes.Buffer

	err_json, _ := json.Marshal(struct_err)

	json.Indent(&indent_err_json, err_json, "", "  ")

	return indent_err_json.Bytes()
}

func TransformJson(structt any) []byte {
	var indent_resp_json bytes.Buffer

	resp_json, err := json.Marshal(structt)
	if err != nil {
		message := Message{Message: err.Error()}
		message_json := jsonErr(message)
		return message_json
	}

	err = json.Indent(&indent_resp_json, resp_json, "", "  ")
	if err != nil {
		message := Message{Message: err.Error()}
		message_json := jsonErr(message)
		return message_json
	}

	return indent_resp_json.Bytes()
}

func VerifyNome(nome string) (bool, []byte) {
	var check string
	database := OpenDB()

	err := database.QueryRow("SELECT email FROM crud WHERE nome = ?", nome).Scan(&check)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("Usuário não encontrado!")
			return false, TransformJson(Message{Message: err.Error()})
		} else {
			return false, TransformJson(Message{Message: err.Error()})
		}
	}

	return true, nil
}

func VerifyEmail(email string) (bool, []byte) {
	var check string

	err := database.QueryRow("SELECT nome FROM crud WHERE email = ?", email).Scan(&check)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("Usuário não encontrado!")
			return false, TransformJson(Message{Message: err.Error()})
		} else {
			return false, TransformJson(Message{Message: err.Error()})
		}
	}

	return true, nil
}

func VerifyId(id int) (bool, []byte) {
	var check string

	err := database.QueryRow("SELECT nome FROM crud WHERE id = ?", id).Scan(&check)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err := errors.New("Usuário não encontrado!")
			message := Message{Message: err.Error()}
			return false, TransformJson(message)
		} else {
			message := Message{Message: err.Error()}
			return false, TransformJson(message)
		}
	}

	return true, nil
}

func DataGetUser(email, nome string, id int) Message {
	if email != "" {
		err := database.QueryRow("SELECT nome, id FROM crud WHERE email = ?", email).Scan(&nome, &id)
		if err != nil {
			message := Message{Message: err.Error()}
			return message
		}

		message := Message{
			Message: "",
			Id:      id,
			Nome:    nome,
			Email:   email,
		}

		return message
	}

	if nome != "" {
		err := database.QueryRow("SELECT email, id FROM crud WHERE nome = ?", nome).Scan(&email, &id)
		if err != nil {
			message := Message{Message: err.Error()}
			return message
		}

		message := Message{
			Message: "",
			Id:      id,
			Nome:    nome,
			Email:   email,
		}

		return message
	}

	if id != 0 {
		err := database.QueryRow("SELECT nome, email FROM crud WHERE id = ?", id).Scan(&nome, &email)
		if err != nil {
			message := Message{Message: err.Error()}
			return message
		}

		message := Message{
			Message: "",
			Id:      id,
			Nome:    nome,
			Email:   email,
		}

		return message
	}

	return Message{}
}
