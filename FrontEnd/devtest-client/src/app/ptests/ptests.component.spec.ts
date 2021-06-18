import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PtestsComponent } from './ptests.component';

describe('PtestsComponent', () => {
  let component: PtestsComponent;
  let fixture: ComponentFixture<PtestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PtestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PtestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
