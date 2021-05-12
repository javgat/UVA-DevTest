import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtAdminTeamComponent } from './lt-admin-team.component';

describe('LtAdminTeamComponent', () => {
  let component: LtAdminTeamComponent;
  let fixture: ComponentFixture<LtAdminTeamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtAdminTeamComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtAdminTeamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
