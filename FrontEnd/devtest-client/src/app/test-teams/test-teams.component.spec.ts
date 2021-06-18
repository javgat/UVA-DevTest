import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestTeamsComponent } from './test-teams.component';

describe('TestTeamsComponent', () => {
  let component: TestTeamsComponent;
  let fixture: ComponentFixture<TestTeamsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestTeamsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestTeamsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
