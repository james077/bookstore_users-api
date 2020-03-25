package users

import(
	"fmt"
	"github.com/james077/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64] *User)	
)

func (user *User) Get() *errors.RestErr{
	result := usersDB[user.Id]
	if result == nil{
		return errors.NewNotFoundError(fmt.Sprintf("Usuario %d no encontrado",user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr{
	current := usersDB[user.Id]
	if current != nil{
		if current.Email != user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("Email %s ya ha sido registrado",user.Email))	
		}
		return errors.NewBadRequestError(fmt.Sprintf("Usuario %d ya existe",user.Id))

	}
	usersDB[user.Id] = user
	return nil
}