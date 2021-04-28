import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, PublishedTestService, Question, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, Pregunta, Tipo, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-answering-pquestion',
  templateUrl: './answering-pquestion.component.html',
  styleUrls: ['./answering-pquestion.component.css']
})
export class AnsweringPQuestionComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  testid: number
  preguntaid: number
  openAnswer?: Answer
  pregunta: Question

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private ptestS: PublishedTestService, private answerS: AnswerService) {
    super(session, router, data, userS);
    this.testid = 0
    this.preguntaid = 0
    this.pregunta = new Pregunta()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.preguntaid = params['questionid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getOpenAnswer(true)
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
    if (this.testid != undefined && this.testid != 0) {
      this.getOpenAnswer(true)
    }
  }

  getOpenAnswer(primera: boolean) {
    this.userS.getOpenAnswersFromUserTest(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        if (resp.length == 0) {
          this.cambiarMensaje(new Mensaje("No existe una respuesta iniciada para el test por el usuario", Tipo.ERROR, true))
        } else {
          this.openAnswer = resp[0]
          this.getPQuestionFromPTest(true)
        }
      },
      err => this.handleErrRelog(err, "obtener respuesta no finalizada de usuario en test", primera, this.getOpenAnswer, this)
    )
  }

  getPQuestionFromPTest(primera: boolean) {
    this.ptestS.getQuestionFromPublishedTests(this.testid, this.preguntaid).subscribe(
      resp => {
        this.pregunta = resp
        this.getQuestionAnswersQuestion(true)
      },
      err => {
        this.handleErrRelog(err, "obtener pregunta publicada de test publicado", primera, this.getPQuestionFromPTest, this)
      }
    )
  }

  getQuestionAnswersQuestion(primera: boolean) {
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.getQuestionAnswerFromAnswer(this.openAnswer.id, this.preguntaid).subscribe(
      resp => {
        this.pregunta.isRespondida = true
      },
      err => {
        if (err.status == 410) {
          this.pregunta.isRespondida = false
        } else {
          this.handleErrRelog(err, "obtener respuestas de una pregunta del test realizandose", primera, this.getQuestionAnswersQuestion, this)
        }
      }
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string {
    return tipoPrint(tipo, eleccionUnica)
  }


}
