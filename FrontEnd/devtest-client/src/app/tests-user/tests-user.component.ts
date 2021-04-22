import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Test, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-tests-user',
  templateUrl: './tests-user.component.html',
  styleUrls: ['./tests-user.component.css']
})
export class TestsUserComponent extends LoggedInTeacherController implements OnInit {

  routeSub: Subscription
  username: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.username=""
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
    });
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
    this.borrarMensaje()
  }

}
