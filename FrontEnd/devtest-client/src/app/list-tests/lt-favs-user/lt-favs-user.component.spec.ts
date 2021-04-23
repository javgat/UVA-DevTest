import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LtFavsUserComponent } from './lt-favs-user.component';

describe('LtFavsUserComponent', () => {
  let component: LtFavsUserComponent;
  let fixture: ComponentFixture<LtFavsUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LtFavsUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LtFavsUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
