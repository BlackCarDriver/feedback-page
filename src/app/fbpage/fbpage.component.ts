import { Component, OnInit } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-fbpage',
  templateUrl: './fbpage.component.html',
  styleUrls: ['./fbpage.component.css']
})


export class FbpageComponent implements OnInit {
  
  constructor(private http: HttpClient) { }
  
  userid = "temp userid";    //user's open id, get from other page
  location = "temp location";  //where the problem come from, get from other page
  selectFileName = "未选择任意图片...";  //the name of images which is going upload 
  server = "http://localhost:4700"

  ngOnInit() {
    var inputs = document.getElementById('inputfile');
    inputs.addEventListener('change', this.showImgName.bind(this));
    var form = document.getElementById('fbform');
    form.addEventListener('submit', this.commit.bind(this));
  }

  showImgName(){
    var fileName = (<HTMLInputElement>document.getElementById('inputfile')).value;
    this.selectFileName = fileName;
  }

  commit(){
    let type =  (<HTMLInputElement>document.getElementById('fbtype')).value;
    // let location =  (<HTMLInputElement>document.getElementById('fblocation')).value;
    // let email =  (<HTMLInputElement>document.getElementById('fbemail')).value;
    let describe =  (<HTMLInputElement>document.getElementById('fbdescribe')).value;
    let fileName =  (<HTMLInputElement>document.getElementById('inputfile')).value;
    let images : any
    const fileNameReg = /(!?)^.*\.(jpg)|(png)|(jpeg)$/;
    if( fileName!=""){
      if (fileNameReg.test(fileName)==false ){
        alert("请上传正确格式的图片文件");
        return;
      }
      let fileSize = (<HTMLInputElement>document.getElementById('inputfile')).files[0].size; ;
       console.log(fileSize);
       if (fileSize > 3 << 20) {
         alert("请上传小于3M 的图片");
         return;
       }
       images =  (<HTMLInputElement>document.getElementById('inputfile')).files[0];
    }

    let input = new FormData();
    input.append('userid', this.userid);
    input.append('location', this.location);
    input.append('type', type);
    input.append('images', images);
    input.append('describe', describe);

    let feedbackUrl = this.server +"/feedback/postfeedback";
    let response = this.http.post<string>(feedbackUrl, input);
    response.subscribe(result=>{alert(result);});
  }

}
