package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

//feedback data
type FeedBackData struct {
	UserId   string `json:"userid"`
	Location string `json:"location"` //Where the problem occurred
	Type     string `json:"type"`
	Images   string `json:"images"` //screenshot saving name
	Describe string `json:"describe"`
}

func init() {
	var err error
	confstr := fmt.Sprintf()
	db, err = sql.Open("postgres", confstr)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connect Scuess!")
}

func SaveFeedBack(data *FeedBackData) error {
	insertTP := `INSERT INTO public.t_feedback(openid, pblocation, fbtype, images, describe)VALUES ($1,$2,$3,$4,$5);`
	res, err := db.Exec(insertTP, data.UserId, data.Location, data.Type, data.Images, data.Describe)
	if err != nil {
		return err
	}
	RoasAffected, _ := res.RowsAffected()
	if RoasAffected == 0 {
		return fmt.Errorf("No row was affacted!")
	}
	return nil
}
