import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnswersUserComponent } from './answers-user.component';

describe('AnswersUserComponent', () => {
  let component: AnswersUserComponent;
  let fixture: ComponentFixture<AnswersUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnswersUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnswersUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
