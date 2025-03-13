package api

import (
	"net/http"
)

var ErrorPersistenceWrite = HttpError{
	Status:      http.StatusInternalServerError,
	Description: "Problem writing to persistence layer",
}

var ErrorPersistenceRead = HttpError{
	Status:      http.StatusInternalServerError,
	Description: "Problem reading from persistence layer",
}

var ErrorPostNotFound = HttpError{
	Status:      http.StatusNotFound,
	Description: "Post not found",
}

var ErrorPostForbidden = HttpError{
	Status:      http.StatusForbidden,
	Description: "You are not authorized to perform this action for this post",
}
