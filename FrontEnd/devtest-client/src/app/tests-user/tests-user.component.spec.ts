import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestsUserComponent } from './tests-user.component';

describe('TestsUserComponent', () => {
  let component: TestsUserComponent;
  let fixture: ComponentFixture<TestsUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestsUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestsUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
