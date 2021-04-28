import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnswerQAnswersComponent } from './answer-qanswers.component';

describe('AnswerQAnswersComponent', () => {
  let component: AnswerQAnswersComponent;
  let fixture: ComponentFixture<AnswerQAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnswerQAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnswerQAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
