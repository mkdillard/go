package scroll

import (
  "fmt"

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
    if b != nil {
      v := b.Get([]byte(key))
      if v != nil {
       *value = bytesToString(v)
      }
    } else {
      fmt.Println("Dictionary bucket was nil")
    }
    return nil
  })
}

func Write(db *bolt.DB, key string, value string) error {
  return db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("Dictionary"))
    if b != nil {
      err := b.Put([]byte(key), []byte(value))
      return err
    }
    fmt.Println("bucket Dictionary was nil")
    return nil
  })
}

func bytesToString(data []byte) string {
  return string(data[:])
}
