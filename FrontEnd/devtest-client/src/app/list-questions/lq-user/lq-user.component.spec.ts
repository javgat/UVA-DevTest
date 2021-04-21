import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqUserComponent } from './lq-user.component';

describe('LqUserComponent', () => {
  let component: LqUserComponent;
  let fixture: ComponentFixture<LqUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
