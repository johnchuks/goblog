package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/johnchuks/goblog/accountservice/model"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) []byte
	Seed()
	UpdateAccount(accountId string) ([]byte, error)
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) seedAccounts() {

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

// "QueryAccount" Get Account information from bucket with accountID
func (bc *BoltClient) QueryAccount(accountId string) (val []byte) {
	bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		val = b.Get([]byte(accountId))
		return nil
	})
	return val
}

func (bc *BoltClient) UpdateAccount(accountId string) (val []byte, err error) {
		err = bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(accountId), []byte("Johnbosco"))
			if err != nil {
				return err
			}
			val = b.Get([]byte(accountId))
			return nil
		})
		if err != nil {
			return nil, err
		}
		return val, nil
} 
