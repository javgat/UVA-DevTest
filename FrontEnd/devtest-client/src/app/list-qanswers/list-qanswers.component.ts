import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, PublishedTestService, Question, QuestionAnswer, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { bgcolorQAnswerPuntuacion, Mensaje, Respuesta, Tipo, tipoPrint } from '../shared/app.model';
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
  questionid: number
  questions: Question[]
  mostrarAutor: boolean
  mensajeListaVacia: string
  canSearchByUsername: boolean
  likeUsername?: string
  editLikeUsername: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, protected answerS: AnswerService, protected ptestS: PublishedTestService) {
    super(session, router, data, userS)
    this.questionAnswers = []
    this.questions = []
    this.testid = 0
    this.answerid = 0
    this.questionid = 0
    this.mostrarAutor = false
    this.canSearchByUsername = false
    this.likeUsername = undefined
    this.editLikeUsername = ""
    this.mensajeListaVacia = "¡Vaya! Parece que no hay ninguna Respuesta a una Pregunta para mostrar."
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.answerid = params['answerid']
      this.questionid = params['questionid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getWantedQAnswers()
        this.getWantedQuestionsTest()
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
    if (this.isDefinido(this.testid) && (this.isDefinido(this.answerid) || this.isDefinido(this.questionid))) {
      this.getWantedQAnswers()
      this.getWantedQuestionsTest()
    }
  }

  isDefinido(num: number | undefined): boolean{
    return (num!=undefined && num!=0)
  }

  saveQAnswers(resp: QuestionAnswer[]){
    this.borrarMensaje()
    this.questionAnswers = resp
  }

  saveQuestionsTest(resp: Question[]){
    this.borrarMensaje()
    this.questions = resp
  }

  getWantedQAnswers(): void{
    this.cambiarMensaje(new Mensaje("Descargando respuestas del usuario a las preguntas del test... ", Tipo.DOWNLOADING, true))
    this.getQAnswers(true)
  }

  getWantedQuestionsTest(): void{
    this.cambiarMensaje(new Mensaje("Descargando preguntas del test... ", Tipo.DOWNLOADING, true))
    this.getQuestionsTest(true)
  }

  getQAnswers(primera: boolean): void{

  }

  getQuestionsTest(primera: boolean): void{
    
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

  bgcolorQAnswerPuntuacion(punt: number){
    return bgcolorQAnswerPuntuacion(punt)
  }

  canSearchUsername(): boolean{
    return this.canSearchByUsername
  }

  clickSearchUsername(){
    this.likeUsername = this.editLikeUsername
    this.getWantedQAnswers()
  }

  clickBorrarUsername(){
    this.likeUsername = undefined
    this.editLikeUsername = ""
    this.getWantedQAnswers()
  }


}
