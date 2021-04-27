import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptests',
  templateUrl: './ptests.component.html',
  styleUrls: ['./ptests.component.css']
})
export class PtestsComponent extends LoggedInController implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }
  
  ngOnDestroy(): void {
    super.onDestroy()
  }
}
