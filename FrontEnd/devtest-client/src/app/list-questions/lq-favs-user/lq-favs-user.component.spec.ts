import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LqFavsUserComponent } from './lq-favs-user.component';

describe('LqFavsUserComponent', () => {
  let component: LqFavsUserComponent;
  let fixture: ComponentFixture<LqFavsUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LqFavsUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LqFavsUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
