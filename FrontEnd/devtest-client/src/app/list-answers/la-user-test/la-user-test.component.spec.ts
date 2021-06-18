import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaUserTestComponent } from './la-user-test.component';

describe('LaUserTestComponent', () => {
  let component: LaUserTestComponent;
  let fixture: ComponentFixture<LaUserTestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LaUserTestComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaUserTestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
