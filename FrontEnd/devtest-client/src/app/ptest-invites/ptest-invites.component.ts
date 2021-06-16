import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EmailUser, Message, PublishedTestService, Team, TeamService, Test, TestService, User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen, Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptest-invites',
  templateUrl: './ptest-invites.component.html',
  styleUrls: ['./ptest-invites.component.css']
})
export class PtestInvitesComponent extends LoggedInTeacherController implements OnInit {

  teams : Team[]
  test: Test
  iusers: User[]
  routeSub: Subscription
  id: number
  addTeamTeamname: string
  addUserUsername: string
  kickingTeamname: string
  kickingUsername: string
  isInAdminTeam: boolean
  lookingForTeams: boolean
  customMessageNotification: string
  enviaMensaje: boolean
  createUser: boolean
  checkedSendEmail: boolean
  autousers: User[]
  autoteams: Team[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, 
      private tS: TestService, private route: ActivatedRoute, private ptS: PublishedTestService, private teamS: TeamService) {
    super(session, router, data, userS)
    this.test = new Examen()
    this.teams = []
    this.iusers = []
    this.autoteams = []
    this.autousers = []
    this.id = 0
    this.isInAdminTeam = false
    this.addTeamTeamname = ""
    this.kickingTeamname = ""
    this.kickingUsername = ""
    this.addUserUsername = ""
    this.customMessageNotification = ""
    this.lookingForTeams = true
    this.enviaMensaje = false
    this.createUser = false
    this.checkedSendEmail = true
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      this.getTest(true)
    });
    this.changeGetAutoTeams()
    this.changeGetAutoUsers()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.borrarMensaje()
    this.routeSub.unsubscribe()
  }

  doHasUserAction() {
    if (this.id!=undefined && this.id != 0)
      this.getIsInAdminTeam(true)
  }

  getTest(primera: boolean) {
    this.tS.getTest(this.id).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.getTeamsTest(true)
        this.getUsersTest(true)
        if (!this.getSessionUser().isEmpty())
          this.getIsInAdminTeam(true)
      },
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  isLookingTeams(): boolean{
    return this.lookingForTeams
  }

  isLookingUsers(): boolean{
    return !this.lookingForTeams
  }

  lookForTeams(){
    this.lookingForTeams = true
  }

  lookForUsers(){
    this.lookingForTeams = false
  }

  getTeamsTest(primera: boolean){
    this.ptS.getTeamsFromPublishedTest(this.id).subscribe(
      resp =>{
        this.teams = resp
      },
      err => this.handleErrRelog(err, "obtener equipos invitados al test", primera, this.getTeamsTest, this)
    )
  }

  getUsersTest(primera: boolean){
    this.ptS.getUsersFromPublishedTest(this.id).subscribe(
      resp=> this.iusers=resp,
      err => this.handleErrRelog(err, "obtener usuarios invitados al test", primera, this.getUsersTest, this)
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

  isPermisosAdministracion() : boolean{
    return (this.getSessionUser().getUsername() == this.test.username) || this.isInAdminTeam
  }

  checkPermisosAdministracion(): boolean {
    return (this.getSessionUser().isAdmin() || this.isPermisosAdministracion())
  }

  inviteTeamSubmit(){
    if(!this.enviaMensaje)
      this.inviteTeam(true)
    else
      this.inviteTeamMessage(true)
  }

  inviteTeam(primera: boolean){
    this.ptS.inviteTeamToPublishedTest(this.addTeamTeamname, this.id).subscribe(
      resp =>{
        this.cambiarMensaje(new Mensaje("Equipo invitado con éxito", Tipo.SUCCESS, true))
        this.getTeamsTest(true)
      },
      err => this.handleErrRelog(err, "invitar equipo a realizar un test", primera, this.inviteTeam, this)
    )
  }

  inviteTeamMessage(primera: boolean){
    var message: Message
    message = {
      body: this.customMessageNotification,
      sendEmail: this.checkedSendEmail
    }
    this.ptS.inviteTeamToPublishedTest(this.addTeamTeamname, this.id, message).subscribe(
      resp =>{
        this.cambiarMensaje(new Mensaje("Equipo invitado con éxito", Tipo.SUCCESS, true))
        this.getTeamsTest(true)
      },
      err => this.handleErrRelog(err, "invitar equipo a realizar un test con mensaje personalizado", primera, this.inviteTeamMessage, this)
    )
  }

  inviteUserSubmit(){
    if(!this.enviaMensaje)
      this.inviteUser(true)
    else
      this.inviteUserMessage(true)
  }

  inviteUser(primera: boolean){
    this.inviteUserMessage(primera)
  }

  inviteUserMessage(primera: boolean){
    var message: Message
    message = {
      body: this.customMessageNotification,
      sendEmail: this.checkedSendEmail
    }
    this.ptS.inviteUserToPublishedTest(this.addUserUsername, this.id, message).subscribe(
      resp=>{
        this.cambiarMensaje(new Mensaje("Usuario invitado con éxito", Tipo.SUCCESS, true))
        this.getUsersTest(true)
      },
      err => this.handleErrRelog(err, "invitar usuario a realizar un test con mensaje personalizado", primera, this.inviteUser, this)
    )
  }

  uninviteTeamClick(teamname: string){
    this.kickingTeamname = teamname
    this.uninviteTeam(true)
  }

  uninviteTeam(primera: boolean){
    this.ptS.removeTeamToPublishedTest(this.kickingTeamname, this.id).subscribe(
      resp=>{
        this.cambiarMensaje(new Mensaje("Equipo expulsado del test con éxito", Tipo.SUCCESS, true))
        this.getTeamsTest(true)
      },
      err => this.handleErrRelog(err, "expulsar a un equipo de realizar un test", primera, this.uninviteTeam, this)
    )
  }

  uninviteUserClick(username: string){
    this.kickingUsername = username
    this.uninviteUser(true)
  }

  uninviteUser(primera: boolean){
    this.ptS.removeUserToPublishedTest(this.kickingUsername, this.id).subscribe(
      resp => {
        this.cambiarMensaje(new Mensaje("Usuario expulsado del test con éxito", Tipo.SUCCESS, true))
        this.getUsersTest(true)
      },
      err => this.handleErrRelog(err, "expulsar a un usuario de realizar un test", primera, this.uninviteUser, this)
    )
  }

  changeEnviaMensaje(){
    this.enviaMensaje = true
  }

  changeNoEnviaMensaje(){
    this.enviaMensaje = false
  }

  changeNotAddUser(){
    this.createUser = false
  }

  changeAddUser(){
    this.createUser = true
  }

  createInviteUserSubmit(){
    this.createInviteUser(true)
  }

  createInviteUser(primera: boolean){
    var eu: EmailUser
    eu = {
      email: this.addUserUsername
    }
    this.userS.postEmailUser(eu).subscribe(
      resp =>{
        this.inviteUserSubmit()
        this.cambiarMensaje(new Mensaje("Usuario creado con éxito.", Tipo.SUCCESS, true))
      },
      err => this.handleErrRelog(err, "crear usuario para añadir a test", primera, this.createInviteUser, this)
    )
  }

  changeFlexSendEmail(){
    this.checkedSendEmail = !this.checkedSendEmail
  }

  changeGetAutoUsers(){
    this.getAutoUsers(true)
  }

  changeGetAutoTeams(){
    this.getAutoTeams(true)
  }

  getAutoUsers(primera: boolean){
    this.userS.getUsers(undefined, this.addUserUsername, undefined, undefined, 20).subscribe(
      resp=>{
        this.autousers=resp
      },
      err => this.handleErrRelog(err, "obtener usuarios que empiezan por ese username", primera, this.getAutoUsers, this)
    )
  }

  getAutoTeams(primera: boolean){
    this.teamS.getTeams(this.addTeamTeamname, undefined, 20).subscribe(
      resp=>{
        this.autoteams=resp
      },
      err => this.handleErrRelog(err, "obtener equipos que empiezan por ese teamname", primera, this.getAutoTeams, this)
    )
  }

}
