package repository

import (
	"codeticket/internal/events/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlEventRepository struct {
	db *sql.DB
}

// // NewMysqlEventRepository creates a new MySQL event repository.
// func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
// 	return &mysqlEventRepository{db: db}, nil
// }

// func (r *mysqlEventRepository) ListEvents() ([]domain.Event, error) {}

// // FindEventByID returns an event by its ID, including associated spots and tickets.
// func (r *mysqlEventRepository) FindEventByID(eventID string) (*domain.Event, error) {}

// // CreateEvent inserts a new event into the database.
func (r *mysqlEventRepository) CreateEvent(event *domain.Event) error {
	query := `
		INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, event.ID, event.Name, event.Location, event.Organization, event.Rating, event.Date.Format("2006-01-02 15:04:05"), event.ImageURL, event.Capacity, event.Price, event.PartnerID)
	return err
}

// // FindSpotByID returns a spot by its ID, including the associated ticket (if any).
// func (r *mysqlEventRepository) FindSpotByID(spotID string) (*domain.Spot, error) {}

// // CreateSpot inserts a new spot into the database.
func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_id, name, status, ticket_id)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}

// // CreateTicket inserts a new ticket into the database.
func (r *mysqlEventRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, event_id, spot_id, ticket_kind, price)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketKind, ticket.Price)
	return err
}

// // ReserveSpot updates a spot's status to reserved.
// func (r *mysqlEventRepository) ReserveSpot(spotID, ticketID string) error {}

// // FindSpotsByEventID returns all spots for a given event ID.
// func (r *mysqlEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {}

// func (r *mysqlEventRepository) FindSpotByName(eventID, name string) (*domain.Spot, error) {}
