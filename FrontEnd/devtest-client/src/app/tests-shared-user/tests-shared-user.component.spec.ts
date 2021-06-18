import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestsSharedUserComponent } from './tests-shared-user.component';

describe('TestsSharedUserComponent', () => {
  let component: TestsSharedUserComponent;
  let fixture: ComponentFixture<TestsSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestsSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestsSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
