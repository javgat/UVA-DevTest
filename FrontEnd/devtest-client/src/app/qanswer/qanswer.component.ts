import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AnswerService, Option, PublishedTestService, Question, QuestionAnswer, QuestionService, Review, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Pregunta, RespuestaPregunta, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-qanswer',
  templateUrl: './qanswer.component.html',
  styleUrls: ['./qanswer.component.css']
})
export class QanswerComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  questionid: number
  testid: number
  answerid: number
  qa: RespuestaPregunta
  question: Pregunta
  options: Option[]
  showCorregir: boolean
  editPuntuacion: number
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, private answerS: AnswerService, private ptestS: PublishedTestService, private qS: QuestionService) {
    super(session, router, data, userS)
    this.questionid = 0
    this.testid = 0
    this.answerid = 0
    this.options = []
    this.question = new Pregunta()
    this.editPuntuacion = 0
    this.qa = new RespuestaPregunta()
    this.showCorregir = false
    this.routeSub = this.route.params.subscribe(params => {
      this.questionid = params['questionid']
      this.testid = params['testid']
      this.answerid = params['answerid']
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getQAnswer(true)
        this.getPregunta(true)
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
    if (this.isDefinido(this.testid) && this.isDefinido(this.answerid) && this.isDefinido(this.questionid)) {
      this.getQAnswer(true)
      this.getPregunta(true)
    }
  }

  isDefinido(num: number | undefined): boolean {
    return (num != undefined && num != 0)
  }

  getQAnswer(primera: boolean) {
    this.answerS.getQuestionAnswerFromAnswer(this.answerid, this.questionid).subscribe(
      resp => {
        this.qa = new RespuestaPregunta(resp)
        this.editPuntuacion = this.qa.puntuacion
      },
      err => this.handleErrRelog(err, "obtener respuesta de una pregunta", primera, this.getQAnswer, this)
    )
  }

  getPregunta(primera: boolean) {
    this.ptestS.getQuestionFromPublishedTests(this.testid, this.questionid).subscribe(
      resp => {
        this.question = Pregunta.constructorFromQuestion(resp)
        if (this.question.tipoPregunta == Question.TipoPreguntaEnum.Opciones) {
          this.getOptions(true)
        }
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
    )
  }

  getOptions(primera: boolean) {
    this.qS.getOptionsFromQuestion(this.questionid).subscribe(
      resp => {
        this.options = resp
      },
      err => this.handleErrRelog(err, "obtener opciones de respuesta de pregunta", primera, this.getOptions, this)
    )
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined): string {
    return tipoPrint(tipo, eleUni)
  }

  opcionSeleccionada(opIn: number | undefined): boolean {
    if (opIn == undefined) return false
    return this.qa.indicesOpciones.filter(x => x == opIn).length > 0
  }

  radioEsCorrecta(opc: Option): boolean {
    return (opc.correcta || false) && this.opcionSeleccionada(opc.indice)
  }

  radioEsIncorrecta(opc: Option): boolean {
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice))
  }

  radioShowTick(opc: Option): boolean {
    return (opc.correcta || false)
  }

  radioShowCross(opc: Option): boolean {
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice))
  }

  checkEsCorrecta(opc: Option): boolean {
    return (opc.correcta || false) && this.opcionSeleccionada(opc.indice)
  }

  checkEsIncorrecta(opc: Option): boolean {
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice)) || (opc.correcta==true && !this.opcionSeleccionada(opc.indice))
  }

  checkShowTick(opc: Option): boolean {
    return (opc.correcta || false)
  }

  checkShowCross(opc: Option): boolean {
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice))
  }

  isRadioQuestion(): boolean{
    return this.question.eleccionUnica==true
  }

  isCheckQuestion(): boolean{
    return !this.isRadioQuestion()
  }
  
  clickShowCorregir(){
    this.showCorregir = true
  }

  clickNotShowCorregir(){
    this.showCorregir = false
  }

  submitUpdateCorrection(){
    this.showCorregir = false
    this.updateCorrection(true)
  }

  updateCorrection(primera: boolean){
    let review: Review
    review = {
      puntuacion : this.editPuntuacion
    }
    this.answerS.putReview(this.answerid, this.questionid, review).subscribe(
      resp => this.getQAnswer(true),
      err => this.handleErrRelog(err, "actualizar correccion", primera, this.updateCorrection, this)
    )
  }

  setAsNotCorregidaClick(){
    this.setAsNotCorregida(true)
  }

  setAsNotCorregida(primera: boolean){
    this.answerS.deleteReview(this.answerid, this.questionid).subscribe(
      resp => this.getQAnswer(true),
      err => this.handleErrRelog(err, "borrar correccion", primera, this.setAsNotCorregida, this)
    )
  }

  calcValor(porcentaje: number | undefined, valorFinal: number | undefined): number{
    if(porcentaje == undefined || valorFinal == undefined) return 0
    return (porcentaje * valorFinal)/100
  }

}
