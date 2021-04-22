import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TagService, Test, TestService, UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-tests',
  templateUrl: './tests.component.html',
  styleUrls: ['./tests.component.css']
})
export class TestsComponent extends LoggedInTeacherController implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private tS: TestService, private tagS: TagService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
  }

}
