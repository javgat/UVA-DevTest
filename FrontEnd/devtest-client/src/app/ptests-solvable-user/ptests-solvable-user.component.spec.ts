import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsSolvableUserComponent } from './ptests-solvable-user.component';

describe('PtestsSolvableUserComponent', () => {
  let component: PtestsSolvableUserComponent;
  let fixture: ComponentFixture<PtestsSolvableUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsSolvableUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsSolvableUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
