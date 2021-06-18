import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqaQuestionComponent } from './lqa-question.component';

describe('LqaQuestionComponent', () => {
  let component: LqaQuestionComponent;
  let fixture: ComponentFixture<LqaQuestionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqaQuestionComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqaQuestionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
