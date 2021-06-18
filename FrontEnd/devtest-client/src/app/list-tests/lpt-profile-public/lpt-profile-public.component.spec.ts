import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptProfilePublicComponent } from './lpt-profile-public.component';

describe('LptProfilePublicComponent', () => {
  let component: LptProfilePublicComponent;
  let fixture: ComponentFixture<LptProfilePublicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptProfilePublicComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptProfilePublicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
