import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptInvitedTeamComponent } from './lpt-invited-team.component';

describe('LptInvitedTeamComponent', () => {
  let component: LptInvitedTeamComponent;
  let fixture: ComponentFixture<LptInvitedTeamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptInvitedTeamComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptInvitedTeamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
