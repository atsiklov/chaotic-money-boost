package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists") // неуникальное поле в БД
	ErrUserNotFound      = errors.New("user not found")

	ErrChgeTemplateNotFound = errors.New("challenge template not found")

	ErrChgeShowcaseNotFound = errors.New("challenge showcase not found")
)
