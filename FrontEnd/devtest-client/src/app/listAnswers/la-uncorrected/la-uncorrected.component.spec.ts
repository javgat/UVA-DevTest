import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaUncorrectedComponent } from './la-uncorrected.component';

describe('LaUncorrectedComponent', () => {
  let component: LaUncorrectedComponent;
  let fixture: ComponentFixture<LaUncorrectedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LaUncorrectedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaUncorrectedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
