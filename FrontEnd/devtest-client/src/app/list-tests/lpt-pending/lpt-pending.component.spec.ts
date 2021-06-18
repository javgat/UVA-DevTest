import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptPendingComponent } from './lpt-pending.component';

describe('LptPendingComponent', () => {
  let component: LptPendingComponent;
  let fixture: ComponentFixture<LptPendingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptPendingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptPendingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
