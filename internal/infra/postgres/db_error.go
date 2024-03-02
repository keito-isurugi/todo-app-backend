// DBからのエラーを変換する

package postgres

type DBError interface {
	Error() string
}

// NotFoundError レコードが見つからない場合のエラー
type NotFoundError struct {
	Message string
}

// BadRequestError リクエストが不正な場合のエラー
type BadRequestError struct {
	Message string
}

// DuplicateError レコードが重複している場合のエラー
type DuplicateError struct {
	Message string
}

type CheckConstraint struct {
	Message string
}

// DuplicateUniqueKeyError 複合ユニークキー制約のエラー
type DuplicateUniqueKeyError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *DuplicateError) Error() string {
	return e.Message
}

func (e *DuplicateUniqueKeyError) Error() string {
	return e.Message
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func (e *CheckConstraint) Error() string {
	return e.Message
}
