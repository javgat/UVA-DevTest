import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AnswerService, PublishedTestService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQAnswersComponent } from '../list-qanswers.component';

@Component({
  selector: 'app-lqa-answer',
  templateUrl: '../list-qanswers.component.html',
  styleUrls: ['../list-qanswers.component.css']
})
export class LqaAnswerComponent extends ListQAnswersComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
      route: ActivatedRoute, answerS: AnswerService, ptestS: PublishedTestService) {
    super(session, router, data, userS, route, answerS, ptestS)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

  getQAnswers(primera: boolean) {
    this.answerS.getQuestionAnswersFromAnswer(this.answerid).subscribe(
      resp => this.saveQAnswers(resp),
      err => this.handleErrRelog(err, "obtener respuestas a preguntas del test", primera, this.getQAnswers, this)
    )
  }

  getQuestionsTest(primera: boolean) {
    this.ptestS.getQuestionsFromPublishedTests(this.testid).subscribe(
      resp => this.saveQuestionsTest(resp),
      err => this.handleErrRelog(err, "obtener preguntas del test", primera, this.getQuestionsTest, this)
    )
  }

}
