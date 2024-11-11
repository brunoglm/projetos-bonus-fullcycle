package repository

import (
	"codeticket/internal/events/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlEventRepository struct {
	db *sql.DB
}

// NewMysqlEventRepository creates a new MySQL event repository.
func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
	return &mysqlEventRepository{db: db}, nil
}

func (r *mysqlEventRepository) ListEvents() ([]domain.Event, error) {}

// FindEventByID returns an event by its ID, including associated spots and tickets.
func (r *mysqlEventRepository) FindEventByID(eventID string) (*domain.Event, error) {}

// CreateEvent inserts a new event into the database.
func (r *mysqlEventRepository) CreateEvent(event *domain.Event) error {}

// FindSpotByID returns a spot by its ID, including the associated ticket (if any).
func (r *mysqlEventRepository) FindSpotByID(spotID string) (*domain.Spot, error) {}

// CreateSpot inserts a new spot into the database.
func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {}

// CreateTicket inserts a new ticket into the database.
func (r *mysqlEventRepository) CreateTicket(ticket *domain.Ticket) error {}

// ReserveSpot updates a spot's status to reserved.
func (r *mysqlEventRepository) ReserveSpot(spotID, ticketID string) error {}

// FindSpotsByEventID returns all spots for a given event ID.
func (r *mysqlEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {}

func (r *mysqlEventRepository) FindSpotByName(eventID, name string) (*domain.Spot, error) {}
