import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqAdminTeamComponent } from './lq-admin-team.component';

describe('LqAdminTeamComponent', () => {
  let component: LqAdminTeamComponent;
  let fixture: ComponentFixture<LqAdminTeamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqAdminTeamComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqAdminTeamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
