package persistence

import (
	"database/sql"
	"errors"
	"maker_checker_sqlite/src/domain/entities"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

const dbDir = "data"         // Directory for the database
const dbFile = "messages.db" // SQLite database file

type SQLiteMessageRepository struct {
	db *sql.DB
}

// NewSQLiteMessageRepository initializes a SQLite-based message repository
func NewSQLiteMessageRepository(dbFilePath string) (*SQLiteMessageRepository, error) {
	// Ensure the data directory exists
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		return nil, err
	}

	// Full path to the database file
	fullPath := filepath.Join(dbDir, dbFile)

	db, err := sql.Open("sqlite3", fullPath)
	if err != nil {
		return nil, err
	}

	repo := &SQLiteMessageRepository{db: db}
	err = repo.init()
	return repo, err
}

// init initializes the database schema
func (repo *SQLiteMessageRepository) init() error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id TEXT PRIMARY KEY,
		content TEXT,
		recipient TEXT,
		status TEXT,
		approvers TEXT,
		max_approvers INTEGER
	);
	`
	_, err := repo.db.Exec(query)
	return err
}

// Save saves a message to the database
func (repo *SQLiteMessageRepository) Save(message *entities.Message) error {
	query := `
	INSERT OR REPLACE INTO messages (id, content, recipient, status, approvers, max_approvers)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := repo.db.Exec(query, message.ID, message.Content, message.Recipient, message.Status, message.Approvers, message.MaxApprovers)
	return err
}

// FindByID retrieves a message by its ID
func (repo *SQLiteMessageRepository) FindByID(id string) (*entities.Message, error) {
	query := `SELECT id, content, recipient, status, approvers, max_approvers FROM messages WHERE id = ?`
	row := repo.db.QueryRow(query, id)

	var message entities.Message
	err := row.Scan(&message.ID, &message.Content, &message.Recipient, &message.Status, &message.Approvers, &message.MaxApprovers)
	if err == sql.ErrNoRows {
		return nil, errors.New("message not found")
	}
	return &message, err
}
