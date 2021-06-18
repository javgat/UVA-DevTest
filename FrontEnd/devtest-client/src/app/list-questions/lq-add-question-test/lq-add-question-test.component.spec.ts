import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqAddQuestionTestComponent } from './lq-add-question-test.component';

describe('LqAddQuestionTestComponent', () => {
  let component: LqAddQuestionTestComponent;
  let fixture: ComponentFixture<LqAddQuestionTestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqAddQuestionTestComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqAddQuestionTestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
