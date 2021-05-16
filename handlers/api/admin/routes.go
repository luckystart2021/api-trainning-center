package admin

import (
	"api-trainning-center/handlers/api/admin/account"
	"api-trainning-center/handlers/api/admin/album"
	"api-trainning-center/handlers/api/admin/child_category"
	"api-trainning-center/handlers/api/admin/child_subject"
	"api-trainning-center/handlers/api/admin/class"
	"api-trainning-center/handlers/api/admin/contact"
	"api-trainning-center/handlers/api/admin/course"
	"api-trainning-center/handlers/api/admin/fee"
	"api-trainning-center/handlers/api/admin/holiday"
	"api-trainning-center/handlers/api/admin/information"
	"api-trainning-center/handlers/api/admin/news"
	"api-trainning-center/handlers/api/admin/photo"
	"api-trainning-center/handlers/api/admin/question"
	"api-trainning-center/handlers/api/admin/schedule"
	"api-trainning-center/handlers/api/admin/seo"
	"api-trainning-center/handlers/api/admin/seo_tag"
	"api-trainning-center/handlers/api/admin/slide"
	"api-trainning-center/handlers/api/admin/student"
	"api-trainning-center/handlers/api/admin/subject"
	"api-trainning-center/handlers/api/admin/suite"
	"api-trainning-center/handlers/api/admin/teacher"
	"api-trainning-center/handlers/api/admin/training_cost"
	"api-trainning-center/handlers/api/admin/upload"
	"api-trainning-center/handlers/api/admin/vehicle"
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
		router.Group(suite.Router(db, client))
		router.Group(class.Router(db, client))
		router.Group(student.Router(db, client))
		router.Group(seo.Router(db, client))
		router.Group(seo_tag.Router(db, client))
		router.Group(vehicle.Router(db, client))
		router.Group(album.Router(db, client))
		router.Group(upload.Router(client))
		router.Group(photo.Router(db, client))
		router.Group(teacher.Router(db, client))
		router.Group(schedule.Router(db, client))
		router.Group(subject.Router(db, client))
		router.Group(child_subject.Router(db, client))
		router.Group(holiday.Router(db, client))
		router.Group(fee.Router(db, client))
		router.Group(training_cost.Router(db, client))
	}
}
