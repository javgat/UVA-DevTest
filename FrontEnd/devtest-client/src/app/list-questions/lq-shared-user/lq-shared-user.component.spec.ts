import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqSharedUserComponent } from './lq-shared-user.component';

describe('LqSharedUserComponent', () => {
  let component: LqSharedUserComponent;
  let fixture: ComponentFixture<LqSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
