package infrastructure

import "github.com/google/uuid"

func GetUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return u.String(), err
	}
	uuID := u.String()
	return uuID, err
}
