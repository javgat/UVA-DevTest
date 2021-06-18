import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NvTeacherAdminComponent } from './nv-teacher-admin.component';

describe('NvTeacherAdminComponent', () => {
  let component: NvTeacherAdminComponent;
  let fixture: ComponentFixture<NvTeacherAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NvTeacherAdminComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NvTeacherAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
