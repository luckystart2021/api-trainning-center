package user

import (
	"api-trainning-center/handlers/api/user/about"
	"api-trainning-center/handlers/api/user/contact"
	"api-trainning-center/handlers/api/user/information"
	"api-trainning-center/handlers/api/user/news"
	"api-trainning-center/handlers/api/user/photo"
	"api-trainning-center/handlers/api/user/question"
	"api-trainning-center/handlers/api/user/schedule"
	"api-trainning-center/handlers/api/user/seo"
	"api-trainning-center/handlers/api/user/slide"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		router.Group(contact.Router(db))
		router.Group(information.Router(db))
		router.Group(question.Router(db))
		router.Group(about.Router(db))
		router.Group(news.Router(db))
		router.Group(slide.Router(db))
		router.Group(seo.Router(db))
		router.Group(photo.Router(db))
		router.Group(schedule.Router(db))
	}
}
