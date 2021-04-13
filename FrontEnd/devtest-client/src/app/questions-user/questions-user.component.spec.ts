import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionsUserComponent } from './questions-user.component';

describe('QuestionsUserComponent', () => {
  let component: QuestionsUserComponent;
  let fixture: ComponentFixture<QuestionsUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QuestionsUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionsUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
