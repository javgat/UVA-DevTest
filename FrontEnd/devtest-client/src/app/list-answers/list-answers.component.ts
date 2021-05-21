import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-answers',
  templateUrl: './list-answers.component.html',
  styleUrls: ['./list-answers.component.css']
})
export class ListAnswersComponent extends LoggedInController implements OnInit {

  testid: number
  answers: Answer[]
  routeSub: Subscription
  editLikeUsername: string
  likeUsername?: string
  buscarUsuario: boolean
  isInAdminTeam: boolean
  test: Examen
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.testid = 0
    this.answers = []
    this.editLikeUsername = ""
    this.buscarUsuario = true
    this.isInAdminTeam = false
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getPTest(true)
        this.getPTestAnswers()
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.routeSub.unsubscribe()
    super.onDestroy()
  }


  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0) {
      this.getIsInAdminTeam(true)
      this.getPTestAnswers()
      this.getPTest(true)
    }
  }
  
  gotTest() {
    if (!this.getSessionUser().isEmpty())
      this.getIsInAdminTeam(true)
  }

  getPTest(primera: boolean) {
    this.userS.getSolvableTestFromUser(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.gotTest()
      },
      err => {
        this.handleErrRelog(err, "obtener test publicado", primera, this.getPTest, this)
      }
    )
  }

  getPTestAnswers(){
    if(this.likeUsername == undefined)
      this.getPTestAllAnswers(true)
    else
      this.getPTestAnswersFromUser(true)
  }

  //Sobreescribir
  getPTestAllAnswers(primera: boolean){
    this.ptestS.getAnswersFromPublishedTests(this.testid).subscribe(
      resp => {
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas de test", primera, this.getPTestAllAnswers, this)
      }
    )
  }

  //Sobreescribir
  getPTestAnswersFromUser(primera: boolean){
    if(this.likeUsername==undefined) return
    this.userS.getAnswersFromUserAnsweredTest(this.likeUsername, this.testid).subscribe(
      resp =>{
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas de test y usuario", primera, this.getPTestAnswersFromUser, this)
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

  clickSearchUsername(){
    this.likeUsername = this.editLikeUsername
    this.getPTestAnswers()
  }

  clickBorrarUsername(){
    this.likeUsername = undefined
    this.getPTestAnswers()
  }

  modoUsuario(): boolean{
    return !this.isModoTestAdmin()
  }

  printDate(d: Date | undefined): string{
    if(d==undefined) return ""
    var date = new Date(d)
    return date.toLocaleString()
  }

  isModoTestAdmin(): boolean {
    return this.isInAdminTeam || this.test.username == this.getSessionUser().getUsername() || this.getSessionUser().isAdmin()
  }
}
