import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Team, TeamService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-teams',
  templateUrl: './admin-teams.component.html',
  styleUrls: ['./admin-teams.component.css']
})
export class AdminTeamsComponent extends AdminComponent implements OnInit {

  teams: Team[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService) {
    super(session, router, data, userS, teamS)
    this.teams = []
  }

  ngOnInit(): void {
  }

  getTeams(primera: boolean) {
    this.teamS.getTeams().subscribe(
      resp => {
        this.teams = resp
      },
      err => {
        this.handleErrRelog(err, "Obtener equipos del panel de administración", primera, this.getTeams, this)
      }
    )
  }

  doAdminAction() {
    this.getTeams(true)
  }

}
