import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsSharedUserComponent } from './ptests-shared-user.component';

describe('PtestsSharedUserComponent', () => {
  let component: PtestsSharedUserComponent;
  let fixture: ComponentFixture<PtestsSharedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsSharedUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsSharedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
