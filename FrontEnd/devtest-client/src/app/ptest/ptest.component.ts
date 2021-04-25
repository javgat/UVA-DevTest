import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, Question, Tag, Test, TestService, UserService } from '@javgat/devtest-api';
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
  test: Test
  preguntas: Question[]
  id: number
  tags: Tag[]
  isInAdminTeam: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private ptestS: PublishedTestService) {
    super(session, router, data, userS);
    this.test = new Examen()
    this.preguntas = []
    this.tags = []
    this.id=0
    this.isInAdminTeam = false
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      this.getPTest(true)
    });
  }

  ngOnInit(): void {
  }

  doHasUserAction() {
    if (this.id != undefined && this.id != 0)
      this.getIsInAdminTeam(true)
  }

  gotTest(){
    this.getPreguntasTest(true)
    //this.getTags(true)
    if (!this.getSessionUser().isEmpty())
      this.getIsInAdminTeam(true)
    //this.getIsFavorita(true)
  }

  getPTest(primera: boolean) {
    this.ptestS.getPublicPublishedTest(this.id).subscribe(
      resp => {
        this.test = new Examen(resp.title, resp.description, resp.accesoPublico, resp.editable, resp.maxMinutes, resp.username, resp.id, resp.accesoPublicoNoPublicado)
        this.gotTest()
      },
      err => {
        if(err.status==410){
          this.getPrivatePTest(true)
        }else{
          this.handleErrRelog(err, "obtener test publico", primera, this.getPTest, this)
        }
      }
    )
  }

  getPrivatePTest(primera: boolean){
    this.ptestS.getPublishedTest(this.id).subscribe(
      resp => {
        this.test = new Examen(resp.title, resp.description, resp.accesoPublico, resp.editable, resp.maxMinutes, resp.username, resp.id, resp.accesoPublicoNoPublicado)
        this.gotTest()
      },
      err => {
        this.handleErrRelog(err, "obtener test privado", primera, this.getPrivatePTest, this)
      }
    )
  }

  getPreguntasTest(primera: boolean){
    this.ptestS.getQuestionsFromPublishedTests(this.id).subscribe(
      resp => {
        this.preguntas = resp
      },
      err => {
        this.handleErrRelog(err, "obtener preguntas de test publicado", primera, this.getPreguntasTest, this)
      }
    )
  }

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if (err.status != 410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string{
    return tipoPrint(tipo, eleccionUnica)
  }

}
