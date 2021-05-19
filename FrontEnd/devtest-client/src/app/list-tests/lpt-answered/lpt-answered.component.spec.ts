import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptAnsweredComponent } from './lpt-answered.component';

describe('LptAnsweredComponent', () => {
  let component: LptAnsweredComponent;
  let fixture: ComponentFixture<LptAnsweredComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptAnsweredComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptAnsweredComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
