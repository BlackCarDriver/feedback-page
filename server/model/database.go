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
	Id       int64     `json:"id"`
	UserId   string    `json:"userid"`
	Email    string    `json:"email"`
	Time     time.Time `json:"time"`
	Status   int64     `json:"state"`
	Type     string    `json:"pbtype"`
	Location string    `json:"location"` //Where the problem occurred
	Describe string    `json:"describe"`
	Imgurl   string    `json:"imgurl"` //screenshot saving name
}

func init() {
	var err error
	confstr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	)
	db, err = sql.Open("postgres", confstr)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connect Scuess!")
}

//insert a feedback record into dbase, time and state will set up by default value.
func SaveFeedBack(d *FeedBackData) error {
	if d == nil {
		return fmt.Errorf("Receive a null pointer!")
	}
	insertTP := `INSERT INTO public.t_feedback2(user_id, fb_location, fb_type, images_name, describe, email)VALUES ($1,$2,$3,$4,$5,$6);`
	res, err := db.Exec(insertTP, d.UserId, d.Location, d.Type, d.Imgurl, d.Describe, d.Email)
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
	selectTP := `SELECT id, user_id, fb_location, fb_type, images_name, describe, fb_time, fb_status, email  FROM public.t_feedback2 order by fb_time desc limit 12 offset $1;`
	rows, err := db.Query(selectTP, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]FeedBackData, 0)
	for rows.Next() {
		t := FeedBackData{}
		imgName := ""
		err = rows.Scan(&t.Id, &t.UserId, &t.Location, &t.Type, &imgName, &t.Describe, &t.Time, &t.Status, &t.Email)
		if err != nil {
			return nil, fmt.Errorf("Error when scan from rows: %v", err)
		}
		t.Imgurl = fbimgPath + imgName
		data = append(data, t)
	}
	return &data, nil
}

//change a feedback record's state to is read
func UpdateState(fbid int) error {
	if fbid < 0 {
		return fmt.Errorf("Receive a numebr small than 0!")
	}
	updateTP := `update t_feedback2 set fb_status=1 where id=$1`
	_, res := db.Exec(updateTP, fbid)
	return res
}

//count how many feedback record have save in database
func CountFbRecord() (int, error) {
	countTP := `select count(*) from t_feedback2;`
	row := db.QueryRow(countTP)
	count := 0
	err := row.Scan(&count)
	return count, err
}
