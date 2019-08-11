package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"

	"./model"
)

var (
	fbImgDir string //the path saving feedback upload images
)

func init() {
	fbImgDir = "./fbimages"
}

// the feedback result strcut return to user
type fbStatus struct {
	Status   bool   `json:"status"`
	Describe string `json:"descirbe"`
}

func main() {
	http.HandleFunc("/feedback/postfeedback", FeedBackHandle)
	http.HandleFunc("/", Test)
	fmt.Println("the server is running...")
	err := http.ListenAndServe("localhost:4700", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("...")
}

func FeedBackHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("AllowMethods", "POST")
	w.Header().Set("content-type", "application/json; charset=utf-8")
	var err error
	result := fbStatus{}
	fbdata := model.FeedBackData{}

	err = r.ParseMultipartForm(5 << 20)
	if err != nil {
		fmt.Println(err)
		result.Status = false
		result.Describe = fmt.Sprintf("Http Request ParseMutipartForm false: %v", err)
		goto tail
	}

	//get feedback data from postbody and check the value latter
	fbdata.Type = getMultipartFormValue(r.MultipartForm, "type")
	fbdata.UserId = getMultipartFormValue(r.MultipartForm, "userid")
	fbdata.Location = getMultipartFormValue(r.MultipartForm, "location")
	fbdata.Describe = getMultipartFormValue(r.MultipartForm, "describe")
	if checkRes := checkFbData(&fbdata); checkRes != "" {
		result.Status = false
		result.Describe = fmt.Sprintf("The feedback received data was incomplete: %s", checkRes)
		goto tail
	}

	//handle the uploaded screenshot images if have
	if imgs, _ := r.MultipartForm.File["images"]; len(imgs) > 0 {
		img := imgs[0]
		size := img.Size
		name := img.Filename
		//check file size
		if size > 3<<20 {
			result.Status = false
			result.Describe = fmt.Sprint("The size of upload images is too larger! please upload images that no more than 3mb!")
			goto tail
		}
		reg, _ := regexp.Compile(`(!?)^[^\.]+\.(jpg)|(png)|(jpeg)$`)
		//check file name
		if !reg.MatchString(name) {
			result.Status = false
			result.Describe = fmt.Sprintf("Unexpect file name!")
			goto tail
		}
		//read file
		tmpfile, err := img.Open()
		if err != nil {
			result.Status = false
			result.Describe = fmt.Sprintf("Can not read uploaded images!")
			goto tail
		}
		defer tmpfile.Close()
		//change to another name and save
		newFileName := fmt.Sprintf("%s%s", GetRandomString(20), path.Ext(name))
		cur, err := os.Create(fmt.Sprintf("%s/%s", fbImgDir, newFileName))
		if err != nil {
			result.Status = false
			result.Describe = fmt.Sprintf("Save images faill! :%v", err)
			goto tail
		}
		io.Copy(cur, tmpfile)
		cur.Close()
		fbdata.Images = newFileName
		fmt.Println(size, newFileName)
	}

	err = model.SaveFeedBack(&fbdata)
	if err != nil {
		fmt.Println(err)
	}

	result.Status = true
tail:
	if result.Status == false {
		fmt.Println(result.Describe)
	}
	jsonResp, _ := json.Marshal(result)
	_, err = w.Write(jsonResp)
	if err != nil {
		fmt.Println(err)
	}
}

func checkFbData(data *model.FeedBackData) string {
	if data.UserId == "" {
		return "Userid not found!"
	}
	if data.Location == "" {
		return "problem location is empty!"
	}
	if data.Type == "" {
		return "feedback type not found!"
	}
	if data.Describe == "" {
		return "problem description is empty!"
	}
	return ""
}

//get first value in mutipaart form according to key
func getMultipartFormValue(f *multipart.Form, key string) string {
	arrays := f.Value[key]
	if len(arrays) == 0 {
		return ""
	}
	return arrays[0]
}

//create an random string that with length l
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
