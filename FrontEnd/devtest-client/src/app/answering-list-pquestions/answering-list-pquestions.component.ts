import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, PublishedTestService, Question, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen, Mensaje, Tipo, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-answering-list-pquestions',
  templateUrl: './answering-list-pquestions.component.html',
  styleUrls: ['./answering-list-pquestions.component.css']
})
export class AnsweringListPQuestionsComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  testid: number
  openAnswer?: Answer
  test: Test
  preguntas: Question[]
  mostrarDescripcion: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private ptestS: PublishedTestService, private answerS: AnswerService) {
    super(session, router, data, userS);
    this.testid = 0
    this.preguntas = []
    this.mostrarDescripcion = false
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
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
          this.getPTest(true)
          this.getPreguntasAnswer(true)
        }
      },
      err => this.handleErrRelog(err, "obtener respuesta no finalizada de usuario en test", primera, this.getOpenAnswer, this)
    )
  }

  getPTest(primera: boolean) {
    this.userS.getSolvableTestFromUser(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
      },
      err => {
        this.handleErrRelog(err, "obtener test publicado", primera, this.getPTest, this)
      }
    )
  }

  getPreguntasAnswer(primera: boolean) {
    if(this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.getQuestionsFromAnswer(this.openAnswer.id).subscribe(
      resp => {
        this.preguntas = resp
      },
      err => {
        this.handleErrRelog(err, "obtener preguntas de test publicado respondiendo", primera, this.getPreguntasAnswer, this)
      }
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string{
    return tipoPrint(tipo, eleccionUnica)
  }

  finishAnswerClick(){
    this.finishAnswer(true)
  }

  finishAnswer(primera: boolean){
    if(this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.finishAnswer(this.openAnswer.id).subscribe(
      resp => {
        console.log("respuesta enviada con éxito")
        this.router.navigate(['/pt', this.testid])
        this.cambiarMensaje(new Mensaje("respuesta enviada con éxito", Tipo.SUCCESS, true))
      },
      err => this.handleErrRelog(err, "marcar como finalizada una respuesta", primera, this.finishAnswer, this)
    )
  }

  getNumeroPreguntasTotales(): number{
    return this.preguntas.length
  }

  getNumeroPreguntasRespondidas(): number{
    return this.preguntas.filter((val, i, quests) => val.isRespondida==true).length
  }

  switchMostrarDescripcion(){
    this.mostrarDescripcion = !this.mostrarDescripcion
  }

}
