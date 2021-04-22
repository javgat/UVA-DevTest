import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListTestsComponent } from './list-tests.component';

describe('ListTestsComponent', () => {
  let component: ListTestsComponent;
  let fixture: ComponentFixture<ListTestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListTestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListTestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
