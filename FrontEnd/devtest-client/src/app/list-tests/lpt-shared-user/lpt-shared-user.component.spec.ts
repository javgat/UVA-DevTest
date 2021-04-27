import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptSharedUserComponent } from './lpt-shared-user.component';

describe('LptSharedUserComponent', () => {
  let component: LptSharedUserComponent;
  let fixture: ComponentFixture<LptSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
