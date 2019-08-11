import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-fbpage',
  templateUrl: './fbpage.component.html',
  styleUrls: ['./fbpage.component.css']
})


export class FbpageComponent implements OnInit {
  
  constructor() { }
  
  selectFileName = "未选择任意图片...";  //the name of images which is going upload 

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
    let location =  (<HTMLInputElement>document.getElementById('fblocation')).value;
    let email =  (<HTMLInputElement>document.getElementById('fbemail')).value;
    let describe =  (<HTMLInputElement>document.getElementById('fbdescribe')).value;
    let fileName =  (<HTMLInputElement>document.getElementById('inputfile')).value;
    let fileSize = (<HTMLInputElement>document.getElementById('inputfile')).files[0].size; ;
    const fileNameReg = /^.*\.(jpg)|(png)|(JPG)|(PNG)$/;
    if( fileName!=""){
       if (fileNameReg.test(fileName)==false ){
        alert("请上传正确格式的图片文件");
        return;
       }
       console.log(fileSize);
       if (fileSize > 3 << 20) {
         alert("请上传小于3M 的图片");
         return;
       }
    }
    let file =  (<HTMLInputElement>document.getElementById('inputfile')).files[0];
    console.log(type, location, email,file, describe);
   
  }

}
