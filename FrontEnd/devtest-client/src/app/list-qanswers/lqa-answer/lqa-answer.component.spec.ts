import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqaAnswerComponent } from './lqa-answer.component';

describe('LqaAnswerComponent', () => {
  let component: LqaAnswerComponent;
  let fixture: ComponentFixture<LqaAnswerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqaAnswerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqaAnswerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
