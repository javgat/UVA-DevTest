import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Option, Prueba, Question, QuestionService, Tag, TagService, Team, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Mensaje, Pregunta, PruebaEjecucion, Tipo, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';


@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css']
})
export class QuestionComponent extends LoggedInTeacherController implements OnInit {

  question: Question
  questionEdit: Question
  routeSub: Subscription
  id: number
  tipoPrint: string
  opciones: Option[]
  op: Option | undefined
  nuevaOpcion: Option
  deleteIndex: number
  newTag: Tag
  tags: Tag[]
  isInAdminTeam: boolean
  deletingTag: string
  mantenerMensaje: boolean
  isFavorita: boolean
  testid?: number
  autotags: Tag[]
  showExtraInfo: boolean
  pruebaEdit: Prueba
  pruebas: Prueba[]
  pruebasidsModificando: Set<number>
  pruebaidActualizar: number
  pruebaidRecuperar: number
  private editandoRespuesta: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService,
    private route: ActivatedRoute, private tagS: TagService) {
    super(session, router, data, userS)
    this.isInAdminTeam = false
    this.editandoRespuesta = false
    this.newTag = {
      tag: ""
    }
    this.opciones = []
    this.pruebas = []
    this.pruebaidRecuperar = 0
    this.pruebasidsModificando = new Set()
    this.pruebaidActualizar = 0
    this.tags = []
    this.deleteIndex = -1
    this.deletingTag = ""
    this.mantenerMensaje = false
    this.question = new Pregunta()
    this.questionEdit = new Pregunta()
    this.isFavorita = false
    this.showExtraInfo = false
    this.pruebaEdit = new PruebaEjecucion()
    this.nuevaOpcion = {
      correcta: false,
      texto: ""
    }
    this.tipoPrint = ""
    this.setTipoPrint()
    this.id = 0
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.testid = params['testid']
      if (!this.mantenerMensaje) {
        this.borrarMensaje()
      }
      this.getPregunta(true)
    });
    this.autotags = []
    this.changeGetAutoTags()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    if (!this.mantenerMensaje) {
      this.borrarMensaje()
    }
    super.onDestroy()
  }

  doHasUserAction() {
    if (this.id != undefined && this.id != 0){
      this.getIsInAdminTeam(true)
      this.getIsFavorita(true)
    }
  }

  setTipoPrint() {
    switch (this.question.tipoPregunta) {
      case Question.TipoPreguntaEnum.String:
        this.tipoPrint = "Texto"
        break
      case Question.TipoPreguntaEnum.Codigo:
        this.tipoPrint = "Código"
        break
      case Question.TipoPreguntaEnum.Opciones:
        if (this.question.eleccionUnica)
          this.tipoPrint = "Tipo test de respuesta única"
        else
          this.tipoPrint = "Tipo test de respuesta múltiple"
        break
      default:
        this.tipoPrint = ""
    }
  }

  tipPrint(tipo: string, eleUni: boolean | undefined): string {
    return tipoPrint(tipo, eleUni)
  }

  getPregunta(primera: boolean) {
    this.qS.getQuestion(this.id).subscribe(
      resp => {
        this.question = Pregunta.constructorFromQuestion(resp)
        this.questionEdit = Pregunta.constructorFromQuestion(resp)
        this.setTipoPrint()
        if (this.question.tipoPregunta == Question.TipoPreguntaEnum.Opciones) {
          this.getOptions(true)
        }else if(this.question.tipoPregunta == Question.TipoPreguntaEnum.Codigo){
          this.getPruebas(true)
        }
        this.getTags(true)
        if (!this.getSessionUser().isEmpty()) {
          this.getIsInAdminTeam(true)
          this.getIsFavorita(true)
        }
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
    )
  }

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedQuestionFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if (err.status != 410)
          this.handleErrRelog(err, "saber si el usuario administra la pregunta", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  getOptions(primera: boolean) {
    this.qS.getOptionsFromQuestion(this.id).subscribe(
      resp => {
        this.opciones = resp
      },
      err => this.handleErrRelog(err, "obtener opciones de respuesta de pregunta", primera, this.getOptions, this)
    )
  }

  getPruebas(primera: boolean) {
    this.qS.getPruebasFromQuestion(this.id).subscribe(
      resp => {
        this.pruebas = resp
      },
      err => this.handleErrRelog(err, "obtener pruebas de respuesta de pregunta", primera, this.getPruebas, this)
    )
  }

  isPermisosAdministracion(): boolean {
    return this.getSessionUser().isAdmin() || (this.getSessionUser().getUsername() == this.question.username) || this.isInAdminTeam
  }

  checkPermisosEdicion(): boolean {
    return this.question.editable && this.isPermisosAdministracion()
  }

  changeCorrectaOpc(op: Option) {
    this.op = op
    this.changeCorrecta(true)
  }

  changeCorrecta(primera: boolean) {
    if (this.op == undefined) return
    this.op.correcta = !this.op.correcta
    if (this.op.indice == undefined) {
      this.cambiarMensaje(new Mensaje("Error, falta el campo indice en la opcion", Tipo.ERROR, true))
      return
    }
    this.qS.putOption(this.id, this.op.indice, this.op).subscribe(
      resp => {
        this.cambiarMensaje(new Mensaje("Valor de opción modificado", Tipo.SUCCESS, true))
      },
      err => this.handleErrRelog(err, "cambiar campo de correccion de una opcion", primera, this.changeCorrecta, this)
    )
  }

  addOptionSubmit() {
    //this.nuevaOpcion.preguntaid = this.id
    this.addOption(true)
  }

  addOption(primera: boolean) {
    this.qS.postOption(this.id, this.nuevaOpcion).subscribe(
      resp => {
        this.getOptions(true)
        this.nuevaOpcion.texto = ""
      },
      err => this.handleErrRelog(err, "añadir nueva opcion de respuesta", primera, this.addOption, this)
    )
  }

  deleteOpc(indice: number | undefined) {
    if (indice != undefined)
      this.deleteIndex = indice
    this.deleteOpcion(true)
  }

  deleteOpcion(primera: boolean) {
    this.qS.deleteOption(this.id, this.deleteIndex).subscribe(
      resp => this.getOptions(true),
      err => this.handleErrRelog(err, "eliminar una opcion", primera, this.deleteOpcion, this)
    )
  }

  addTagSubmit() {
    this.addTag(true)
  }

  addTag(primera: boolean) {
    this.qS.addTagToQuestion(this.id, this.newTag.tag).subscribe(
      resp => {
        this.getTags(true)
      },
      err => {
        if (err.status == 409) {
          this.cambiarMensaje(new Mensaje("Esa etiqueta ya está añadida", Tipo.ERROR, true))
        } else {
          this.handleErrRelog(err, "añadir una etiqueta a una pregunta", primera, this.addTag, this)
        }
      }
    )
  }

  getTags(primera: boolean) {
    this.qS.getTagsFromQuestion(this.id).subscribe(
      resp => this.tags = resp,
      err => this.handleErrRelog(err, "obtener etiquetas de una pregunta", primera, this.getTags, this)
    )
  }

  onSelectTipoString() {
    this.questionEdit.tipoPregunta = Question.TipoPreguntaEnum.String
    this.questionEdit.eleccionUnica = undefined
  }

  onSelectTipoRadio() {
    this.questionEdit.tipoPregunta = Question.TipoPreguntaEnum.Opciones
    this.questionEdit.eleccionUnica = true
  }

  onSelectTipoCheck() {
    this.questionEdit.tipoPregunta = Question.TipoPreguntaEnum.Opciones
    this.questionEdit.eleccionUnica = false
  }

  onSelectTipoCode() {
    this.questionEdit.tipoPregunta = Question.TipoPreguntaEnum.Codigo
    this.questionEdit.eleccionUnica = undefined
  }

  putQuestionSubmit() {
    this.stopEditarRespuesta()
    this.putQuestion(true)
  }

  putQuestion(primera: boolean) {
    this.qS.putQuestion(this.id, this.questionEdit).subscribe(
      resp => {
        this.getPregunta(true)
        this.cambiarMensaje(new Mensaje("Pregunta actualizada", Tipo.SUCCESS, true))
      },
      err => this.handleErrRelog(err, "modificar pregunta", primera, this.putQuestion, this)
    )
  }

  deleteTagClick(tag: string) {
    this.deletingTag = tag
    this.deleteTag(true)
  }

  deleteTag(primera: boolean) {
    this.qS.removeTagFromQuestion(this.id, this.deletingTag).subscribe(
      resp => this.getTags(true),
      err => this.handleErrRelog(err, "eliminar etiqueta de una pregunta", primera, this.deleteTag, this)
    )
  }

  checkCloneQuestion(): boolean {
    if (this.question.accesoPublicoNoPublicada) {
      return this.getSessionUser().isTeacherOrAdmin()
    } else {
      return this.isPermisosAdministracion()
    }
  }

  cloneQuestionClick() {
    this.cloneQuestion(true)
  }

  cloneQuestion(primera: boolean) {
    if (this.question.id == undefined) return
    this.userS.copyQuestion(this.getSessionUser().getUsername(), this.question.id).subscribe(
      resp => {
        this.cambiarMensaje(new Mensaje("Pregunta clonada con éxito", Tipo.SUCCESS, true))
        this.mantenerMensaje = true
        this.router.navigate(['/q', resp.id])
      },
      err => this.handleErrRelog(err, "clonar pregunta", primera, this.cloneQuestion, this)
    )
  }

  getIsFavorita(primera: boolean) {
    this.userS.getFavoriteQuestion(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isFavorita = true,
      err => {
        if (err.status == 410) {
          this.isFavorita = false
        } else {
          this.handleErrRelog(err, "ver si la pregunta esta marcada como favorita", primera, this.getIsFavorita, this)
        }
      }
    )
  }

  changeFavorita() {
    if (this.isFavorita) {
      this.removeFavorita(true)
    } else {
      this.addFavorita(true)
    }
  }

  addFavorita(primera: boolean) {
    this.userS.addQuestionFavorite(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => {
        this.getPregunta(true)
      },
      err => {
        this.handleErrRelog(err, "marcar como favorita una pregunta", primera, this.addFavorita, this)
      }
    )
  }

  removeFavorita(primera: boolean) {
    this.userS.removeQuestionFavorite(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.getPregunta(true),
      err => {
        this.handleErrRelog(err, "desmarcar como favorita una pregunta", primera, this.removeFavorita, this)
      }
    )
  }

  changeGetAutoTags() {
    this.getAutoTags(true)
  }

  getAutoTags(primera: boolean) {
    this.tagS.getTags(this.newTag.tag, "moreQuestion", 20).subscribe(
      resp => {
        this.autotags = resp
      },
      err => this.handleErrRelog(err, "obtener tags de preguntas mas comunes", primera, this.getAutoTags, this)
    )
  }

  anadirOpcionDisabled(): boolean {
    return this.nuevaOpcion.texto == "" || this.nuevaOpcion.texto == undefined
  }

  clickEditarRespuesta(){
    this.editandoRespuesta = true
  }

  clickStopEditarRespuesta(){
    this.stopEditarRespuesta()
  }

  stopEditarRespuesta(){
    this.editandoRespuesta = false
  }

  isEditandoRespuesta(): boolean{
    return this.editandoRespuesta
  }

  needsStopEditarRespuesta(): boolean{
    return this.question.tipoPregunta == Question.TipoPreguntaEnum.Opciones || this.question.tipoPregunta == Question.TipoPreguntaEnum.Codigo
  }

  showVolverTest(): boolean{
    return this.testid!=undefined
  }

  showVolverMisQuestions(): boolean{
    return this.getSessionUser().getUsername() == this.question.username && !this.showVolverTest()
  }

  showVolverQuestionsCompartidos(): boolean{
    return this.isInAdminTeam && !this.showVolverMisQuestions() && !this.showVolverTest()
  }

  showVolverQuestions(): boolean{
    return !this.showVolverMisQuestions() && !this.showVolverQuestionsCompartidos() && !this.showVolverTest()
  }

  swapShowExtraInfo(){
    this.showExtraInfo = !this.showExtraInfo
  }

  showMostrarExtraInfo(): boolean{
    return this.showExtraInfo
  }

  isCodeQuestion(): boolean{
    return this.question.tipoPregunta == Question.TipoPreguntaEnum.Codigo
  }
  
  isCorreccionAutomatica(): boolean{
    return this.question.autoCorrect
  }

  clickAddPrueba(){
    this.pruebaEdit = new PruebaEjecucion()
  }

  postPruebaSubmit(){
    this.postPrueba(true)
  }

  postPrueba(primera: boolean){
    this.qS.postPrueba(this.id, this.pruebaEdit).subscribe(
      resp => {
        this.getPruebas(true)
      },
      err => this.handleErrRelog(err, "crear nueva prueba de ejecución", primera, this.postPrueba, this)
    )
  }

  modificarPrueba(pruebaid: number | undefined){
    this.pruebasidsModificando.add(pruebaid || 0)
  }

  stopModificarPrueba(pruebaid: number | undefined){
    this.pruebasidsModificando.delete(pruebaid || 0)
  }

  modificandoPrueba(pruebaid: number | undefined){
    return this.pruebasidsModificando.has(pruebaid || 0)
  }

  cancelarModificarPrueba(pruebaid: number| undefined){
    if(pruebaid == undefined) return
    this.stopModificarPrueba(pruebaid)
    this.pruebaidRecuperar = pruebaid
    this.recuperarPrueba(true)
  }

  recuperarPrueba(primera: boolean){
    this.qS.getPruebaFromQuestion(this.id, this.pruebaidRecuperar).subscribe(
      resp => {
        this.pruebas.map( (p, i, arr) => {
          if(p.id != undefined && p.id == this.pruebaidRecuperar){
            p.entrada = resp.entrada
            p.salida = resp.salida
            p.visible = resp.visible
            p.postEntrega = resp.postEntrega
          }
        } )
      },
      err => this.handleErrRelog(err, "recuperar valor original de prueba", primera, this.recuperarPrueba, this)
    )
  }

  guardarModificarPrueba(pruebaid: number | undefined){
    if(pruebaid == undefined) return
    this.stopModificarPrueba(pruebaid)
    this.pruebaidActualizar = pruebaid
    this.actualizarPrueba(true)
  }

  actualizarPrueba(primera: boolean){

    let prus = this.pruebas.filter( (p, i, arr) => {
      return (p.id != undefined && p.id == this.pruebaidActualizar)
    })
    if(prus.length==0) return
    let pru = prus[0]
    this.qS.putPrueba(this.id, this.pruebaidActualizar, pru).subscribe(
      resp => {
        this.getPruebas(true)
      },
      err => this.handleErrRelog(err, "actualizar valor de prueba", primera, this.actualizarPrueba, this)
    )
  }
}
