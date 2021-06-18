import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestInvitesComponent } from './ptest-invites.component';

describe('PtestInvitesComponent', () => {
  let component: PtestInvitesComponent;
  let fixture: ComponentFixture<PtestInvitesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestInvitesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestInvitesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
