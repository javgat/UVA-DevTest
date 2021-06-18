import { Component, OnInit } from '@angular/core';
import { ListTeamsComponent } from '../list-teams.component';
import { Router } from '@angular/router';
import { Team, TeamService, User, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-lt-admin',
  templateUrl: '../list-teams.component.html',
  styleUrls: ['../list-teams.component.css']
})
export class LtAdminComponent extends ListTeamsComponent implements OnInit {

  private userSub: Subscription

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService) {
    super(session, router, data, userS, teamS)
    this.userSub = this.session.sessionUser.subscribe(
      valor => {
        if (valor.getRol() == User.RolEnum.Administrador) {
          this.getTeamsFilter()
        } else if (!valor.isEmpty()) {
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.userSub.unsubscribe()
    super.ngOnDestroy()
  }

  getTeamsFilter(){
    this.getTeams(true)
  }


  getTeams(primera: boolean) {
    this.teamS.getTeams(undefined, this.likeTeamname).subscribe(
      resp => {
        this.teams = resp
      },
      err => {
        this.handleErrRelog(err, "Obtener equipos del panel de administración", primera, this.getTeams, this)
      }
    )
  }

  
}
