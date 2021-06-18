import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestComponent } from './ptest.component';

describe('PtestComponent', () => {
  let component: PtestComponent;
  let fixture: ComponentFixture<PtestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
