import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QanswerComponent } from './qanswer.component';

describe('QanswerComponent', () => {
  let component: QanswerComponent;
  let fixture: ComponentFixture<QanswerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QanswerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QanswerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
