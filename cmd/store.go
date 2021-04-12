package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/status-im/go-waku/waku/v2/protocol"
	store "github.com/status-im/go-waku/waku/v2/protocol/waku_store"
)

type DBStore struct {
	store.MessageProvider
	db *sql.DB
}

func NewDBStore(path string) (*DBStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	result := &DBStore{db: db}

	err = result.createTable()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DBStore) createTable() error {
	sqlStmt := `CREATE TABLE IF NOT EXISTS messages (
		id BLOB PRIMARY KEY,
		timestamp INTEGER NOT NULL,
		contentTopic BLOB NOT NULL,
		payload BLOB,
		version INTEGER NOT NULL DEFAULT 0
	) WITHOUT ROWID;`
	_, err := d.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBStore) Stop() {
	d.db.Close()
}

func (d *DBStore) Put(cursor *protocol.Index, message *protocol.WakuMessage) error {
	stmt, err := d.db.Prepare("INSERT INTO messages (id, timestamp, contentTopic, payload, version) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cursor.Digest, uint64(message.Timestamp), message.ContentTopic, message.Payload, message.Version)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBStore) GetAll() ([]*protocol.WakuMessage, error) {
	rows, err := d.db.Query("SELECT timestamp, contentTopic, payload, version FROM messages ORDER BY timestamp ASC")
	if err != nil {
		return nil, err
	}

	var result []*protocol.WakuMessage

	defer rows.Close()

	for rows.Next() {
		var timestamp int64
		var contentTopic string
		var payload []byte
		var version uint32

		err = rows.Scan(&timestamp, &contentTopic, &payload, &version)
		if err != nil {
			log.Fatal(err)
		}

		msg := new(protocol.WakuMessage)
		msg.ContentTopic = contentTopic
		msg.Payload = payload
		msg.Timestamp = float64(timestamp)
		msg.Version = version

		result = append(result, msg)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
