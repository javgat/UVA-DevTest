import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, Question, Tag, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptest',
  templateUrl: './ptest.component.html',
  styleUrls: ['./ptest.component.css']
})
export class PtestComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  test: Examen
  preguntas: Question[]
  id: number
  tags: Tag[]
  isInAdminTeam: boolean
  respuestaIniciadaId?: number
  isRespuestaIniciada: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private ptestS: PublishedTestService) {
    super(session, router, data, userS);
    this.preguntas = []
    this.test = new Examen()
    this.tags = []
    this.id = 0
    this.isInAdminTeam = false
    this.isRespuestaIniciada = false
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && (!this.getSessionLogin().isLoggedIn() || this.getSessionUser().getUsername() != "")) {
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

  hasPermissions(): boolean {
    return this.canVerPTests()
  }

  doHasUserAction() {
    if (this.id != undefined && this.id != 0) {
      this.getIsInAdminTeam(true)
      this.getPTest(true)
    }
  }

  gotTest() {
    this.getPreguntasTest(true)
    this.getTags(true)
    if (!this.getSessionUser().isEmpty())
      this.getIsInAdminTeam(true)
    this.getIsRespuestaIniciada(true)
  }

  getPTest(primera: boolean) {
    if (this.getSessionLogin().isLoggedIn()) {
      this.userS.getSolvableTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
        resp => {
          this.test = Examen.constructorFromTest(resp)
          this.gotTest()
        },
        err => {
          this.handleErrRelog(err, "obtener test publicado", primera, this.getPTest, this)
        }
      )
    } else {
      this.getPublicTest(primera)
    }
  }

  getPublicTest(primera: boolean) {
    this.ptestS.getPublicPublishedTest(this.id).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.gotTest()
      },
      err => {
        this.handleErrRelog(err, "obtener test publicado publico", primera, this.getPTest, this)
      }
    )
  }

  getPreguntasTest(primera: boolean) {
    if (!this.isModoTestAdmin()) return
    this.ptestS.getQuestionsFromPublishedTests(this.id).subscribe(
      resp => {
        this.preguntas = resp
      },
      err => {
        this.handleErrRelog(err, "obtener preguntas de test publicado", primera, this.getPreguntasTest, this)
      }
    )
  }

  getTags(primera: boolean) {
    this.ptestS.getTagsFromPublishedTest(this.id).subscribe(
      resp => this.tags = resp,
      err => this.handleErrRelog(err, "obtener etiquetas de un test publicado", primera, this.getTags, this)
    )
  }

  getIsInAdminTeam(primera: boolean) {
    if (this.getSessionLogin().isLoggedIn()) {
      this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
        resp => {
          this.isInAdminTeam = true
          this.getPreguntasTest(true)
        },
        err => {
          if (err.status != 410)
            this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
        }
      )
    }
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string {
    return tipoPrint(tipo, eleccionUnica)
  }

  isModoTestAdmin(): boolean {
    return this.isInAdminTeam || (this.test.username == this.getSessionUser().getUsername() && this.test.username!="") || this.getSessionUser().isAdmin()
  }

  getIsRespuestaIniciada(primera: boolean) {
    if (this.getSessionLogin().isLoggedIn()) {
      this.userS.getOpenAnswersFromUserTest(this.getSessionUser().getUsername(), this.id).subscribe(
        resp => {
          if (resp.length == 0) {
            this.isRespuestaIniciada = false
          } else {
            this.isRespuestaIniciada = true
            this.respuestaIniciadaId = resp[0].id
          }
        },
        err => this.handleErrRelog(err, "obtener informacion de si hay respuesta iniciada", primera, this.getIsRespuestaIniciada, this)
      )
    } else {
      this.isRespuestaIniciada = false
    }
  }

  startAnswerClick() {
    this.startAnswer(true)
  }

  startAnswer(primera: boolean) {
    this.userS.startAnswer(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => {
        this.router.navigate(['/pt', this.id, "answering"])
      },
      err => {
        this.handleErrRelog(err, "iniciar respuesta a test publicado", primera, this.startAnswer, this)
      }
    )
  }

  mostrarVerMisRespuestas(): boolean{
    if(this.test.cantidadRespuestasDelUsuario==undefined) return false
    return this.isLoggedIn() && this.test.cantidadRespuestasDelUsuario>0 && (!this.isRespuestaIniciada || this.test.cantidadRespuestasDelUsuario>1)
  }

  testHasMaxIntentos(): boolean{
    return this.test.maxIntentos>0
  }

  quedanMasIntentos(): boolean{
    return (this.test.maxIntentos<1) || this.test.maxIntentos>this.test.cantidadRespuestasDelUsuario
  }
}
