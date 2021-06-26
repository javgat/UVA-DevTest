import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AnswerService, Option, Prueba, PublishedTestService, Question, QuestionAnswer, QuestionService, Review, Test, UserService } from '@javgat/devtest-api';
import { CodeModel } from '@ngstack/code-editor';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { bgcolorQAnswerPuntuacion, Examen, Pregunta, RespuestaPregunta, ResultadoPruebas, tipoPrint } from '../shared/app.model';
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
  isInAdminTeam: boolean
  test: Test
  theme = 'vs-dark';
  codeModel: CodeModel
  isMostrandoPruebas: boolean
  collapsedPruebaIds: Set<number>

  pruebas: Prueba[]
  resPruebas: ResultadoPruebas

  codeOptions = {
    contextmenu: true,
    minimap: {
      enabled: true,
    },
  };
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, private answerS: AnswerService, private ptestS: PublishedTestService, private qS: QuestionService) {
    super(session, router, data, userS)
    this.questionid = 0
    this.testid = 0
    this.answerid = 0
    this.isInAdminTeam = false
    this.options = []
    this.isMostrandoPruebas = false
    this.question = new Pregunta()
    this.pruebas = []
    this.collapsedPruebaIds = new Set()
    this.resPruebas = new ResultadoPruebas()
    this.editPuntuacion = 0
    this.qa = new RespuestaPregunta()
    this.codeModel = {
      language: 'cpp',
      uri: 'main.cpp',
      value: this.qa.respuesta || "",
    }
    this.showCorregir = false
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.questionid = params['questionid']
      this.testid = params['testid']
      this.answerid = params['answerid']
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getIsInAdminTeam(true)
        this.getQAnswer(true)
        this.getPregunta(true)
        this.getPTest(true)
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
      this.getIsInAdminTeam(true)
      this.getPTest(true)
    }
  }

  isDefinido(num: number | undefined): boolean {
    return (num != undefined && num != 0)
  }

  getQAnswer(primera: boolean) {
    this.answerS.getQuestionAnswerFromAnswer(this.answerid, this.questionid).subscribe(
      resp => {
        this.qa = new RespuestaPregunta(resp)
        this.recargarEditorCodigo()
        this.editPuntuacion = this.qa.puntuacion
        this.getResultadoPruebas()
        if(this.qa.estado == QuestionAnswer.EstadoEnum.Ejecutando){
          setTimeout(() => {this.getQAnswer(true)}, this.TIEMPO_RECARGA_ESTADO_COMPILACION)
        }
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
        } else if (this.question.tipoPregunta == Question.TipoPreguntaEnum.Codigo) {
          this.getPruebas(true)
        }
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
    )
  }

  getOptions(primera: boolean) {
    if (this.getSessionUser().isTeacherOrAdmin())
      this.getOptionsAsTeacher(primera)
    else
      this.getOptionsAsStudent(primera)
  }

  getOptionsAsTeacher(primera: boolean) {
    this.qS.getOptionsFromQuestion(this.questionid).subscribe(
      resp => {
        this.options = resp
      },
      err => this.handleErrRelog(err, "obtener opciones de respuesta de pregunta como propietario", primera, this.getOptions, this)
    )
  }

  getOptionsAsStudent(primera: boolean) {
    this.ptestS.getOptionsFromPublishedQuestion(this.testid, this.questionid).subscribe(
      resp => {
        this.options = resp
      },
      err => this.handleErrRelog(err, "obtener opciones de respuesta de pregunta como usuario", primera, this.getOptions, this)
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

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        this.isInAdminTeam = true
      },
      err => {
        if (err.status != 410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
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
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice)) || (opc.correcta == true && !this.opcionSeleccionada(opc.indice))
  }

  checkShowTick(opc: Option): boolean {
    return (opc.correcta || false)
  }

  checkShowCross(opc: Option): boolean {
    return (!(opc.correcta) && this.opcionSeleccionada(opc.indice))
  }

  isRadioQuestion(): boolean {
    return this.question.eleccionUnica == true
  }

  isCheckQuestion(): boolean {
    return !this.isRadioQuestion()
  }

  clickShowCorregir() {
    this.showCorregir = true
  }

  clickNotShowCorregir() {
    this.showCorregir = false
  }

  submitUpdateCorrection() {
    this.showCorregir = false
    this.updateCorrection(true)
  }

  updateCorrection(primera: boolean) {
    let review: Review
    review = {
      puntuacion: this.editPuntuacion
    }
    this.answerS.putReview(this.answerid, this.questionid, review).subscribe(
      resp => this.getQAnswer(true),
      err => this.handleErrRelog(err, "actualizar correccion", primera, this.updateCorrection, this)
    )
  }

  setAsNotCorregidaClick() {
    this.setAsNotCorregida(true)
  }

  setAsNotCorregida(primera: boolean) {
    this.answerS.deleteReview(this.answerid, this.questionid).subscribe(
      resp => this.getQAnswer(true),
      err => this.handleErrRelog(err, "borrar correccion", primera, this.setAsNotCorregida, this)
    )
  }

  calcValor(porcentaje: number | undefined, valorFinal: number | undefined): number {
    if (porcentaje == undefined || valorFinal == undefined) return 0
    return (porcentaje * valorFinal) / 100
  }

  isModoTestAdmin(): boolean {
    return this.isInAdminTeam || this.test.username == this.getSessionUser().getUsername()
  }

  bgcolorQAnswerPuntuacion(punt: number) {
    return bgcolorQAnswerPuntuacion(punt)
  }

  isAutoCorrect(): boolean {
    return this.question.autoCorrect
  }

  recargarEditorCodigo() {
    this.codeModel = {
      language: 'cpp',
      uri: 'main.cpp',
      value: this.qa.respuesta || "",
    }
  }

  getPruebas(primera: boolean) {
    if (this.isModoTestAdmin() || this.getSessionUser().isAdmin()) {
      this.getFullPruebas(primera)
    } else {
      this.getVisiblePruebas(primera)
    }
  }

  getFullPruebas(primera: boolean) {
    this.answerS.getPublishedPruebasFromQuestionTest(this.answerid, this.questionid).subscribe(
      resp => this.pruebas = resp,
      err => this.handleErrRelog(err, "obtener todas las pruebas de pregunta", primera, this.getFullPruebas, this)
    )
  }

  getVisiblePruebas(primera: boolean) {
    this.answerS.getVisiblePublishedPruebasFromQuestionTest(this.answerid, this.questionid).subscribe(
      resp => this.pruebas = resp,
      err => this.handleErrRelog(err, "obtener pruebas visibles de pregunta", primera, this.getVisiblePruebas, this)
    )
  }

  getResultadoPruebas() {
    if (this.isModoTestAdmin() || this.getSessionUser().isAdmin()) {
      this.getFullResultadoPruebas(true)
    } else {
      this.getPreResultadoPruebas(true)
    }
  }

  getPreResultadoPruebas(primera: boolean) {
    this.answerS.getPreTesting(this.answerid, this.questionid).subscribe(
      resp => {
        this.resPruebas = resp
      },
      err => this.handleErrRelog(err, "obtener resultado de pretesting de pruebas", primera, this.getPreResultadoPruebas, this)
    )
  }

  getFullResultadoPruebas(primera: boolean) {
    this.answerS.getFullTesting(this.answerid, this.questionid).subscribe(
      resp => {
        this.resPruebas = resp
      },
      err => this.handleErrRelog(err, "obtener resultado de fulltesting de pruebas", primera, this.getFullResultadoPruebas, this)
    )
  }

  isVisibleEstadoEjecucion(): boolean {
    return this.qa.estado != QuestionAnswer.EstadoEnum.NoProbado
  }

  isEstadoProbado(): boolean {
    return this.qa.estado == QuestionAnswer.EstadoEnum.Probado
  }

  printEstado(): string {
    switch (this.qa.estado) {
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

  isMostrarPruebas(): boolean {
    return this.isMostrandoPruebas
  }

  switchMostrarPruebas() {
    this.isMostrandoPruebas = !this.isMostrandoPruebas
  }

  isErrorCompilacion(): boolean {
    return this.qa.estado == QuestionAnswer.EstadoEnum.ErrorCompilacion
  }

  isEjecutando(): boolean{
    return this.qa.estado == QuestionAnswer.EstadoEnum.Ejecutando
  }

  getErrorCompilacionString(): string {
    return this.qa.errorCompilacion || ""
  }

  showTextoVisibles(): boolean {
    return !(this.isModoTestAdmin() || this.getSessionUser().isAdmin())
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

  getPruebaIndice(id: number | undefined): string{
    if(id==undefined) return ""
    for(let i = 0; i<this.pruebas.length; i++){
      if(id==this.pruebas[i].id){
        let valor = i+1
        return valor.toString()
      }
    }
    return ""
  }

  isCollapsed(id: number | undefined): boolean{
    if(id==undefined) return true
    return this.collapsedPruebaIds.has(id)
  }

  switchCollapse(id: number | undefined){
    if(id==undefined) return
    if(this.collapsedPruebaIds.has(id)){
      this.collapsedPruebaIds.delete(id)
    }else{
      this.collapsedPruebaIds.add(id)
    }
  }

}
