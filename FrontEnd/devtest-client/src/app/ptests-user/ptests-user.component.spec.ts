import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsUserComponent } from './ptests-user.component';

describe('PtestsUserComponent', () => {
  let component: PtestsUserComponent;
  let fixture: ComponentFixture<PtestsUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
