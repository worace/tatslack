package tatslack_test

import (
	"github.com/worace/tatslack"
	"io/ioutil"
	"os"
	"testing"
)

//ensure we can save stuff
func TestDB_SaveMessages(t *testing.T) {
	db := OpenDB()
	defer db.Close()

	a := []*tatslack.Message{
		{Type: "message", UserID: "1234", Text: "heres a message", TS: "142141414193832.14141"},
		{Type: "message", UserID: "14141", Text: "mess 2", TS: "1421414141.141488921"},
		{Type: "message", UserID: "123454634", Text: "mess 3", TS: "142141414858.14141"},
		{Type: "message", UserID: "1234636k4", Text: "message 4", TS: "14214123141.14141"},
	}

	channel := "C123456"
	if err := db.SaveMessages(channel, a); err != nil {
		t.Fatal(err)
	}

	messages, err := db.Messages(channel)
	if err != nil {
		t.Fatal(err)
	} else if len(messages) != 4 {
		t.Fatalf("unexepected len: %d", len(messages))
	}
}

func OpenDB() *tatslack.DB {
	db, err := tatslack.Open(tempfile())
	if err != nil {
		panic(err)
	}
	return db
}

func tempfile() string {
	f, _ := ioutil.TempFile("", "tatslack")
	f.Close()
	os.Remove(f.Name())
	return f.Name()
}
