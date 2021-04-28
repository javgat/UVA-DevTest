import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PquestionQAnswersComponent } from './pquestion-qanswers.component';

describe('PquestionQAnswersComponent', () => {
  let component: PquestionQAnswersComponent;
  let fixture: ComponentFixture<PquestionQAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PquestionQAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PquestionQAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
