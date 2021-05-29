import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AnswerService, PublishedTestService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQAnswersComponent } from '../list-qanswers.component';

@Component({
  selector: 'app-lqa-question',
  templateUrl: '../list-qanswers.component.html',
  styleUrls: ['../list-qanswers.component.css']
})
export class LqaQuestionComponent extends ListQAnswersComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    route: ActivatedRoute, answerS: AnswerService, ptestS: PublishedTestService) {
    super(session, router, data, userS, route, answerS, ptestS)
    this.mostrarAutor = true
    this.canSearchByUsername = true
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.ngOnDestroy()
  }

  getQAnswers(primera: boolean) {
    this.ptestS.getQuestionAnswersFromPublishedTestQuestion(this.testid, this.questionid, this.likeUsername).subscribe(
      resp => this.saveQAnswers(resp),
      err => this.handleErrRelog(err, "obtener respuestas a la pregunta de todos los usuarios", primera, this.getQAnswers, this)
    )
  }

  getQuestionsTest(primera: boolean) {
    this.ptestS.getQuestionsFromPublishedTests(this.testid).subscribe(
      resp => this.saveQuestionsTest(resp),
      err => this.handleErrRelog(err, "obtener preguntas del test", primera, this.getQuestionsTest, this)
    )
  }

}
