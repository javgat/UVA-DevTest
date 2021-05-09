import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NavBarNotLoggedComponent } from './nav-bar-not-logged.component';

describe('NavBarNotLoggedComponent', () => {
  let component: NavBarNotLoggedComponent;
  let fixture: ComponentFixture<NavBarNotLoggedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NavBarNotLoggedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavBarNotLoggedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
