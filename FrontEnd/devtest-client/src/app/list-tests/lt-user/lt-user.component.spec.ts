import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtUserComponent } from './lt-user.component';

describe('LtUserComponent', () => {
  let component: LtUserComponent;
  let fixture: ComponentFixture<LtUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
