import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListPtestsComponent } from './list-ptests.component';

describe('ListPtestsComponent', () => {
  let component: ListPtestsComponent;
  let fixture: ComponentFixture<ListPtestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListPtestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListPtestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
