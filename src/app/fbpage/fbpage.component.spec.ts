import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FbpageComponent } from './fbpage.component';

describe('FbpageComponent', () => {
  let component: FbpageComponent;
  let fixture: ComponentFixture<FbpageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FbpageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FbpageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
