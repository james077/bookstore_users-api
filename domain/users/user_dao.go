package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/james077/bookstore_users-api/datasources/mysql/users_db"
	"github.com/james077/bookstore_users-api/utils/mysql_utils"
	"github.com/james077/bookstore_utils-go/logger"
	"github.com/james077/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error mientras se preparaba la sentencia de obtencion de usuarios", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error al intentar obtener el usuario por id", getErr)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	return nil
}

func (user *User) Save() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error preparando la consulta para guardar usuario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("Error tratando de guardar usuario", saveErr)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error obteniendo el último id de inserción después de crear un nuevo usuario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	user.Id = userId

	return nil
}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error preparando la consulta de actualización de usuario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("Error actualizando usuario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error preparando la consulta de eliminación de usuario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("Error eliminando usario", err)
		return rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error preparando la consulta para encontrar usuarios por  estado", err)
		return nil, rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error buscando usuario por estado", err)
		return nil, rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error al escanear el registro del usuario en la estructura del usuario", err)
			return nil, rest_errors.NewInternalServerError("Error al intentar obtener usuario", errors.New("database error"))
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("No hay usuarios que coincidan con el estado %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error preparando sentencia de consulta de usuario por email y password", err)
		return rest_errors.NewInternalServerError("Error al intentar encontrar usuario", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("Credenciales de usuario inválidas")
		}
		logger.Error("Error al intentar obtener el usuario por correo electrónico y contraseña", getErr)
		return rest_errors.NewInternalServerError("Error al intentar encontrar usuario", errors.New("database error"))
	}
	return nil
}
