import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LgiStudentComponent } from './lgi-student.component';

describe('LgiStudentComponent', () => {
  let component: LgiStudentComponent;
  let fixture: ComponentFixture<LgiStudentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LgiStudentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LgiStudentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
