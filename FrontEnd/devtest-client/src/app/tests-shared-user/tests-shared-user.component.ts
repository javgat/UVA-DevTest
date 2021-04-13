import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-tests-shared-user',
  templateUrl: './tests-shared-user.component.html',
  styleUrls: ['./tests-shared-user.component.css']
})
export class TestsSharedUserComponent extends LoggedInTeacherController implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
  }

}
