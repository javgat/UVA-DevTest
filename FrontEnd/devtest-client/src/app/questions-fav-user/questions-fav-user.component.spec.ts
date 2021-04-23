import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionsFavUserComponent } from './questions-fav-user.component';

describe('QuestionsFavUserComponent', () => {
  let component: QuestionsFavUserComponent;
  let fixture: ComponentFixture<QuestionsFavUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QuestionsFavUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionsFavUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
