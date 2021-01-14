package user

import (
	"api-trainning-center/models/admin"
	"log"
)

func (st Store) ShowAllAccount() ([]admin.User, error) {
	users, err := admin.RetrieveAccounts(st.db)
	if err != nil {
		log.Fatalln("[ShowAllAccount]  error ", err)
		return []admin.User{}, err
	}

	return users, nil
}
