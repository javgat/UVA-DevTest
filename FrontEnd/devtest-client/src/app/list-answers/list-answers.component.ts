import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { EnumOrderByAnswer, Examen, Mensaje, Tipo } from '../shared/app.model';
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
  mensajeListaVacia: string
  orderBy: EnumOrderByAnswer
  canOrderByPuntuacion: boolean
  canOrderByDuracion: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.testid = 0
    this.answers = []
    this.editLikeUsername = ""
    this.buscarUsuario = true
    this.isInAdminTeam = false
    this.orderBy = EnumOrderByAnswer.newStartDate
    this.test = new Examen()
    this.canOrderByDuracion = false
    this.canOrderByPuntuacion = false
    this.mensajeListaVacia = "¡Vaya! Parece que aún no hay respuestas para mostrar en esta lista"
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
    this.cambiarMensaje(new Mensaje("Descargando respuestas... ", Tipo.DOWNLOADING, true))
    if(this.likeUsername == undefined)
      this.getPTestAllAnswers(true)
    else
      this.getPTestAnswersFromUser(true)
  }

  ptestAnswersRecieved(resp: Answer[]){
    this.borrarMensaje()
    this.answers = resp
  }

  //Sobreescribir
  getPTestAllAnswers(primera: boolean){
    this.ptestS.getAnswersFromPublishedTests(this.testid).subscribe(
      resp => {
        this.ptestAnswersRecieved(resp)
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
        this.ptestAnswersRecieved(resp)
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

  printDuracion(di: Date | undefined, df: Date | undefined): string{
    if(di==undefined || df == undefined) return "--"
    var d1 = new Date(di)
    var d2 = new Date(df)
    var time = d2.getTime() - d1.getTime()
    var minutes = time/60000
    return minutes.toLocaleString('en-us', {minimumFractionDigits: 2, maximumFractionDigits: 2})
  }

  isModoTestAdmin(): boolean {
    return this.isInAdminTeam || this.test.username == this.getSessionUser().getUsername() || this.getSessionUser().isAdmin()
  }

  clickOrderByPuntuacion(){
    if(this.orderBy == EnumOrderByAnswer.morePuntuacion){
      this.orderBy = EnumOrderByAnswer.lessPuntuacion
    }else if(this.orderBy == EnumOrderByAnswer.lessPuntuacion){
      this.orderBy = EnumOrderByAnswer.newStartDate
    }else{
      this.orderBy = EnumOrderByAnswer.morePuntuacion
    }
    this.getPTestAnswers()
  }

  isMorePuntuacionSelected(): boolean{
    return this.orderBy == EnumOrderByAnswer.morePuntuacion
  }

  isLessPuntuacionSelected(): boolean{
    return this.orderBy == EnumOrderByAnswer.lessPuntuacion
  }

  clickOrderByDuracion(){
    if(this.orderBy == EnumOrderByAnswer.moreDuracion){
      this.orderBy = EnumOrderByAnswer.lessDuracion
    }else if(this.orderBy == EnumOrderByAnswer.lessDuracion){
      this.orderBy = EnumOrderByAnswer.newStartDate
    }else{
      this.orderBy = EnumOrderByAnswer.moreDuracion
    }
    this.getPTestAnswers()
  }

  isMoreDuracionSelected(): boolean{
    return this.orderBy == EnumOrderByAnswer.moreDuracion
  }

  isLessDuracionSelected(): boolean{
    return this.orderBy == EnumOrderByAnswer.lessDuracion
  }
  
  canOrderPuntuacion(): boolean{
    return this.canOrderByPuntuacion
  }

  canOrderDuracion(): boolean{
    return this.canOrderByDuracion
  }
}
