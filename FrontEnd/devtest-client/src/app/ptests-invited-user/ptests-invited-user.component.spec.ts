import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsInvitedUserComponent } from './ptests-invited-user.component';

describe('PtestsInvitedUserComponent', () => {
  let component: PtestsInvitedUserComponent;
  let fixture: ComponentFixture<PtestsInvitedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsInvitedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsInvitedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
