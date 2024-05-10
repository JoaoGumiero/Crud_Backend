package postgres

import (
	"context"
	"fmt"

	t "github.com/JoaoGumiero/Crud_Backend/Ticket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TicketDAO struct {
	db *pgxpool.Pool
}

func NewTicketDAO(db *pgxpool.Pool) *TicketDAO {
	return &TicketDAO{db: db}
}

func (r *TicketDAO) CreateTicketDAO(ctx context.Context, NewTicket t.Ticket) (*t.Ticket, error) {
	query := `INSERT INTO Tickets (id, title, analysis_date, solving_date, description, sender_queue, reciever_queue, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := r.db.QueryRow(ctx, query).Scan(&NewTicket.ID)
	if err != nil {
		return nil, err
	}
	return &NewTicket, nil
}

func (r *TicketDAO) UpdateTicketDAO(ctx context.Context, NewTicket t.Ticket, id string) (*t.Ticket, error) {
	query := `UPDATE Tickets SET title = $2, analysis_date = $3, solving_date = $4, description = $5, sende_queue = $6, reciever_queue = $7, status = $8 Where id = $1`
	err := r.db.QueryRow(ctx, query, id, NewTicket.Title, NewTicket.Analysis_Date, NewTicket.Solving_Date, NewTicket.Description, NewTicket.SenderQeue, NewTicket.RecieverQeue, NewTicket.Status).Scan(&NewTicket.ID)
	if err != nil {
		return nil, err
	}
	return &NewTicket, nil
}

func (r *TicketDAO) GetTicketByIdDAO(ctx context.Context, id string) (*t.Ticket, error) {
	var ticket t.Ticket
	query := `SELECT * FROM Ticket Where ID = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&ticket.ID, &ticket.Title, &ticket.Analysis_Date, &ticket.Solving_Date, &ticket.Description, &ticket.SenderQeue, &ticket.RecieverQeue, &ticket.Status)
	if err != nil {
		return nil, err
	}
	return &ticket, err
}

func (r *TicketDAO) GetAllTicketsDAO(ctx context.Context) error {
	query := `SELECT * FROM Tickets`
	_, err := r.db.Query(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketDAO) DeleteTicketDAO(ctx context.Context, id string) error {
	query := `DELETE * FROM Tickets Where id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketDAO) Migrate() error {
	query :=
		`CREATE TYPE ticket_status AS ENUM (
		'Open',
		'InProgress',
		'Solved',
		'Closed'
	);
	
	CREATE TABLE tickets (
		ID            VARCHAR(255) PRIMARY KEY,
		Title         VARCHAR(255) NOT NULL,
		Analysis_Date TIMESTAMP NOT NULL,
		Solving_Date  TIMESTAMP NOT NULL,
		Description   TEXT,
		SenderQueue   VARCHAR(255),
		ReceiverQueue VARCHAR(255),
		Status        ticket_status NOT NULL
	);`
	fmt.Println("Creating status type and tickets table...")
	_, err := r.db.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}
