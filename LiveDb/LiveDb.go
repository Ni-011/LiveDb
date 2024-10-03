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
	collection := collection{} // collection to store the user created bucket
	// create a read-write bucket with user sent name
	err := ldb.DB.Update(func(tx *bbolt.Tx) error {
		var (
			err    error
			bucket *bbolt.Bucket
		)
		bucket = tx.Bucket([]byte(name)) // search for the bucket
		if bucket == nil {               // if bucket doesn't exist, make it
			bucket, err = tx.CreateBucket([]byte(name))
			if err != nil {
				return err
			}
		}
		collection.Bucket = bucket // update collection with the bucket found/made
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &collection, nil // return the collection
}

func (ldb *LiveDb) Insert(collectionName string, data Data) (uuid.UUID, error) {
	id := uuid.New()

	// get/make the collection user wants
	collection, err := ldb.CreateCollection(collectionName)
	if err != nil {
		return id, err
	}

	// inserting user's data into the collection
	ldb.DB.Update(func(tx *bbolt.Tx) error { // read/write method
		for key, value := range data {
			if err := collection.Put([]byte(key), []byte(value)); err != nil {
				return err
			}
		}

		// putting id into the collection
		if err := collection.Put([]byte("id"), []byte(id.String())); err != nil {
			return err
		}
		return nil
	})

	return id, nil
}

func (lbd *LiveDb) Select (collection string, key string, query any) {

}
