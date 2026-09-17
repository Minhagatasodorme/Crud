package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Message struct {
	Message string `json:"error"`
	Id      int    `json:"id"`
	Nome    string `json:"nome"`
	Email   string `json:"email"`
}

var database sql.DB = *OpenDB()

func GetUser(nome, email string) []byte {

	if email == "" && nome == "" {
		err := errors.New("Parâmetros vazios!")
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	if email != "" && nome != "" {
		err := errors.New("Envie apenas Nome ou Email!")
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	if email == "" && nome != "" {
		ok, data := VerifyNome(nome)
		if ok {
			data_message := DataGetUser("", nome, 0)
			if data_message.Message != "" {
				return data
			} else {
				return TransformJson(data_message)
			}
		} else {
			return data
		}

	} else {
		ok, data := VerifyEmail(email)
		if ok {
			data_message := DataGetUser(email, "", 0)
			if data_message.Message != "" {
				return data
			} else {
				return TransformJson(data_message)
			}
		} else {
			return data
		}
	}
}

func GetUserAll() []byte {
	users := make([]Message, 0)

	rows, err := database.Query("SELECT * FROM crud")
	if err != nil {
		Message := Message{Message: err.Error()}
		return TransformJson(Message)
	}
	defer rows.Close()

	for rows.Next() {
		var user Message
		err = rows.Scan(&user.Id, &user.Nome, &user.Email)
		if err != nil {
			message := Message{Message: err.Error()}
			return TransformJson(message)
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	Json, err := json.MarshalIndent(users, "  ", "   ")
	if err != nil {
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	return Json
}

func PostUser(nome, email string) []byte {
	var check_email string
	var check_nome string

	err := database.QueryRow("SELECT email FROM crud WHERE nome = ?", nome).Scan(&check_email)
	if err == nil {
		err = errors.New("Nome já em uso!")
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	err = database.QueryRow("SELECT nome FROM crud WHERE email = ?", email).Scan(&check_nome)
	if err == nil {
		err = errors.New("Email já em uso!")
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	result, err := database.Exec("INSERT INTO crud (nome, email) VALUES (?, ?)", nome, email)
	if err != nil {
		message := Message{Message: err.Error()}
		return TransformJson(message)
	}

	id, err := result.LastInsertId()

	message := Message{
		Message: "Sucess",
		Id:      int(id),
		Nome:    nome,
		Email:   email,
	}

	return TransformJson(message)

}

func DeleteUser(Id string, nome, email string) []byte {
	id, err := strconv.Atoi(Id)
	if err != nil {
		return TransformJson(Message{Message: err.Error()})
	}

	if id != 0 {
		ok, data := VerifyId(id)
		if ok {
			data_get := DataGetUser("", "", id)
			if data_get.Message != "" {
				return TransformJson(data_get)
			}

			_, err := database.Exec("DELETE FROM crud WHERE id = ?", id)
			if err != nil {
				return TransformJson(Message{Message: err.Error()})
			}

			message := Message{
				Message: "Sucess",
				Id:      id,
				Nome:    data_get.Nome,
				Email:   data_get.Email,
			}

			return TransformJson(message)
		} else {
			return data
		}
	}

	if nome != "" {
		ok, data := VerifyNome(nome)
		if ok {
			data_get := DataGetUser("", nome, 0)
			if data_get.Message != "" {
				return TransformJson(data_get)
			}

			result, err := database.Exec("DELETE FROM crud WHERE nome = ?", nome)
			fmt.Printf("result: %v\n", result)
			if err != nil {
				return TransformJson(Message{Message: err.Error()})
			}

			message := Message{
				Message: "Sucess",
				Id:      data_get.Id,
				Nome:    nome,
				Email:   data_get.Email,
			}

			return TransformJson(message)
		} else {
			return TransformJson(data)
		}
	}

	if email != "" {
		ok, data := VerifyEmail(email)
		if ok {
			data_get := DataGetUser(email, "", 0)
			if data_get.Message != "" {
				return TransformJson(data_get)
			}

			result, err := database.Exec("DELETE FROM crud WHERE email = ?", email)
			fmt.Printf("result: %v\n", result)
			if err != nil {
				return TransformJson(Message{Message: err.Error()})
			}

			message := Message{
				Message: "Sucess",
				Id:      data_get.Id,
				Nome:    data_get.Nome,
				Email:   email,
			}

			return TransformJson(message)
		} else {
			return TransformJson(data)
		}
	}

	return []byte("")
}
