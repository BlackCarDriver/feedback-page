import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FbshowComponent } from './fbshow.component';

describe('FbshowComponent', () => {
  let component: FbshowComponent;
  let fixture: ComponentFixture<FbshowComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FbshowComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FbshowComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
