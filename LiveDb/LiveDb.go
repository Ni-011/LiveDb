package livedb

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	AdminDB = "Admin"
)

type collection struct {
	bucket *bbolt.Bucket
}

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

// creating collections
func (ldb LiveDb) createCollection (name string) (*collection, error) {
	collection := collection{} // collection to store the user created bucket
	// create a read-write bucket with user sent name
	err := ldb.DB.Update(func (tx *bbolt.Tx) error {
		bucket, error := tx.CreateBucket([]byte(name))
		if error != nil {
			return error
		}
		collection.bucket = bucket
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &collection, nil // return the collection
}



// // create a bucket in the db, add the user info
// 	DB.Update(func(tx *bbolt.Tx) error {
// 		id := uuid.New() // uuid id of user

// 		// put the users info into the bucket
// 		for key, val := range User {
// 			err := bucket.Put([]byte(key), []byte(val))
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		// addding the id to the bucket
// 		if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
// 			return err
// 		}
// 		return nil
// 	})