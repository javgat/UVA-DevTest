import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Question, QuestionService, UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-question-create',
  templateUrl: './question-create.component.html',
  styleUrls: ['./question-create.component.css']
})
export class QuestionCreateComponent extends LoggedInTeacherController implements OnInit {

  question: Question =  {
    title: "",
    question: "",
    estimatedTime: 0,
    autoCorrect: false,
    editable: true,
    username: "",
    tipoPregunta: Question.TipoPreguntaEnum.String,
    accesoPublicoNoPublicada: true
  }
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS : QuestionService) {
    super(session, router, data, userS)
  }

  doHasUserAction(){
    this.question.username = this.getSessionUser().getUsername()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
  }

  questionSubmit(){
    this.questionSend(true)
  }

  questionSend(primera: boolean){
    this.qS.postQuestion(this.getSessionUser().getUsername(), this.question).subscribe(
      resp=>{
        this.cambiarMensaje(new Mensaje("Pregunta creada con éxito", Tipo.SUCCESS, true))
        this.router.navigate(['/q', resp.id])
      },
      err => this.handleErrRelog(err, "crear una pregunta nueva", primera, this.questionSend, this)
    )
  }

  onSelectTipoString(){
    this.question.tipoPregunta = Question.TipoPreguntaEnum.String
    this.question.eleccionUnica = undefined
  }

  onSelectTipoRadio(){
    this.question.tipoPregunta = Question.TipoPreguntaEnum.Opciones
    this.question.eleccionUnica = true
  }

  onSelectTipoCheck(){
    this.question.tipoPregunta = Question.TipoPreguntaEnum.Opciones
    this.question.eleccionUnica = false
  }

  onSelectTipoCode(){
    this.question.tipoPregunta = Question.TipoPreguntaEnum.Codigo
    this.question.eleccionUnica = undefined
  }

}
