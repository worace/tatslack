package tatslack

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"time"
)

//DB stores a ref to bolt db
//containing slack data
type DB struct {
	db *bolt.DB
}

func Open(path string) (*DB, error) {
	d, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return nil, err
	}

	db := &DB{db: d}
	if err := db.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("messages"))
		return nil
	}); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

//close underlying connection
func (db *DB) Close() error {
	return db.db.Close()
}

//message returns a message by channel
func (db *DB) Messages(channel string) (a []*Message, err error) {
	db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages")).Bucket([]byte(channel))
		if b == nil {
			return nil
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			m := &Message{}
			if err := json.Unmarshal(v, &m); err != nil {
				return err
			}
			a = append(a, m)
		}
		return nil
	})
	return a, err
}

func (db *DB) SaveMessages(channel string, a []*Message) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.Bucket([]byte("messages")).CreateBucketIfNotExists([]byte(channel))
		if err != nil {
			return err
		}

		for _, m := range a {
			buf, err := json.Marshal(m)
			if err != nil {
				return err
			}

			if err := b.Put([]byte(m.TS), buf); err != nil {
				return err
			}
		}
		return nil
	})
}
