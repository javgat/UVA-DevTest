import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, Tag, Test, TestService, UserService, ValorFinal } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen, Mensaje, Tipo, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';
@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css']
})
export class TestComponent extends LoggedInTeacherController implements OnInit {

  routeSub: Subscription
  id: number
  test: Test
  testEdit: Test
  preguntas: Question[]
  addQuestionId: number
  preguntaChange?: Question
  isInAdminTeam: boolean
  tags: Tag[]
  newTag : string
  deletingTag: string
  mantenerMensaje: boolean
  addQuestionById: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private testS: TestService) {
    super(session, router, data, userS)
    this.isInAdminTeam = false
    this.id = 0
    this.addQuestionId = 0
    this.test = new Examen()
    this.testEdit = new Examen()
    this.preguntas = []
    this.tags = []
    this.newTag = ""
    this.deletingTag = ""
    this.mantenerMensaje = false
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      if(!this.mantenerMensaje){
        this.borrarMensaje()
      }
      this.getTest(true)
    });
    this.addQuestionById = false
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
    this.routeSub.unsubscribe()
    if(!this.mantenerMensaje){
      this.borrarMensaje()
    }
  }

  doHasUserAction() {
    if (this.id!=undefined && this.id != 0)
      this.getIsInAdminTeam(true)
  }

  getTest(primera: boolean) {
    this.testS.getTest(this.id).subscribe(
      resp => {
        this.test = new Examen(resp.title, resp.description, resp.accesoPublico, resp.editable, resp.maxMinutes, resp.username, resp.id, resp.accesoPublicoNoPublicado)
        this.testEdit = new Examen(resp.title, resp.description, resp.accesoPublico, resp.editable, resp.maxMinutes, resp.username, resp.id, resp.accesoPublicoNoPublicado)
        this.getPreguntasTest(true)
        this.getTags(true)
        if (!this.getSessionUser().isEmpty())
          this.getIsInAdminTeam(true)
      },
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  getTags(primera: boolean){
    this.testS.getTagsFromTest(this.id).subscribe(
      resp => this.tags = resp,
      err => this.handleErrRelog(err, "obtener etiquetas de test", primera, this.getTags, this)
    )
  }

  addTagSubmit(){
    this.addTag(true)
  }

  addTag(primera: boolean){
    this.testS.addTagToTest(this.id, this.newTag).subscribe(
      resp => {
        this.getTags(true)
        this.newTag = ""
      },
      err => this.handleErrRelog(err, "añadir etiqueta a test", primera, this.addTag, this)
    )
  }

  deleteTagClick(tag: string){
    this.deletingTag = tag
    this.deleteTag(true)
  }

  deleteTag(primera: boolean){
    this.testS.removeTagFromTest(this.id, this.deletingTag).subscribe(
      resp => {
        this.getTags(true)
      },
      err => this.handleErrRelog(err, "eliminar etiqueta de test", primera, this.deleteTag, this)
    )
  }

  getPreguntasTest(primera: boolean){
    this.testS.getQuestionsFromTest(this.id).subscribe(
      resp => this.preguntas = resp,
      err => this.handleErrRelog(err, "obtener preguntas de un test", primera, this.getPreguntasTest, this)   
    )
  }

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if(err.status!=410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  isPermisosAdministracion(): boolean{
    return this.getSessionUser().isAdmin() || (this.getSessionUser().getUsername() == this.test.username) || this.isInAdminTeam
  }

  checkPermisosEdicion(): boolean {
    return this.test.editable && this.isPermisosAdministracion()
  }

  addQuestionSubmit(){
    this.addQuestion(true)
  }

  isQuestionInTest(idQ: number): boolean{
    return !this.preguntas.every(element =>{
      return idQ != element.id
    })
  }

  addQuestion(primera: boolean){
    if(this.isQuestionInTest(this.addQuestionId)){
      this.cambiarMensaje(new Mensaje("La pregunta ya está en el test", Tipo.ERROR, true))
      return
    }
    let vF : ValorFinal = {
      valorFinal: 1
    }
    this.testS.addQuestionToTest(this.id, this.addQuestionId, vF).subscribe(
      resp => {
        this.getPreguntasTest(true)
      },
      err => this.handleErrRelog(err, "añadir pregunta a test", primera, this.addQuestion, this)
    )
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined): string{
    return tipoPrint(tipo, eleUni)
  }

  changeValueSubmit(pregunta : Question){
    this.preguntaChange = pregunta
    this.changeValue(true)
  }

  changeValue(primera: boolean){
    if(this.preguntaChange == null || this.preguntaChange.id == null || this.preguntaChange.valorFinal == null){
      return
    }
    let vF : ValorFinal = {
      valorFinal: this.preguntaChange.valorFinal
    }
    this.testS.addQuestionToTest(this.id, this.preguntaChange.id, vF).subscribe(
      resp => this.cambiarMensaje(new Mensaje("Valor cambiado con éxito", Tipo.SUCCESS, true)),
      err => this.handleErrRelog(err, "cambiar valor de una pregunta", primera, this.changeValue, this)
    )
  }

  putTestSubmit(){
    this.putTest(true)
  }

  putTest(primera: boolean){
    this.testS.putTest(this.id, this.testEdit).subscribe(
      resp => this.getTest(true),
      err => this.handleErrRelog(err, "actualizar datos de test", primera, this.putTest, this)
    )
  }

  checkCloneTest(): boolean {
    if (this.test.accesoPublicoNoPublicado) {
      return this.getSessionUser().isTeacherOrAdmin()
    } else {
      return this.isPermisosAdministracion()
    }
  }

  cloneTestClick(){
    this.cloneTest(true)
  }

  cloneTest(primera: boolean){
    this.userS.copyTest(this.getSessionUser().getUsername(), this.id).subscribe(
      resp=>{
        this.cambiarMensaje(new Mensaje("Test clonado con éxito", Tipo.SUCCESS, true))
        this.mantenerMensaje = true
        this.router.navigate(['/et', resp.id])
      },
      err => this.handleErrRelog(err, "clonar test", primera, this.cloneTest, this)
    )
  }

  changeAddByID(){
    this.addQuestionById=true
  }

  changeNotAddByID(){
    this.addQuestionById=false
  }

  questionPicked(id: number){
    this.addQuestionId=id
    this.addQuestionSubmit()
  }
}
