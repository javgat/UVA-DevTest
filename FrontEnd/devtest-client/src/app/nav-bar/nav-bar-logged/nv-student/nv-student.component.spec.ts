import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NvStudentComponent } from './nv-student.component';

describe('NvStudentComponent', () => {
  let component: NvStudentComponent;
  let fixture: ComponentFixture<NvStudentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NvStudentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NvStudentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
