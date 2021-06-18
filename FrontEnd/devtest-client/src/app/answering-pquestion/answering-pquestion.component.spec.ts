import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnsweringPQuestionComponent } from './answering-pquestion.component';

describe('AnsweringPQuestionComponent', () => {
  let component: AnsweringPQuestionComponent;
  let fixture: ComponentFixture<AnsweringPQuestionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnsweringPQuestionComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnsweringPQuestionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
