import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileTeamsComponent } from './profile-teams.component';

describe('ProfileTeamsComponent', () => {
  let component: ProfileTeamsComponent;
  let fixture: ComponentFixture<ProfileTeamsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfileTeamsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfileTeamsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
