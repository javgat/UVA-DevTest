import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, QuestionService, Team, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Pregunta } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-question-teams',
  templateUrl: './question-teams.component.html',
  styleUrls: ['./question-teams.component.css']
})
export class QuestionTeamsComponent extends LoggedInTeacherController implements OnInit {

  teams : Team[]
  userTeams: Team[]
  question: Question
  routeSub: Subscription
  id: number
  addTeamTeamname: string
  kickingTeamname: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.question = new Pregunta(undefined, "", "", 0, false, true, "", undefined, undefined, Question.TipoPreguntaEnum.String, undefined)
    this.teams = []
    this.userTeams = []
    this.id = 0
    this.addTeamTeamname = ""
    this.kickingTeamname = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      this.getPregunta(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.borrarMensaje()
    this.routeSub.unsubscribe()
  }

  doHasUserAction(){
    this.getUserTeams(true)
  }

  getUserTeams(primera: boolean){
    this.userS.getTeamsOfUser(this.getSessionUser().getUsername()).subscribe(
      resp =>{
        this.userTeams = resp
      },
      err => this.handleErrRelog(err, "obtener equipos de usuario", primera, this.getUserTeams, this)
    )
  }

  getPregunta(primera: boolean) {
    this.qS.getQuestion(this.id).subscribe(
      resp => {
        this.question = new Pregunta(resp.id, resp.title, resp.question, resp.estimatedTime,
          resp.autoCorrect, resp.editable, resp.username, resp.eleccionUnica, resp.solucion,
          resp.tipoPregunta, resp.valorFinal)
        this.getTeamsPregunta(true)
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
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

  checkPermisosAdministracion(): boolean {
    return this.question.editable && (this.getSessionUser().isAdmin() || this.isPermisosAdministracion())
  }

  addTeamSubmit(){
    this.addTeam(true)
  }

  addTeam(primera: boolean){
    this.qS.addTeamToQuestion(this.addTeamTeamname, this.id).subscribe(
      resp => this.getTeamsPregunta(true),
      err => this.handleErrRelog(err, "añadir equipo a una pregunta", primera, this.addTeam, this)
    )
  }

  kickTeam(teamname: string){
    this.kickingTeamname = teamname
    this.kickT(true)
  }

  kickT(primera: boolean){
    this.qS.removeTeamToQuestion(this.kickingTeamname, this.id).subscribe(
      resp => this.getTeamsPregunta(true),
      err => this.handleErrRelog(err, "eliminar equipo de una pregunta", primera, this.kickT, this)
    )
  }

}
