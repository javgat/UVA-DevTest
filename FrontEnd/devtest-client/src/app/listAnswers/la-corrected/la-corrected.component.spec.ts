import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaCorrectedComponent } from './la-corrected.component';

describe('LaCorrectedComponent', () => {
  let component: LaCorrectedComponent;
  let fixture: ComponentFixture<LaCorrectedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LaCorrectedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaCorrectedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
