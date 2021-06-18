import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LgiTeacherComponent } from './lgi-teacher.component';

describe('LgiTeacherComponent', () => {
  let component: LgiTeacherComponent;
  let fixture: ComponentFixture<LgiTeacherComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LgiTeacherComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LgiTeacherComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
