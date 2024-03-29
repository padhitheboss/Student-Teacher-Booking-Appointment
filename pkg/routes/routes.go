package routes

//Register Routes
import (
	"github.com/StudentTeacher-Booking-Appointment/pkg/controller"
	"github.com/StudentTeacher-Booking-Appointment/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterRoute(r chi.Router) {
	r.Post("/login", controller.Login)
	r.Post("/signup", controller.StudentSignup)
	r.Get("/logout", controller.Logout)
	r.Get("/", middleware.RequireAuth(controller.Home, "student"))
	r.Get("/search", middleware.RequireAuth(controller.SearchTeacher, "student"))
	r.Get("/approve", middleware.RequireAuth(controller.ListUnApprovedStudent, "admin"))
	r.Post("/approve", middleware.RequireAuth(controller.ApproveStudent, "admin"))
	r.Post("/addteacher", middleware.RequireAuth(controller.AddTeacher, "admin"))
	r.Post("/book", middleware.RequireAuth(controller.BookAppointment, "student"))
	r.Get("/getappointments", middleware.RequireAuth(controller.ListAppointments, "teacher"))
	r.Post("/scheduleappointment", middleware.RequireAuth(controller.ScheduleAppointment, "teacher"))
}
