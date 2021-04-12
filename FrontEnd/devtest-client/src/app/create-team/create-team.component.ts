import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Team, User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-create-team',
  templateUrl: './create-team.component.html',
  styleUrls: ['./create-team.component.css']
})
export class CreateTeamComponent extends LoggedInController implements OnInit {

  team: Team

  private sUserSub : Subscription

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
    this.team = {
      teamname: "",
      description: "",
      soloProfesores: false
    }
    this.sUserSub = this.session.sessionUser.subscribe(
      valor => {
        if (!valor.isEmpty() && valor.getRol() == User.RolEnum.Estudiante) {
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.sUserSub.unsubscribe()
    super.onDestroy()
  }

  onSelectSoloProf(valor: boolean) : void{
    this.team.soloProfesores = valor
  }

  teamSubmit(): void {
    this.teamPost(true)
  }

  teamPost(primera: boolean){
    this.userS.postTeam(this.getSessionUser().getUsername(), this.team).subscribe(
      resp=>{
        console.log("Equipo creado con exito")
        this.router.navigate(['/teams', this.team.teamname])
      },
      err =>{
        if(err.status == 409){
          this.cambiarMensaje(new Mensaje("Ya existe un equipo con ese nombre de equipo", Tipo.ERROR, true))
        }else
          this.handleErrRelog(err, "crear nuevo equipo", primera, this.teamPost, this)
      }
    )
  }


}
