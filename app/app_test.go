package main

import (
	"context"
	"testing"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/providers"
)

func TestAddPost(t *testing.T) {
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)
	for id := 0; id < 100000; id++ {
		_, err := entiry.SetPost(db.AddNewPostParams{})
		if err != nil {
			t.Error(err)
		} else {
			// fmt.Print(post)
			t.Log(err)
		}
	}

}

func TestListPost(t *testing.T) {
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)
	_, err := entiry.ListPosts(context.Background())
	if err != nil {
		t.Error(err)
	} else {

		t.Log(err)
	}

}

func TestGetOnePost(t *testing.T) {
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)

	_, err := entiry.GetOnePost(context.Background(), 3)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(err)
	}

}
