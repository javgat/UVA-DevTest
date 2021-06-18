import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtProfilePublicComponent } from './lt-profile-public.component';

describe('LtProfilePublicComponent', () => {
  let component: LtProfilePublicComponent;
  let fixture: ComponentFixture<LtProfilePublicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtProfilePublicComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtProfilePublicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
