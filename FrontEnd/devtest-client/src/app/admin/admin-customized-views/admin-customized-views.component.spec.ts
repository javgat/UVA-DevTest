import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminCustomizedViewsComponent } from './admin-customized-views.component';

describe('AdminCustomizedViewsComponent', () => {
  let component: AdminCustomizedViewsComponent;
  let fixture: ComponentFixture<AdminCustomizedViewsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdminCustomizedViewsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminCustomizedViewsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
