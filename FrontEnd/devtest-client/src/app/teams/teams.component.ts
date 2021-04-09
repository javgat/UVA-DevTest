import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, Team, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { SessionUser } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-teams',
  templateUrl: './teams.component.html',
  styleUrls: ['./teams.component.css']
})
export class TeamsComponent extends LoggedInController implements OnInit {

  teams: Team[]

  private sUserSub: Subscription = new Subscription;
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
    this.teams = []
  }

  ngOnInit(): void {
    this.getTeams(true)
  }

  ngOnDestroy(): void {
    this.sUserSub.unsubscribe()
    super.onDestroy()
  }

  getTeams(primera: boolean) {
    this.sUserSub.unsubscribe()
    this.sUserSub = this.session.sessionUser.subscribe(
      valor => {
        if(!valor.isEmpty()){
          this.getTeamsOfUser(valor.username, primera)
        }
      }
    )
  }

  getTeamsOfUser(username: string, primera: boolean) {
    this.userS.getTeamsOfUser(username).subscribe(
      resp => {
        this.teams = resp
      },
      err => {
        this.handleErrRelog(err, "obtener equipos de usuario", primera, this.getTeams, this)
      }
    )
  }

}
