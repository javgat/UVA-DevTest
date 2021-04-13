import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Option, Question, QuestionService, Tag, Team, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Mensaje, Pregunta, Tipo } from '../shared/app.model';
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
  teams: Team[]
  userTeams: Team[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.newTag = {
      tag: ""
    }
    this.opciones = []
    this.tags = []
    this.deleteIndex = -1
    this.question = new Pregunta(undefined, "", "", 0, false, true, "", undefined, undefined, Question.TipoPreguntaEnum.String, undefined)
    this.questionEdit = new Pregunta(undefined, "", "", 0, false, true, "", undefined, undefined, Question.TipoPreguntaEnum.String, undefined)
    this.teams = []
    this.userTeams = []
    this.nuevaOpcion = {
      correcta: false,
      texto: ""
    }
    this.tipoPrint = ""
    this.setTipoPrint()
    this.id = 0
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      this.getPregunta(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    this.borrarMensaje()
    super.onDestroy()
  }

  doHasUserAction(){
    this.getUserTeams(true)
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

  getPregunta(primera: boolean) {
    this.qS.getQuestion(this.id).subscribe(
      resp => {
        this.question = new Pregunta(resp.id, resp.title, resp.question, resp.estimatedTime,
          resp.autoCorrect, resp.editable, resp.username, resp.eleccionUnica, resp.solucion,
          resp.tipoPregunta, resp.valorFinal)
        this.questionEdit = new Pregunta(resp.id, resp.title, resp.question, resp.estimatedTime,
          resp.autoCorrect, resp.editable, resp.username, resp.eleccionUnica, resp.solucion,
          resp.tipoPregunta, resp.valorFinal)
        this.setTipoPrint()
        if (this.question.tipoPregunta == Question.TipoPreguntaEnum.Opciones) {
          this.getOptions(true)
        }
        this.getTags(true)
        this.getTeamsPregunta(true)
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
    )
  }

  getUserTeams(primera: boolean){
    this.userS.getTeamsOfUser(this.getSessionUser().getUsername()).subscribe(
      resp =>{
        this.userTeams = resp
      },
      err => this.handleErrRelog(err, "obtener equipos de usuario", primera, this.getUserTeams, this)
    )
  }

  getTeamsPregunta(primera: boolean){
    this.qS.getTeamsFromQuestion(this.id).subscribe(
      resp =>{
        this.teams = resp
      },
      err => this.handleErrRelog(err, "obtener equipos de la pregunta", primera, this.getTeamsPregunta, this)
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

  isInAdminTeam() : boolean{
    for(var i: number = 0; i < this.teams.length; i++){
      for(var j: number = 0; j < this.userTeams.length; j++){
        if(this.teams[i].teamname == this.userTeams[j].teamname)
          return true
      }
    }
    return false
  }

  isPermisosAdministracion() : boolean{
    return (this.getSessionUser().getUsername() == this.question.username) || this.isInAdminTeam()
  }

  checkPermisosEdicion(): boolean {
    return this.question.editable && (this.getSessionUser().isAdmin() || this.isPermisosAdministracion())
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
      resp => { },
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
      err => this.handleErrRelog(err, "añadir una etiqueta a una pregunta", primera, this.addTag, this)
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
}