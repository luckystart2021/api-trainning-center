package admin

import (
	"api-trainning-center/handlers/api/admin/account"
	"api-trainning-center/handlers/api/admin/child_category"
	"api-trainning-center/handlers/api/admin/contact"
	"api-trainning-center/handlers/api/admin/course"
	"api-trainning-center/handlers/api/admin/information"
	"api-trainning-center/handlers/api/admin/news"
	"api-trainning-center/handlers/api/admin/question"
	"api-trainning-center/handlers/api/admin/slide"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	return func(router chi.Router) {
		router.Group(account.RouterLogin(db, client))
		router.Group(contact.Router(db, client))
		router.Group(course.CourseRoute(db, client))
		router.Group(question.Router(db, client))
		router.Group(information.InfoRouter(db, client))
		router.Group(news.Router(db, client))
		router.Group(child_category.Router(db, client))
		router.Group(slide.Router(db, client))
	}
}
