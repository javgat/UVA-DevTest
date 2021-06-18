import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsPendingUserComponent } from './ptests-pending-user.component';

describe('PtestsPendingUserComponent', () => {
  let component: PtestsPendingUserComponent;
  let fixture: ComponentFixture<PtestsPendingUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsPendingUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsPendingUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
