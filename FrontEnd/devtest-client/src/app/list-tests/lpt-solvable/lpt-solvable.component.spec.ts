import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LptSolvableComponent } from './lpt-solvable.component';

describe('LptSolvableComponent', () => {
  let component: LptSolvableComponent;
  let fixture: ComponentFixture<LptSolvableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LptSolvableComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LptSolvableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
