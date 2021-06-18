import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionTeamsComponent } from './question-teams.component';

describe('QuestionTeamsComponent', () => {
  let component: QuestionTeamsComponent;
  let fixture: ComponentFixture<QuestionTeamsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QuestionTeamsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionTeamsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
