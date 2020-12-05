package infrastructure

import "github.com/google/uuid"

func GetUuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		err.Error()
	}
	uuID := u.String()
	return uuID
}
