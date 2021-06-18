import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaUCorrectedComponent } from './la-u-corrected.component';

describe('LaUCorrectedComponent', () => {
  let component: LaUCorrectedComponent;
  let fixture: ComponentFixture<LaUCorrectedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LaUCorrectedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaUCorrectedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
