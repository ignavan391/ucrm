package api

import "github.com/go-chi/chi"

func RegisterRouter(r chi.Router, controller *Controller) {
	r.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/sendVerifyCode", controller.SendVerifyCode)
			r.Post("/signUp", controller.SignUp)
			r.Post("/signIn", controller.SignIn)
		})
	})
}
