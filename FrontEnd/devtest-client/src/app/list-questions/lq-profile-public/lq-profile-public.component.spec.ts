import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqProfilePublicComponent } from './lq-profile-public.component';

describe('LqProfilePublicComponent', () => {
  let component: LqProfilePublicComponent;
  let fixture: ComponentFixture<LqProfilePublicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqProfilePublicComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqProfilePublicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
