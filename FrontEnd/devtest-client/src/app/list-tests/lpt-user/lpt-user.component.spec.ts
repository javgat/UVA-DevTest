import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptUserComponent } from './lpt-user.component';

describe('LptUserComponent', () => {
  let component: LptUserComponent;
  let fixture: ComponentFixture<LptUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
