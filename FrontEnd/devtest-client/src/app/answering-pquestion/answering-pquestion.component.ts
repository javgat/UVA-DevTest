import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, Option, Prueba, PublishedTestService, Question, QuestionAnswer, Testing, UserService } from '@javgat/devtest-api';
import { CodeModel } from '@ngstack/code-editor';
import { CountdownEvent } from 'ngx-countdown';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen, Mensaje, Pregunta, ResultadoPruebas, Tipo, tipoPrint } from '../shared/app.model';
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
  pregunta: Pregunta
  questionAnswer: QuestionAnswer
  newRespuesta: string
  options: Option[]
  modificandoRespuesta: boolean
  test: Examen
  timeOver: boolean
  pruebas: Prueba[]
  isMostrandoPruebas: boolean
  resPruebas: ResultadoPruebas

  theme = 'vs-dark';

  codeModel: CodeModel

  codeOptions = {
    contextmenu: true,
    minimap: {
      enabled: true,
    },
  };
  TIEMPO_RECARGA_ESTADO_COMPILACION: number = 4000 //ms

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private ptestS: PublishedTestService, private answerS: AnswerService) {
    super(session, router, data, userS);
    this.modificandoRespuesta = false
    this.isMostrandoPruebas = false
    this.testid = 0
    this.preguntaid = 0
    this.options = []
    this.pregunta = new Pregunta()
    this.test = new Examen()
    this.timeOver = false
    this.resPruebas = new ResultadoPruebas()
    this.pruebas = []
    this.questionAnswer = {
      idPregunta: 0,
      idRespuesta: 0,
      puntuacion: 0,
      corregida: false,
      respuesta: "",
      estado: QuestionAnswer.EstadoEnum.NoProbado
    }
    this.newRespuesta = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.preguntaid = parseInt(params['questionid'])
      this.questionAnswer.idPregunta = this.preguntaid
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getOpenAnswer(true)
      }
    });
    this.codeModel = {
      language: 'cpp',
      uri: 'main.cpp',
      value: this.questionAnswer.respuesta || "",
    }
    this.recargarEditorCodigo()
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

  resetViewData(): void{
    this.modificandoRespuesta = false
    this.testid = 0
    this.preguntaid = 0
    this.options = []
    this.pregunta = new Pregunta()
    this.questionAnswer = {
      idPregunta: 0,
      idRespuesta: 0,
      puntuacion: 0,
      corregida: false,
      respuesta: ""
    }
    this.newRespuesta = ""
  }

  getOpenAnswer(primera: boolean) {
    this.userS.getOpenAnswersFromUserTest(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        if (resp.length == 0) {
          this.cambiarMensaje(new Mensaje("No existe una respuesta iniciada para el test por el usuario", Tipo.ERROR, true))
          this.router.navigate(['/pt', this.testid])
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
        this.pregunta = Pregunta.constructorFromQuestion(resp)
        this.getQuestionAnswersQuestion(true)
        if (this.pregunta.tipoPregunta == Question.TipoPreguntaEnum.Opciones) {
          this.getOpciones(true)
        }else if(this.pregunta.tipoPregunta == Question.TipoPreguntaEnum.Codigo){
          this.getPruebasVisibles(true)
        }
        this.getPTest(true)
      },
      err => {
        this.handleErrRelog(err, "obtener pregunta publicada de test publicado", primera, this.getPQuestionFromPTest, this)
      }
    )
  }

  getPTest(primera: boolean) {
    this.userS.getSolvableTestFromUser(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.updateIsTimeOver()
      },
      err => {
        this.handleErrRelog(err, "obtener test publicado", primera, this.getPTest, this)
      }
    )
  }

  getOpciones(primera: boolean) {
    this.ptestS.getOptionsFromPublishedQuestion(this.testid, this.preguntaid).subscribe(
      resp => this.options = resp,
      err => this.handleErrRelog(err, "obtener opciones de respuesta multiple", primera, this.getOpciones, this)
    )
  }

  getQuestionAnswersQuestion(primera: boolean) {
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.getQuestionAnswerFromAnswer(this.openAnswer.id, this.preguntaid).subscribe(
      resp => {
        this.pregunta.isRespondida = true
        this.questionAnswer = resp
        this.newRespuesta = this.questionAnswer.respuesta || ""
        this.recargarEditorCodigo()
        this.questionAnswer.idPregunta = this.preguntaid
        if (this.questionAnswer.indicesOpciones == undefined) {
          this.questionAnswer.indicesOpciones = []
        }
        console.log(this.questionAnswer.estado)
        if(this.questionAnswer.estado==QuestionAnswer.EstadoEnum.Probado){
          this.getResultadoPruebas(true)
        }else  if(this.questionAnswer.estado == QuestionAnswer.EstadoEnum.Ejecutando){
          setTimeout(() => {this.getQuestionAnswersQuestion(true)}, this.TIEMPO_RECARGA_ESTADO_COMPILACION)
        }
        this.getPruebasVisibles(true)
      },
      err => {
        if (err.status == 410) {
          this.pregunta.isRespondida = false
          this.questionAnswer.respuesta = ""
          this.newRespuesta = this.questionAnswer.respuesta || ""
          this.recargarEditorCodigo()
        } else {
          this.handleErrRelog(err, "obtener respuestas de una pregunta del test realizandose", primera, this.getQuestionAnswersQuestion, this)
        }
      }
    )
  }

  getResultadoPruebas(primera: boolean){
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.getPreTesting(this.openAnswer.id, this.preguntaid).subscribe(
      resp=>{
        this.resPruebas = resp
      },
      err => this.handleErrRelog(err, "obtener resultado de pretesting de pruebas", primera, this.getResultadoPruebas, this)
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string {
    return tipoPrint(tipo, eleccionUnica)
  }

  sendTextRespuestaClick() {
    this.questionAnswer.respuesta = this.newRespuesta
    this.newRespuesta = ""
    this.sendRespuesta()
  }

  sendRespuesta() {
    this.modificandoRespuesta = false
    if (this.pregunta.isRespondida) {
      this.putRespuesta(true)
    } else {
      this.postRespuesta(true)
    }
  }

  postRespuesta(primera: boolean) {
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.postQuestionAnswer(this.openAnswer.id, this.questionAnswer).subscribe(
      resp => {
        //this.cambiarMensaje(new Mensaje("Respuesta actualizada con éxito", Tipo.SUCCESS, true))
        this.getOpenAnswer(true)
        this.executePreTestingIfCode(true)
      },
      err => {
        this.handleErrRelog(err, "publicar nueva respuesta a una pregunta de test publicado", primera, this.postRespuesta, this)
      }
    )
  }

  putRespuesta(primera: boolean) {
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.putQuestionAnswerFromAnswer(this.openAnswer.id, this.preguntaid, this.questionAnswer).subscribe(
      resp => {
        this.cambiarMensaje(new Mensaje("Respuesta actualizada con éxito", Tipo.SUCCESS, true))
        this.getOpenAnswer(true)
        this.executePreTestingIfCode(true)
      },
      err => {
        this.handleErrRelog(err, "modificar una respuesta a una pregunta de test publicado", primera, this.putRespuesta, this)
      }
    )
  }

  executePreTestingIfCode(primera: boolean){
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    if(this.pregunta.tipoPregunta==Question.TipoPreguntaEnum.Codigo){
      this.answerS.createPreTesting(this.openAnswer.id, this.pregunta.id).subscribe(
        resp => {},
        err => this.handleErrRelog(err, "pedir la compilacion en backend de pretesting", primera, this.executePreTestingIfCode, this)
      )
    }
  }


  isChecked(indiceCheck: number | undefined): boolean {
    if (indiceCheck == undefined) return false
    return this.questionAnswer.indicesOpciones?.includes(indiceCheck) || false
  }

  pressedCheckbox(indicePressed: number | undefined) {
    if (indicePressed == undefined) return
    if (this.isChecked(indicePressed)) {
      const index = this.questionAnswer.indicesOpciones?.indexOf(indicePressed, 0);
      if (index != undefined && index > -1) {
        this.questionAnswer.indicesOpciones?.splice(index, 1);
      }
    } else {
      if(this.questionAnswer.indicesOpciones == undefined)
        this.questionAnswer.indicesOpciones =[]
      this.questionAnswer.indicesOpciones.push(indicePressed)
    }
  }

  pressedRadio(indicePressed: number | undefined) {
    if (indicePressed == undefined) return
    this.questionAnswer.indicesOpciones = [indicePressed]
  }

  sendRespuestaTipoTest(){
    this.sendRespuesta()
  }

  sendRespuestaTipoCode(){
    this.questionAnswer.respuesta = this.newRespuesta
    this.sendRespuesta()
  }

  borrarRespuestaClick(){
    this.newRespuesta = ""
    this.borrarRespuesta(true)
  }

  borrarRespuesta(primera: boolean){
    if (this.openAnswer == undefined || this.openAnswer.id == undefined) return
    this.answerS.deleteQuestionAnswerFromAnswer(this.openAnswer?.id, this.preguntaid).subscribe(
      resp => {
        //this.cambiarMensaje(new Mensaje("Respuesta borrada con éxito", Tipo.SUCCESS, true))
        this.getOpenAnswer(true)
      },
      err => {
        this.handleErrRelog(err, "borrar respuesta a una pregunta", primera, this.borrarRespuesta, this)
      }
    )
  }

  modificarRespuestaClick(){
    this.modificandoRespuesta = true
    this.newRespuesta = this.questionAnswer.respuesta || ""
  }

  hasPrevQuestion(): boolean{
    return this.pregunta.hasPrevious()
  }

  hasNextQuestion(): boolean{
    return this.pregunta.hasNext()
  }

  goPrevQuestion(): void{
    let testid = this.testid
    let prevId = this.pregunta.prevId
    this.resetViewData()
    this.router.navigate(['/pt', testid, 'answering', 'pq', prevId])
  }

  goNextQuestion(): void{
    let testid = this.testid
    let nextId = this.pregunta.nextId
    this.resetViewData()
    this.router.navigate(['/pt', testid, 'answering', 'pq', nextId])
  }

  showWarningExitQuestion(): boolean{
    return (!this.pregunta.isRespondida) || this.modificandoRespuesta
  }

  updateIsTimeOver(){
    if(this.getLeftTime()>=0){
      this.timeOver = false
    }else{
      this.timeOver = true
    }
  }
  
  leftTime?: number

  getLeftTime(): number{
    if(this.leftTime !=undefined)
      return this.leftTime
    if(this.openAnswer==undefined || this.openAnswer.startTime == undefined) return 0
    var date = new Date(this.openAnswer.startTime)
    var now = new Date()
    var ahoraSecs = now.getTime()/1000
    this.leftTime = ((date.getTime()/1000)+(this.test.maxMinutes*60))-ahoraSecs
    return this.leftTime
  }

  timeIsOver(event: CountdownEvent){
    if(event.left == 0)
      this.timeOver = true
  }

  isTimeOver(): boolean{
    if(this.openAnswer==undefined || this.openAnswer.startTime == undefined) return false
    return this.timeOver

  }

  recargarEditorCodigo(){
    this.codeModel = {
      language: 'cpp',
      uri: 'main.cpp',
      value: this.questionAnswer.respuesta || "",
    }
  }

  onCodeChanged(value: any) {
    if(value!=this.newRespuesta){
      this.modificandoRespuesta = true
      this.newRespuesta = value
    }
  }

  isPreguntaCodigo(): boolean{
    return this.pregunta.tipoPregunta=='codigo'
  }

  getPruebasVisibles(primera: boolean){
    if(this.openAnswer==undefined || this.openAnswer.id==undefined) return
    this.answerS.getVisiblePublishedPruebasFromQuestionTest(this.openAnswer.id, this.questionAnswer.idPregunta).subscribe(
      resp => {
        this.pruebas = resp
        if(this.pruebas.length==0){
          this.geUndonetVisiblePruebas(true)
        }
      },
      err => this.handleErrRelog(err, "obtener las pruebas visibles de una pregunta de codigo", primera, this.getPruebasVisibles, this)
    )
  }

  geUndonetVisiblePruebas(primera: boolean){
    if (this.testid == undefined) return
    this.ptestS.getVisiblePruebasFromQuestionTest(this.testid, this.questionAnswer.idPregunta).subscribe(
      resp => {
        this.pruebas = resp
      },
      err => this.handleErrRelog(err, "obtener pruebas visibles de respuesta de pregunta", primera, this.geUndonetVisiblePruebas, this)
    )
  }

  isVisibleEstadoEjecucion(): boolean{
    return this.pregunta.isRespondida && this.questionAnswer.estado!=QuestionAnswer.EstadoEnum.NoProbado
  }

  isEstadoProbado(): boolean{
    return this.questionAnswer.estado == QuestionAnswer.EstadoEnum.Probado
  }

  printEstado(): string{
    switch(this.questionAnswer.estado){
      case QuestionAnswer.EstadoEnum.Ejecutando:
        return "Ejecutando..."
      case QuestionAnswer.EstadoEnum.ErrorCompilacion:
        return "Error de compilación"
      case QuestionAnswer.EstadoEnum.Probado:
        return "Compilado con éxito"
      case QuestionAnswer.EstadoEnum.NoProbado:
        return "No probado"
      default:
        return ""
    }
  }

  isMostrarPruebas(): boolean{
    return this.isMostrandoPruebas
  }

  switchMostrarPruebas(){
    this.isMostrandoPruebas = !this.isMostrandoPruebas
  }

  isErrorCompilacion(): boolean{
    return this.questionAnswer.estado == QuestionAnswer.EstadoEnum.ErrorCompilacion
  }

  getErrorCompilacionString(): string{
    return this.questionAnswer.errorCompilacion || ""
  }

  showEvaluation(): boolean {
    return true
  }

  isPruebaSuperada(pruebaid: number | undefined): boolean {
    if (pruebaid == undefined) return false
    let ps = this.pruebas.filter((p) =>{return p.id == pruebaid})
    if(ps.length<1) return false
    return ps[0].estado == Prueba.EstadoEnum.Correcto
  }

  printEstadoPrueba(estado: string | undefined): string{
    switch(estado){
      case Prueba.EstadoEnum.Correcto:
        return "Superada"
      case Prueba.EstadoEnum.ErrorRuntime:
        return "Error en tiempo de ejecución"
      case Prueba.EstadoEnum.SalidaIncorrecta:
        return "Salida incorrecta"
      case Prueba.EstadoEnum.TiempoExcedido:
        return "Tiempo límite de ejecución sobrepasado"
      default:
        return "No ejecutada"
    }
  }
}
