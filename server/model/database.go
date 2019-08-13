package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	//fbimgPath is the prefix of feedback img url
	fbimgPath = "https://cst.gzhu.edu.cn/ards/static/feedbackimg/"
)

type FeedBackData struct {
	FbId     int64     `json:"id"`
	FbTime   time.Time `json:"time"`
	FbState  int64     `json:"state"`
	UserId   string    `json:"userid"`
	Location string    `json:"location"` //Where the problem occurred
	Type     string    `json:"pbtype"`
	Images   string    `json:"imgurl"` //screenshot saving name
	Describe string    `json:"describe"`
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

//insert a feedback record into database, time and state will set up by default value.
func SaveFeedBack(data *FeedBackData) error {
	if data == nil {
		return fmt.Errorf("Receive a null pointer!")
	}
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

//read at most 12 rows of feedback record from database
func GetFeedBack(offset int) (*[]FeedBackData, error) {
	if offset < 0 {
		return nil, fmt.Errorf("Offset can't small than 0!")
	}
	selectTP := `SELECT id, openid, pblocation, fbtype, images, describe, fbtime, fbstate FROM public.t_feedback order by fbtime desc limit 12 offset $1;`
	rows, err := db.Query(selectTP, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]FeedBackData, 0)
	for rows.Next() {
		t := FeedBackData{}
		err = rows.Scan(&t.FbId, &t.UserId, &t.Location, &t.Type, &t.Images, &t.Describe, &t.FbTime, &t.FbState)
		if err != nil {
			return nil, fmt.Errorf("Error when scan from rows: %v", err)
		}
		t.Images = fbimgPath + t.Images
		data = append(data, t)
	}
	return &data, nil
}

//change a feedback record's state to is read
func UpdateState(fbid int) error {
	if fbid < 0 {
		return fmt.Errorf("Receive a numebr small than 0!")
	}
	updateTP := `update t_feedback set fbstate=1 where id=$1`
	_, res := db.Exec(updateTP, fbid)
	return res
}

//count how many feedback record have save in database
func CountFbRecord() (int, error) {
	countTP := `select count(*) from t_feedback;`
	row := db.QueryRow(countTP)
	count := 0
	err := row.Scan(&count)
	return count, err
}
