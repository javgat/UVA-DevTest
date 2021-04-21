import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-questions-shared-user',
  templateUrl: './questions-shared-user.component.html',
  styleUrls: ['./questions-shared-user.component.css']
})
export class QuestionsSharedUserComponent extends LoggedInTeacherController implements OnInit {

  username: string
  routeSub: Subscription
  
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.username = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.routeSub.unsubscribe()
    this.borrarMensaje()
  }

}
