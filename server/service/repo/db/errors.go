package db

type RollbackError struct {
	err error
}

func NewRollbackError(err error) error {
	return &RollbackError{
		err: err,
	}
}

func (e *RollbackError) Error() string {
	return "rollback: " + e.err.Error()
}
