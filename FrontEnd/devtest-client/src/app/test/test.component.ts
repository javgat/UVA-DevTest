import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishTestParams, Question, Tag, TagService, Test, TestService, UserService } from '@javgat/devtest-api';
import { TestPregunta } from '@javgat/devtest-api/model/testPregunta';
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
  changePosPreguntas: Question[]
  addQuestionId: number
  preguntaChange?: Question
  isInAdminTeam: boolean
  tags: Tag[]
  newTag: string
  deletingTag: string
  mantenerMensaje: boolean
  addQuestionById: boolean
  preguntaQuitando: number
  isFavorita: boolean
  publishedParams: PublishTestParams
  autotags: Tag[]
  tPTempPosicion?: TestPregunta
  pregTempPosicion?: Question
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute,
    private testS: TestService, private tagS: TagService) {
    super(session, router, data, userS)
    this.isInAdminTeam = false
    this.id = 0
    this.addQuestionId = 0
    this.test = new Examen()
    this.testEdit = new Examen()
    this.preguntas = []
    this.changePosPreguntas = []
    this.tags = []
    this.newTag = ""
    this.deletingTag = ""
    this.mantenerMensaje = false
    this.isFavorita = false
    this.preguntaQuitando = 0
    this.publishedParams = {
      title: "",
      accesoPublico: false,
      autoCorrect: false,
      visibilidad: "manual",
      tiempoEstricto: false,
      maxMinutes: 1
    }
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      if (!this.mantenerMensaje) {
        this.borrarMensaje()
      }
      if (!this.getSessionUser().isEmpty()) {
        this.getTest(true)
        this.getIsInAdminTeam(true)
        this.getIsFavorita(true)
      }
    });
    this.addQuestionById = false
    this.autotags = []
    this.changeGetAutoTags()
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
    this.routeSub.unsubscribe()
    if (!this.mantenerMensaje) {
      this.borrarMensaje()
    }
  }

  doHasUserAction() {
    if (this.id != undefined && this.id != 0) {
      this.getTest(true)
      this.getIsInAdminTeam(true)
      this.getIsFavorita(true)
    }
  }

  getTest(primera: boolean) {
    this.testS.getTest(this.id).subscribe(
      resp => {
        if (!resp.editable) {
          this.router.navigate(['/pt', this.id])
        }
        this.test = Examen.constructorFromTest(resp)
        this.testEdit = Examen.constructorFromTest(resp)
        this.publishedParams.title = this.test.title
        this.publishedParams.accesoPublico = this.test.accesoPublico
        this.publishedParams.autoCorrect = this.test.autoCorrect
        this.publishedParams.maxMinutes =  this.test.maxMinutes
        this.publishedParams.tiempoEstricto = this.test.tiempoEstricto
        this.publishedParams.visibilidad = this.test.visibilidad
        this.getPreguntasTest(true)
        this.getTags(true)
        if (!this.getSessionUser().isEmpty()) {
          this.getIsInAdminTeam(true)
          this.getIsFavorita(true)
        }
      },
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  getTags(primera: boolean) {
    this.testS.getTagsFromTest(this.id).subscribe(
      resp => this.tags = resp,
      err => this.handleErrRelog(err, "obtener etiquetas de test", primera, this.getTags, this)
    )
  }

  addTagSubmit() {
    this.addTag(true)
  }

  addTag(primera: boolean) {
    this.testS.addTagToTest(this.id, this.newTag).subscribe(
      resp => {
        this.getTags(true)
        this.newTag = ""
      },
      err => {
        if (err.status == 409) {
          this.cambiarMensaje(new Mensaje("Esa etiqueta ya está añadida", Tipo.ERROR, true))
        } else {
          this.handleErrRelog(err, "añadir etiqueta a test", primera, this.addTag, this)
        }
      }
    )
  }

  deleteTagClick(tag: string) {
    this.deletingTag = tag
    this.deleteTag(true)
  }

  deleteTag(primera: boolean) {
    this.testS.removeTagFromTest(this.id, this.deletingTag).subscribe(
      resp => {
        this.getTags(true)
      },
      err => this.handleErrRelog(err, "eliminar etiqueta de test", primera, this.deleteTag, this)
    )
  }

  getPreguntasTest(primera: boolean) {
    this.testS.getQuestionsFromTest(this.id).subscribe(
      resp => this.preguntas = resp,
      err => this.handleErrRelog(err, "obtener preguntas de un test", primera, this.getPreguntasTest, this)
    )
  }

  getIsInAdminTeam(primera: boolean) {
    if (this.id == undefined || this.getSessionUser().isEmpty()) return
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if (err.status != 410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  isPermisosAdministracion(): boolean {
    return this.getSessionUser().isAdmin() || (this.getSessionUser().getUsername() == this.test.username) || this.isInAdminTeam
  }

  checkPermisosEdicion(): boolean {
    return this.test.editable && this.isPermisosAdministracion()
  }

  addQuestionSubmit() {
    this.addQuestion(true)
  }

  isQuestionInTest(idQ: number): boolean {
    return !this.preguntas.every(element => {
      return idQ != element.id
    })
  }

  getMaxPositionQuestions(): number {
    let first = this.preguntas.map(item => item.posicion).map(i => i || 0)
    let max: number
    console.log(first)
    if (first.length == 0) {
      max = -1
    } else {
      max = Math.max(...first);
    }
    return max
  }

  addQuestion(primera: boolean) {
    if (this.isQuestionInTest(this.addQuestionId)) {
      this.cambiarMensaje(new Mensaje("La pregunta ya está en el test", Tipo.ERROR, true))
      return
    }
    let tp: TestPregunta = {
      valorFinal: 1,
      posicion: this.getMaxPositionQuestions() + 1
    }
    this.testS.addQuestionToTest(this.id, this.addQuestionId, tp).subscribe(
      resp => {
        this.getTest(true)
      },
      err => this.handleErrRelog(err, "añadir pregunta a test", primera, this.addQuestion, this)
    )
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined): string {
    return tipoPrint(tipo, eleUni)
  }

  changeValueSubmit(pregunta: Question) {
    this.preguntaChange = pregunta
    this.changeValue(true)
  }

  changeValue(primera: boolean) {
    if (this.preguntaChange == null || this.preguntaChange.id == null || this.preguntaChange.valorFinal == null) {
      return
    }
    let tp: TestPregunta = {
      valorFinal: this.preguntaChange.valorFinal,
      posicion: this.preguntaChange.posicion || 0
    }
    this.testS.addQuestionToTest(this.id, this.preguntaChange.id, tp).subscribe(
      resp => this.cambiarMensaje(new Mensaje("Valor cambiado con éxito", Tipo.SUCCESS, true)),
      err => this.handleErrRelog(err, "cambiar valor de una pregunta", primera, this.changeValue, this)
    )
  }

  putTestSubmit() {
    this.putTest(true)
  }

  putTest(primera: boolean) {
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

  cloneTestClick() {
    this.cloneTest(true)
  }

  cloneTest(primera: boolean) {
    this.userS.copyTest(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => {
        this.cambiarMensaje(new Mensaje("Test clonado con éxito", Tipo.SUCCESS, true))
        this.mantenerMensaje = true
        this.router.navigate(['/et', resp.id])
      },
      err => this.handleErrRelog(err, "clonar test", primera, this.cloneTest, this)
    )
  }

  changeAddByID() {
    this.addQuestionById = true
  }

  changeNotAddByID() {
    this.addQuestionById = false
  }

  questionPicked(id: number) {
    this.addQuestionId = id
    this.addQuestionSubmit()
  }

  quitarPreguntaClick(id: number | undefined) {
    this.preguntaQuitando = id || this.preguntaQuitando
    this.quitarPregunta(true)
  }

  quitarPregunta(primera: boolean) {
    if (this.test.id == undefined) return
    this.testS.removeQuestionFromTest(this.test.id, this.preguntaQuitando).subscribe(
      resp => this.getTest(true),
      err => this.handleErrRelog(err, "quitar pregunta de test", primera, this.quitarPregunta, this)
    )
  }

  getIsFavorita(primera: boolean) {
    if (this.id == undefined || this.getSessionUser().isEmpty()) return
    this.userS.getFavoriteTest(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isFavorita = true,
      err => {
        if (err.status == 410) {
          this.isFavorita = false
        } else {
          this.handleErrRelog(err, "ver si el Test esta marcado como favorito", primera, this.getIsFavorita, this)
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
    this.userS.addTestFavorite(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => {
        this.getTest(true)
      },
      err => {
        this.handleErrRelog(err, "marcar como favorito un test", primera, this.addFavorita, this)
      }
    )
  }

  removeFavorita(primera: boolean) {
    this.userS.removeTestFavorite(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.getTest(true),
      err => {
        this.handleErrRelog(err, "desmarcar como favorito un test", primera, this.removeFavorita, this)
      }
    )
  }

  publishTestClick() {
    this.publishTest(true)
  }

  publishTest(primera: boolean) {
    this.testS.postPublishedTest(this.publishedParams, this.id).subscribe(
      resp => {
        this.router.navigate(['/pt', resp.id])
      },
      err => this.handleErrRelog(err, "publicar test", primera, this.publishTest, this)
    )
  }

  visibilidadToString(vis: Test.VisibilidadEnum): string {
    return Examen.visibilidadToString(vis)
  }

  printManual(): string {
    return this.visibilidadToString(Test.VisibilidadEnum.Manual)
  }

  printCorregir(): string {
    return this.visibilidadToString(Test.VisibilidadEnum.AlCorregir)
  }

  printEntregar(): string {
    return this.visibilidadToString(Test.VisibilidadEnum.AlEntregar)
  }

  getValueManual(): Test.VisibilidadEnum {
    return Test.VisibilidadEnum.Manual
  }

  getValueCorregir(): Test.VisibilidadEnum {
    return Test.VisibilidadEnum.AlCorregir
  }

  getValueEntregar(): Test.VisibilidadEnum {
    return Test.VisibilidadEnum.AlEntregar
  }

  changeGetAutoTags() {
    this.getAutoTags(true)
  }

  getAutoTags(primera: boolean) {
    this.tagS.getTags(this.newTag, "moreTest", 20).subscribe(
      resp => {
        this.autotags = resp
      },
      err => this.handleErrRelog(err, "obtener tags de tests mas comunes", primera, this.getAutoTags, this)
    )
  }

  prepareModalPosPreguntas() {
    this.changePosPreguntas = this.preguntas
  }

  changePosPreguntasSubmit() {
    this.preguntas = this.changePosPreguntas
    for (let i = 0; i < this.preguntas.length; i++) {
      let tP: TestPregunta = {
        valorFinal: this.preguntas[i].valorFinal || 0,
        posicion: i
      }
      this.tPTempPosicion = tP
      this.pregTempPosicion = this.preguntas[i]
      this.putPosPregunta(true)
    }
  }

  putPosPregunta(primera: boolean){
    if(this.pregTempPosicion==undefined || this.tPTempPosicion == undefined) return
    this.testS.addQuestionToTest(this.id, this.pregTempPosicion.id || 0, this.tPTempPosicion).subscribe(
      resp => {},
      err => this.handleErrRelog(err, "Actualizar posicion de una pregunta", primera, this.putPosPregunta, this)
    )
  }

  newTabCreateQuestion(){
    const url = this.router.serializeUrl(
      this.router.createUrlTree(['/qCreate'])
    );
  
    window.open(url, '_blank');
  }

}
