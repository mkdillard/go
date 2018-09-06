package scroll

import (
  "fmt"
  "strings"

  "github.com/boltdb/bolt"
)

func CreateBuckets(db *bolt.DB) error {
  return db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("Dictionary"))
    if err != nil {
      return fmt.Errorf("create bucket: %s", err)
    }
    return nil
  })
}

func Read(db *bolt.DB, key string, value *string) error {
  return db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("Dictionary"))
    v := b.Get([]byte(key))
    *value = fmt.Sprintf("%v", v)
    return nil
  })
}

func Write(db *bolt.DB, message string) error {
  return db.Update(func(tx *bolt.Tx) error {
    inputs := strings.SplitN(message, " ", 4)
    b := tx.Bucket([]byte("Dictionary"))
    err := b.Put([]byte(inputs[2]), []byte(inputs[3]))
    return err
  })
}
