import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionsSharedUserComponent } from './questions-shared-user.component';

describe('QuestionsSharedUserComponent', () => {
  let component: QuestionsSharedUserComponent;
  let fixture: ComponentFixture<QuestionsSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QuestionsSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionsSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
