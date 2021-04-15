import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-tests-shared-user',
  templateUrl: './tests-shared-user.component.html',
  styleUrls: ['./tests-shared-user.component.css']
})
export class TestsSharedUserComponent extends LoggedInTeacherController implements OnInit {

  routeSub: Subscription
  username: string
  tests: Test[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.tests=[]
    this.username=""
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
      this.getTestsSharedUser(true)
    });
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
    this.borrarMensaje()
  }

  getTestsSharedUser(primera: boolean){
    this.userS.getSharedTestsFromUser(this.username).subscribe(
      resp => this.tests = resp,
      err => this.handleErrRelog(err, "obtener tests compartidos con el usuario", primera, this.getTestsSharedUser, this)
    )
  }
}
