package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TartuDen/webPage4_starting_project/internal/config"
	"github.com/TartuDen/webPage4_starting_project/internal/forms"
	"github.com/TartuDen/webPage4_starting_project/internal/models"
	"github.com/TartuDen/webPage4_starting_project/internal/renderer"
)

// TemplateData holds data sent from handlers to templates

// Repo the repository used by the handlers
var Repo *Repository

// Repositroy is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// MainHandler is a method of the Repository struct that handles requests to the main page.
// It renders the "home.page.html" template to the provided HTTP response writer.
func (m *Repository) MainHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renderer.RendererTemplate(w, "home.page.html", r, &models.TemplateData{})
}

// AboutHandler is a method of the Repository struct that handles requests to the about page.
// It renders the "about.page.html" template to the provided HTTP response writer.
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "this is test data!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send data to the template
	renderer.RendererTemplate(w, "about.page.html", r, &models.TemplateData{
		StringMap: stringMap,
	})
}

// DensRoomHandler renders the room page
func (m *Repository) DensRoomHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "densroom.page.html", r, &models.TemplateData{})
}

// YurecRoomHandler renders the room page
func (m *Repository) YurecRoomHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "yurecroom.page.html", r, &models.TemplateData{})
}

// //YurecRoomHandler renders the room page
// func (m *Repository) YurecPostRoomHandler(w http.ResponseWriter, r *http.Request){
// 	// renderer.RendererTemplate(w, "yurecroom.page.html", r, &models.TemplateData{})
// 	w.Write([]byte("Posted to search availability"))
// }

// AvailableHandler renders the room page
func (m *Repository) AvailableHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "search-availability.page.html", r, &models.TemplateData{})
}

// ContactHandler renders the room page
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "contact.page.html", r, &models.TemplateData{})
}

// BookHandler renders the room page
func (m *Repository) BookHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "bookNow.page.html", r, &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// BookJSON handles request for availability and sends Json response
func (m *Repository) BookJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// BookHandler renders the room page
func (m *Repository) BookPostHandler(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	w.Write([]byte(start + " - " + end))

	// renderer.RendererTemplate(w, "bookNow.page.html", r, &models.TemplateData{})
	w.Write([]byte("Posted on handler"))
}

// ReservationHandler renders the room page
func (m *Repository) ReservationHandler(w http.ResponseWriter, r *http.Request) {
	var emptyReserv models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReserv
	renderer.RendererTemplate(w, "make-reservation.page.html", r, &models.TemplateData{
		Form: forms.NewForm(nil),
		Data: data,
	})
}

// PostMakeReservation handles the POsting of the reservation form
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	reservation := models.Reservation{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone_number"),
	}

	form := forms.NewForm(r.PostForm)

	// form.Has("FirstName",r)

	form.Required("first_name", "last_name", "email")
	form.MinLen("first_name", 4, r)
	form.MinLen("last_name", 4, r)
	form.MaxLen("first_name", 10, r)
	form.MaxLen("last_name", 10, r)

	//manually written valid func for email
	form.EmailFormat("email",r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renderer.RendererTemplate(w, "make-reservation.page.html", r, &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}
