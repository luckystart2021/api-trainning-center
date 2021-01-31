package admin

import (
	"api-trainning-center/handlers/api/admin/account"
	"api-trainning-center/handlers/api/admin/contact"
	"api-trainning-center/handlers/api/admin/course"
	"api-trainning-center/handlers/api/admin/question"
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
	}
}
