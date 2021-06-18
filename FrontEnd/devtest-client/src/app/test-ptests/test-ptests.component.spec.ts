import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestPTestsComponent } from './test-ptests.component';

describe('TestPTestsComponent', () => {
  let component: TestPTestsComponent;
  let fixture: ComponentFixture<TestPTestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestPTestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestPTestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
