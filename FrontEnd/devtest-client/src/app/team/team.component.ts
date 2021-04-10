import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Team, TeamService, User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Equipo, Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.css']
})
export class TeamComponent extends LoggedInController implements OnInit {


  private routeSub: Subscription
  id: string
  equipo: Equipo
  admins: User[]
  miembros: User[]
  addMiembroUsername: string
  usernamePutAdmin: string
  kickingUsername: string
  mightKickUsername: string
  teamEdit: Team

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
    private route: ActivatedRoute, private teamService: TeamService) {
    super(session, router, data, userS)
    this.equipo = new Equipo("", "", false)
    this.admins = []
    this.miembros = []
    this.teamEdit = new Equipo("", "", false)
    this.id = this.addMiembroUsername = this.usernamePutAdmin = this.kickingUsername = this.mightKickUsername = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      if (this.id != "" && this.id != undefined) {
        this.getTeam(true)
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    this.borrarMensaje()
    super.onDestroy()
  }

  getTeam(primera: boolean) {
    this.teamService.getTeam(this.id).subscribe(
      resp => {
        let team = resp
        this.equipo = new Equipo(team.teamname, team.description || "", team.soloProfesores)
        this.getAdmins(true)
        this.getMiembros(true)
      },
      err => {
        this.handleErrRelog(err, "obtener el equipo", primera, this.getTeam, this)
      }
    )
  }

  getAdmins(primera: boolean) {
    this.teamService.getAdmins(this.id).subscribe(
      resp => this.admins = resp,
      err => this.handleErrRelog(err, "obtener admins de equipo", primera, this.getAdmins, this)
    )
  }

  getMiembros(primera: boolean) {
    this.teamService.getMembers(this.id).subscribe(
      resp => this.miembros = resp,
      err => this.handleErrRelog(err, "obtener miembors del equipo", primera, this.getMiembros, this)
    )
  }

  isTeamAdmin(username: string): boolean {
    for (let i = 0; i < this.admins.length; i++) {
      if (this.admins[i].username == username) {
        return true
      }
    }
    return false
  }

  isTeamMiembro(username: string): boolean {
    for (let i = 0; i < this.miembros.length; i++) {
      if (this.miembros[i].username == username) {
        return true
      }
    }
    return false
  }

  checkTeamAdmin(): boolean {
    let username = this.getSessionUser().getUsername()
    return this.isTeamAdmin(username) || this.getSessionUser().isAdmin()
  }

  addMemberSubmit() {
    this.addMember(true)
  }

  addMember(primera: boolean) {
    if (this.isTeamAdmin(this.addMiembroUsername) || this.isTeamMiembro(this.addMiembroUsername)) {
      this.cambiarMensaje(new Mensaje("Ese usuario ya pertenece al equipo", Tipo.ERROR, true))
    } else {
      this.teamService.addMember(this.id, this.addMiembroUsername).subscribe(
        resp => this.getMiembros(true),
        err => this.handleErrRelog(err, "añadir miembro a equipo", primera, this.addMember, this)
      )
    }
  }

  putAdminClick(username: string) {
    this.usernamePutAdmin = username
    this.putAdmin(true)
  }

  putMemberClick(username: string) {
    this.usernamePutAdmin = username
    if (this.admins.length > 1) {
      this.putMember(true)
    } else {
      this.cambiarMensaje(new Mensaje("No se puede dejar al equipo sin administradores", Tipo.ERROR, true))
    }
  }

  putAdmin(primera: boolean) {
    this.teamService.addAdmin(this.id, this.usernamePutAdmin).subscribe(
      resp => {
        this.getMiembros(true)
        this.getAdmins(true)
      },
      err => this.handleErrRelog(err, "otorgar permisos de administracion en un equipo", primera, this.putAdmin, this)
    )
  }

  putMember(primera: boolean) {
    this.teamService.addMember(this.id, this.usernamePutAdmin).subscribe(
      resp => {
        this.getMiembros(true)
        this.getAdmins(true)
      },
      err => this.handleErrRelog(err, "retirar permisos de administracion en un equipo", primera, this.putMember, this)
    )
  }

  mighDeleteClick(username:string){
    this.mightKickUsername = username
  }

  deleteMemberClick(username: string) {
    this.kickingUsername = username
    this.deleteMember(true)
  }

  deleteMember(primera: boolean) {
    this.teamService.deleteUserFromTeam(this.id, this.kickingUsername).subscribe(
      resp=>{
        this.getMiembros(true)
      },
      err => this.handleErrRelog(err, "expulsar miembro de un equipo", primera, this.deleteMember, this)
    )
  }

  updateTeamClick(){
    console.log("updating team...")
    this.updateTeam(true)
  }

  updateTeam(primera: boolean){
    this.teamService.putTeam(this.id, this.teamEdit).subscribe(
      resp => {
        this.router.navigate(['/teams',this.teamEdit.teamname])
      },
      err=> this.handleErrRelog(err, "actualizar equipo", primera, this.updateTeam, this)
    )
  }

}