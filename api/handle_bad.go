package api

func HandleBad() error {
	return HttpError{
		Status:      400,
		Description: "The error was bad, very bad",
	}
}
