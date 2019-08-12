import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-fbshow',
  templateUrl: './fbshow.component.html',
  styleUrls: ['./fbshow.component.css']
})
export class FbshowComponent implements OnInit {
  pageNumber =  10;
  presentPage = 1;
  targetImg = "https://blackcardriver.cn/images/1.jpg";
  describe = "你好，刚才在刷新浏览器的时候页面突然蹦了，刷新多次依然是这样，不知道为什么，请速速解决！";

  mockdata : rows[] = [
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
      {"id":123123, "time":"2019-11 09:11", "userid":"off364i2R67h6AgSnvKthIiNqwrs","pbtype":"登录问题","location":"https://github.com/BlackCarDriver", "state":1},
  ]

  constructor() { }

  ngOnInit() {
  }

}


type rows = {
    id:number;
    time: string;
    userid:string;
    pbtype:string;
    location:string;
    state:number;
}