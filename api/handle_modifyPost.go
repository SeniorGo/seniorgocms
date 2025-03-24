package api

import (
	"context"
	"net/http"
	"time"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/utils"
)

type ModifyPostRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func HandleModifyPost(ctx context.Context, r *http.Request, input *ModifyPostRequest) (*Post, error) {
	logger := utils.GlobalLogger

	postId := r.PathValue("postId")
	p := GetPostPersistence(ctx)

	logger.Info("Modificando post con ID: %s", postId)

	post, err := p.Get(ctx, postId)
	if err != nil {
		logger.Error("Error al obtener el post: %v", err)
		return nil, ErrorPersistenceRead
	}
	if post == nil {
		logger.Warn("Post no encontrado con ID: %s", postId)
		return nil, ErrorPostNotFound
	}

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		logger.Warn("Acceso denegado al post %s para el usuario %s", postId, auth.GetAuth(ctx).User.ID)
		return nil, ErrorPostForbidden
	}

	post.Item.Author = auth.GetAuth(ctx).User // Update user data
	post.Item.ModificationTime = time.Now()

	if input.Title != nil {
		logger.Info("Actualizando título del post %s: %s", postId, *input.Title)
		post.Item.Title = *input.Title
	}

	if input.Body != nil {
		logger.Info("Actualizando contenido del post %s", postId)
		post.Item.Body = *input.Body
	}

	err = post.Item.Validate()
	if err != nil {
		logger.Error("Error de validación: %v", err)
		return nil, err
	}

	err = p.Put(ctx, post)
	if err != nil {
		logger.Error("Error al guardar el post: %v", err)
		return nil, ErrorPersistenceWrite
	}

	logger.Success("Post %s modificado exitosamente", postId)
	return &post.Item, nil
}
