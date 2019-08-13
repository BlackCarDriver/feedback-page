import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'app-fbshow',
  templateUrl: './fbshow.component.html',
  styleUrls: ['./fbshow.component.css']
})

export class FbshowComponent implements OnInit {
  pageNumber = 1;
  presentPage = 1;
  targetImg = "";
  describe = "";
  server = "http://localhost:4700"
  fbdata: rows[] = [];

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.getFeedBackData(0);
  }

  //get data from database, ofs specified how many rows will be skip
  getFeedBackData(ofs: number) {
    let getdataUrl = this.server + "/feedback/getdata";
    let response = this.http.post<result2>(getdataUrl, JSON.stringify({ offset: ofs }));
    response.subscribe(result => {
      if (result.status == false) {
        alert("获取数据失败： " + result.describe);
      } else {
        this.fbdata = result.data;
        this.pageNumber = Math.ceil(result.sum/12);
        console.log(this.fbdata);
      }
    });
  }

  //after a feedback record have been read, update it record's state to 1
  updateState(fbid: number) {
    let updateUrl = this.server + "/feedback/readfb";
    let response = this.http.post<result>(updateUrl, JSON.stringify({ fbid: fbid }));
    response.subscribe(result => {
      if (result.status == false) {
        alert("更新反馈状态失败：" + result.describe);
      }
    });
  }

  nextPage(){
    if (this.presentPage>=this.pageNumber) return;
    this.presentPage+=1;
    this.getFeedBackData((this.presentPage-1)*12);
  }

  leastPage(){
    if (this.presentPage<=0) return;
    this.presentPage -=1;
    this.getFeedBackData((this.presentPage-1)*12);
  }

  //####################### tools function ############################

  setUpDesPanCont(imgurl: string, describe: string) {
    this.targetImg = imgurl;
    this.describe = describe;
    this.showDesPan();
  }
  hideDesPan() {
    let pan = (<HTMLInputElement>document.getElementById('showPan'));
    pan.style.display = "none";
  }
  showDesPan() {
    let pan = (<HTMLInputElement>document.getElementById('showPan'));
    pan.style.display = "initial";
  }

}

type result = {
  status: boolean;
  describe: string;
}

type result2 = {
  status: boolean;
  describe: string;
  data: rows[];
  sum: number;
}


type rows = {
  id: number;
  time: string;
  userid: string;
  pbtype: string;
  location: string;
  state: number;
  imgurl: string;
  describe: string;
}