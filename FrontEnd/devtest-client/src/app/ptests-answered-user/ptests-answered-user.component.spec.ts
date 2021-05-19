import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsAnsweredUserComponent } from './ptests-answered-user.component';

describe('PtestsAnsweredUserComponent', () => {
  let component: PtestsAnsweredUserComponent;
  let fixture: ComponentFixture<PtestsAnsweredUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsAnsweredUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsAnsweredUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
