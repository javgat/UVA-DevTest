import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListAnswersComponent } from './list-answers.component';

describe('ListAnswersComponent', () => {
  let component: ListAnswersComponent;
  let fixture: ComponentFixture<ListAnswersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListAnswersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListAnswersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
