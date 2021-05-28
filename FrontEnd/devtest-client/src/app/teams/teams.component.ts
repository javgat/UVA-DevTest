import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
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

  username: string
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.teams = []
    this.username = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
      this.getTeams(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
  }

  getTeams(primera: boolean) {
    if(this.username == "" || this.username == undefined) return
    this.getTeamsOfUser(this.username, primera)
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
