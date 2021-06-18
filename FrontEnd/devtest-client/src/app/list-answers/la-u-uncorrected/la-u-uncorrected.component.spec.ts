import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaUUncorrectedComponent } from './la-u-uncorrected.component';

describe('LaUUncorrectedComponent', () => {
  let component: LaUUncorrectedComponent;
  let fixture: ComponentFixture<LaUUncorrectedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LaUUncorrectedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaUUncorrectedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
