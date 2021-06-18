import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestUAnswersComponent } from './ptest-u-answers.component';

describe('PtestUAnswersComponent', () => {
  let component: PtestUAnswersComponent;
  let fixture: ComponentFixture<PtestUAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestUAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestUAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
