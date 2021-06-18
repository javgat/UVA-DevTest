import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtSharedUserComponent } from './lt-shared-user.component';

describe('LtSharedUserComponent', () => {
  let component: LtSharedUserComponent;
  let fixture: ComponentFixture<LtSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
