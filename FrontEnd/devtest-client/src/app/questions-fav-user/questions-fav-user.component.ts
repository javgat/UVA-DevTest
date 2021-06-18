import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, QuestionService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-questions-fav-user',
  templateUrl: './questions-fav-user.component.html',
  styleUrls: ['./questions-fav-user.component.css']
})
export class QuestionsFavUserComponent extends LoggedInTeacherController implements OnInit {

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

  ngOnDestroy(): void {
    super.onDestroy()
    this.routeSub.unsubscribe()
    this.borrarMensaje()
  }

}