import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestAnswersComponent } from './ptest-answers.component';

describe('PtestAnswersComponent', () => {
  let component: PtestAnswersComponent;
  let fixture: ComponentFixture<PtestAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
