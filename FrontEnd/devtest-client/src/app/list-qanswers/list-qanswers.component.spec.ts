import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListQAnswersComponent } from './list-qanswers.component';

describe('ListQAnswersComponent', () => {
  let component: ListQAnswersComponent;
  let fixture: ComponentFixture<ListQAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListQAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListQAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
