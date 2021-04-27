import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfilePublicPublishedTestsComponent } from './profile-public-published-tests.component';

describe('ProfilePublicPublishedTestsComponent', () => {
  let component: ProfilePublicPublishedTestsComponent;
  let fixture: ComponentFixture<ProfilePublicPublishedTestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfilePublicPublishedTestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfilePublicPublishedTestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
