import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtAdminComponent } from './lt-admin.component';

describe('LtAdminComponent', () => {
  let component: LtAdminComponent;
  let fixture: ComponentFixture<LtAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtAdminComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
