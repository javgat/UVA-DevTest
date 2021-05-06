import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, PublishedTestService, Question, QuestionAnswer, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Respuesta, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-answer-qanswers',
  templateUrl: './answer-qanswers.component.html',
  styleUrls: ['./answer-qanswers.component.css']
})
export class AnswerQAnswersComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  testid: number
  answerid: number
  answer: Answer
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, protected answerS: AnswerService, protected ptestS: PublishedTestService) {
    super(session, router, data, userS)
    this.testid = 0
    this.answerid = 0
    this.answer = new Respuesta()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.answerid = params['answerid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getAnswer(true)
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
  }

  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0 && this.answerid != undefined && this.answerid != 0) {
      this.getAnswer(true)
    }
  }

  getAnswer(primera: boolean) {
    this.answerS.getAnswer(this.answerid).subscribe(
      resp => {
        this.answer = new Respuesta(resp)
      },
      err => this.handleErrRelog(err, "obtener respuesta a examen", primera, this.getAnswer, this)
    )
  }

}
