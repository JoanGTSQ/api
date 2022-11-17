package api

import "errors"

var (
	ERR_NOT_MATCH_LOGIN = errors.New("ALFA-001")

	ERR_NOT_FOUND       = errors.New("ALFA-002")
	ERR_USERID_REQUIRED = errors.New("ALFA-003")
	ERR_ID_INVALID      = errors.New("ALFA-004")

	ERR_MAIL_REQUIRED        = errors.New("ALFA-005")
	ERR_MAIL_IS_N0T_VALID    = errors.New("ALFA-006")
	ERR_MAIL_IS_TAKEN        = errors.New("ALFA-007")
	ERR_MAIL_INVALID         = errors.New("ALFA-008")
	ERR_MAIL_NOT_EXIST       = errors.New("ALFA-009")
	ERR_NOT_FOUND_MODEL      = "models: resource not found"
	ERR_INVALID_LITERAL_JSON = errors.New("EOF")
	ERR_PSSWD_INCORRECT      = errors.New("ALFA-010")
	ERR_PSSWD_TOO_SHORT      = errors.New("ALFA-011")
	ERR_PSSWD_REQUIRED       = errors.New("ALFA-012")
	ERR_PSSWD_SAME_RESET     = errors.New("ALFA-003")

	ERR_REMMEMBER_TOO_SHOT = errors.New("ALFA-014")
	ERR_REMMEMBER_REQUIRED = errors.New("ALFA-015")

	ERR_PSSWD_RESET_TOKEN_EXPIRED    = errors.New("ALFA-016")
	ERR_PSSWD_RESET_TOKEN_DUPLICATED = errors.New("ALFA-017")

	ERR_JWT_TOKEN_EXPIRED      = errors.New("ALFA-018")
	ERR_JWT_CLAIMS_INVALID     = errors.New("ALFA-019")
	ERR_JWT_TOKEN_REQUIRED     = errors.New("ALFA-020")
	ERR_NOT_ENOUGH_PERMISSIONS = errors.New("ALFA-021")

	ERR_INVALID_JSON = errors.New("ALFA-022")

	ERR_NOT_SAME_USER = errors.New("ALFA-023")

  ERR_UNDEFINED_TYPE_TRANSACTION = errors.New("BRAVO-001")
)
