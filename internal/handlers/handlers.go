package handlers

import (
	"net/http"

	"github.com/test/internal/config"
	"github.com/test/internal/driver"
	"github.com/test/internal/forms"
	"github.com/test/internal/models"
	"github.com/test/internal/render"
	"github.com/test/internal/repository"
	"github.com/test/internal/repository/dbrepo"
)

// Initialize the repository for the application
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// Function to create a new test Repository
func NewTestingRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) ShowHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}
