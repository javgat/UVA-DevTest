import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfilePublicTestsComponent } from './profile-public-tests.component';

describe('ProfilePublicTestsComponent', () => {
  let component: ProfilePublicTestsComponent;
  let fixture: ComponentFixture<ProfilePublicTestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfilePublicTestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfilePublicTestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
