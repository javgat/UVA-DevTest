import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfilePublicQuestionsComponent } from './profile-public-questions.component';

describe('ProfilePublicQuestionsComponent', () => {
  let component: ProfilePublicQuestionsComponent;
  let fixture: ComponentFixture<ProfilePublicQuestionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfilePublicQuestionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfilePublicQuestionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
