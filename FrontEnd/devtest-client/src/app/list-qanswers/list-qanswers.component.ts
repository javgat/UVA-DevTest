import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AnswerService, PublishedTestService, Question, QuestionAnswer, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-qanswers',
  templateUrl: './list-qanswers.component.html',
  styleUrls: ['./list-qanswers.component.css']
})
export class ListQAnswersComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  questionAnswers: QuestionAnswer[]
  testid: number
  answerid: number
  questions: Question[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, protected answerS: AnswerService, protected ptestS: PublishedTestService) {
    super(session, router, data, userS)
    this.questionAnswers = []
    this.questions = []
    this.testid = 0
    this.answerid = 0
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.answerid = params['answerid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getQAnswers(true)
        this.getQuestionsTest(true)
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
      this.getQAnswers(true)
      this.getQuestionsTest(true)
    }
  }

  getQAnswers(primera: boolean) {
    this.answerS.getQuestionAnswersFromAnswer(this.answerid).subscribe(
      resp => this.questionAnswers = resp,
      err => this.handleErrRelog(err, "obtener respuestas a preguntas del test", primera, this.getQAnswers, this)
    )
  }

  getQuestionsTest(primera: boolean) {
    this.ptestS.getQuestionsFromPublishedTests(this.testid).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas del test", primera, this.getQuestionsTest, this)
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined) {
    return tipoPrint(tipo, eleccionUnica)
  }

  getPregunta(idPreg: number | undefined): Question{
    return this.questions.filter(x => x.id == idPreg)[0]
  }

  calcValor(porcentaje: number | undefined, valorFinal: number | undefined): number{
    if(porcentaje == undefined || valorFinal == undefined) return 0
    return (porcentaje * valorFinal)/100
  }

}
