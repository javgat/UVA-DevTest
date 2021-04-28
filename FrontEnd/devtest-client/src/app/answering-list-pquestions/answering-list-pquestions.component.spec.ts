import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnsweringListPQuestionsComponent } from './answering-list-pquestions.component';

describe('AnsweringListPQuestionsComponent', () => {
  let component: AnsweringListPQuestionsComponent;
  let fixture: ComponentFixture<AnsweringListPQuestionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnsweringListPQuestionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnsweringListPQuestionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
