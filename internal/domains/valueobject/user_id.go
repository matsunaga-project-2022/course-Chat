package valueobject

type UserID string

func NewUserID(id string) (UserID, error) {
	if len(id) > 1 {
		return "", ErrInvalidInput
	}

	return UserID(id), nil
}
