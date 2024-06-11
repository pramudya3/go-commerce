package utils

type CustomError struct {
	Message string
}

func (c *CustomError) Error() string {
	return c.Message
}
