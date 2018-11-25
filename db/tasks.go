package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

//keys in bolt are byte slices
var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// check that this is not 'init' and does not run automatically.
//This is run everytime the function is called for the first time
//from the main.go file
func Init(dbPath string) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	// db.Update takes in a function that checks to see if have the bucket
	// created already or not
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// connect to the bucket
		b := tx.Bucket(taskBucket)
		// get the next possible id
		id64, _ := b.NextSequence()
		id = int(id64)
		//put the thing into the database.
		key := itob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int {
	return int(binary.BigEndian.Uint64(v))
}
