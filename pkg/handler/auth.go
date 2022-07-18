package handler

import "github.com/labstack/echo/v4"

//ОБЩИЕ ОБРАБОТЧИКИ

func (h *Handler) signUp(c echo.Context) error { //POST
	return nil
}

func (h *Handler) signIn(c echo.Context) error { //POST
	return nil
}

func (h *Handler) deleteProfile(c echo.Context) error { //DELETE
	return nil
}

func (h *Handler) readPost(c echo.Context) error { //GET
	return nil
}

func (h *Handler) checkProfile(c echo.Context) error { //GET
	return nil
}

func (h *Handler) createComment(c echo.Context) error { //POST
	return nil
}

func (h *Handler) editComment(c echo.Context) error { //PUT
	return nil
}

func (h *Handler) deleteComment(c echo.Context) error { //DELETE
	return nil
}

//ВРАЧА ОБРАБОТЧИКИ

func (h *Handler) acceptPatient(c echo.Context) error { //POST
	return nil
}

func (h *Handler) unAcceptPatient(c echo.Context) error { //DELETE
	return nil
}

func (h *Handler) ratePatient(c echo.Context) error { //POST
	return nil
}

func (h *Handler) createPost(c echo.Context) error { //POST
	return nil
}

func (h *Handler) editPost(c echo.Context) error { //PUT
	return nil
}

func (h *Handler) deletePost(c echo.Context) error { //DELETE
	return nil
}

//ПАЦИЕНТА ОБРАБОТЧИКИ

func (h *Handler) createRequest(c echo.Context) error { //POST
	return nil
}

func (h *Handler) deleteRequest(c echo.Context) error { //DELETE
	return nil
}

func (h *Handler) rateDoctor(c echo.Context) error { //POST
	return nil
}

func (h *Handler) createDiscussion(c echo.Context) error { //POST
	return nil
}

func (h *Handler) editDiscussion(c echo.Context) error { //PUT
	return nil
}

func (h *Handler) deleteDiscussion(c echo.Context) error { //DELETE
	return nil
}
