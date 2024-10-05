package LiveDb

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	AdminDB = "Admin"
)

type Data map[string]string

// a cluster of buckets
type collection struct {
	*bbolt.Bucket
}

// the main db
type LiveDb struct {
	DB *bbolt.DB
}

// constructor to initiate the LiveDb instance
func NewLiveDB() (*LiveDb, error) {
	// creating a bbolt db
	DbName := fmt.Sprintf("%s.LiveDb", AdminDB)
	DB, err := bbolt.Open(DbName, 0666, nil)
	if err != nil {
		return nil, err
	}
	// creating the LiveDb instance with its values
	return &LiveDb{
		DB: DB,
	}, nil
}

// creating collecions,
func (ldb LiveDb) CreateCollection(name string) (*collection, error) {
	tx, err := ldb.DB.Begin(true) // read/write method
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // defer the commit

	Bucket, err := tx.CreateBucketIfNotExists([]byte(name)) // find or create the bucket
	if err != nil {
		return nil, err
	}

	return &collection{Bucket}, nil
}

func (ldb *LiveDb) Insert(collectionName string, data Data) (uuid.UUID, error) {
	id := uuid.New()

	tx, err := ldb.DB.Begin(true) // read/write method
	if err != nil {
		return id, err
	}
	defer tx.Rollback() // defer the commit

	Bucket, err := tx.CreateBucketIfNotExists([]byte(collectionName)) // find or create the bucket
	if err != nil {
		return id, err
	}

	// inserting user's data into the bucket 
	for key, value := range data {
		if err := Bucket.Put([]byte(key), []byte(value)); err != nil {
			return id, err
		}
	}

	// putting id into the collection
	if err := Bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}

	return id, tx.Commit()
}


// return bucket data using a query, read only
func (lbd *LiveDb) Select(collection string,  query Data) (Data, error) {
	tx, err := lbd.DB.Begin(true) // read/write method
	if err != nil {
		return nil, err
	}

	bucket := tx.Bucket([]byte(collection)) // find user requested bucket
	if bucket == nil {
		// bucket must exist
		return nil, fmt.Errorf("collection %s not found", collection)
	}
}
