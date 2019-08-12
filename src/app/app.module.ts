import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule} from '@angular/common/http';
import { AppComponent } from './app.component';
import { FbpageComponent } from './fbpage/fbpage.component';
import { FbshowComponent } from './fbshow/fbshow.component';

@NgModule({
  declarations: [
    AppComponent,
    FbpageComponent,
    FbshowComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
