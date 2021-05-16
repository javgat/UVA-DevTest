import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, AnswerService, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen, Respuesta } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-answer',
  templateUrl: './answer.component.html',
  styleUrls: ['./answer.component.css']
})
export class AnswerComponent extends LoggedInController implements OnInit {

  routeSub: Subscription
  testid: number
  answerid: number
  answer: Respuesta
  isInAdminTeam: boolean
  test: Test
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private answerS: AnswerService) {
    super(session, router, data, userS)
    this.testid = 0
    this.answer = new Respuesta()
    this.answerid = 0
    this.isInAdminTeam = false
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.answerid = params['answerid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getIsInAdminTeam(true)
        this.getAnswer(true)
        this.getPTest(true)
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
    this.routeSub.unsubscribe()
  }

  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0 && this.answerid != undefined && this.answerid != 0) {
      this.getIsInAdminTeam(true)
      this.getAnswer(true)
      this.getPTest(true)
    }
  }

  getAnswer(primera: boolean) {
    this.answerS.getAnswer(this.answerid).subscribe(
      resp => {
        this.answer = new Respuesta(resp)
      },
      err => this.handleErrRelog(err, "obtener respuesta a examen", primera, this.getAnswer, this)
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

  isModoTestAdmin(): boolean {
    return this.isInAdminTeam || this.test.username == this.getSessionUser().getUsername()
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

  setAsCorregidaClick(){
    this.setAsCorregida(true)
  }

  setAsCorregida(primera: boolean){
    this.answerS.setAnswerCorrected(this.answerid).subscribe(
      resp => this.getAnswer(true),
      err => this.handleErrRelog(err, "marcar respuesta como corregida", primera, this.setAsCorregida, this)
    )
  }

  setAsNotCorregidaClick(){
    this.setAsNotCorregida(true)
  }

  setAsNotCorregida(primera: boolean){
    this.answerS.setAnswerNotCorrected(this.answerid).subscribe(
      resp => this.getAnswer(true),
      err => this.handleErrRelog(err, "marcar pregunta como no corregida", primera, this.setAsNotCorregida, this)
    )
  }

  clickHacerVisible(){
    if(this.answer.visibleParaUsuario)
      this.hacerNoVisible(true)
    else
      this.hacerVisible(true)
  }

  hacerVisible(primera: boolean){
    this.answerS.setAnswerVisible(this.answerid).subscribe(
      resp => this.getAnswer(true),
      err => this.handleErrRelog(err, "hacer visible para usuario la respuesta", primera, this.hacerVisible, this)
    )
  }

  hacerNoVisible(primera: boolean){
    this.answerS.setAnswerNotVisible(this.answerid).subscribe(
      resp => this.getAnswer(true),
      err => this.handleErrRelog(err, "hacer no visible para usuario la respuesta", primera, this.hacerNoVisible, this)
    )
  }

  printDate(d: Date | undefined): string{
    if(d==undefined) return ""
    var date = new Date(d)
    return date.toLocaleString()
  }

}
