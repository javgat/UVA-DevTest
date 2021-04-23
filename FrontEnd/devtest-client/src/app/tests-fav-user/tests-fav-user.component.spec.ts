import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestsFavUserComponent } from './tests-fav-user.component';

describe('TestsFavUserComponent', () => {
  let component: TestsFavUserComponent;
  let fixture: ComponentFixture<TestsFavUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestsFavUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestsFavUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
