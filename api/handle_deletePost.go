package api

import (
	"context"
	"net/http"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/utils"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	logger := utils.GlobalLogger

	postId := r.PathValue("postId")
	logger.Info("Intentando eliminar post con ID: %s", postId)

	p := GetPostPersistence(ctx)

	post, err := p.Get(ctx, postId)
	if err != nil {
		logger.Error("Error al obtener el post: %v", err)
		return ErrorPersistenceRead
	}
	if post == nil {
		logger.Warn("Post no encontrado con ID: %s", postId)
		return ErrorPostNotFound
	}

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		logger.Warn("Acceso denegado al post %s para el usuario %s", postId, auth.GetAuth(ctx).User.ID)
		return ErrorPostForbidden
	}

	err = p.Delete(ctx, postId)
	if err != nil {
		logger.Error("Error al eliminar el post: %v", err)
		return ErrorPersistenceWrite
	}

	logger.Success("Post %s eliminado exitosamente", postId)
	w.WriteHeader(http.StatusNoContent)
	return nil
}
