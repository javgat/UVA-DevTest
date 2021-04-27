import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptInvitedUserComponent } from './lpt-invited-user.component';

describe('LptInvitedUserComponent', () => {
  let component: LptInvitedUserComponent;
  let fixture: ComponentFixture<LptInvitedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptInvitedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptInvitedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
