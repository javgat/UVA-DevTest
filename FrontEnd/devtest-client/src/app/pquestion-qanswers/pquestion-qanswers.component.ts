import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-pquestion-qanswers',
  templateUrl: './pquestion-qanswers.component.html',
  styleUrls: ['./pquestion-qanswers.component.css']
})
export class PquestionQAnswersComponent extends LoggedInTeacherController implements OnInit {

  routeSub: Subscription
  questionid: number
  testid: number
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.questionid = 0
    this.testid = 0
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.questionid = params['questionid']
      this.borrarMensaje()
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
  }

}
