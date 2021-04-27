import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptests-shared-user',
  templateUrl: './ptests-shared-user.component.html',
  styleUrls: ['./ptests-shared-user.component.css']
})
export class PtestsSharedUserComponent extends LoggedInTeacherController implements OnInit {

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
