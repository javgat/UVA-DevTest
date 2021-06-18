import { Component, Input, OnInit } from '@angular/core';
import { ListTeamsComponent } from '../list-teams.component';
import { Router } from '@angular/router';
import { Team, TeamService, User, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { Subscription } from 'rxjs';
@Component({
  selector: 'app-ltm-user',
  templateUrl: '../list-teams.component.html',
  styleUrls: ['../list-teams.component.css']
})
export class LtmUserComponent extends ListTeamsComponent implements OnInit {

  @Input() username: string | undefined;

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService) {
    super(session, router, data, userS, teamS)
    this.getTeamsFilter()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

  getTeamsFilter(){
    this.getTeams(true)
  }

  getTeams(primera: boolean) {
    if (this.username == "" || this.username == undefined) return
    this.getTeamsOfUser(this.username, primera)
  }

  getTeamsOfUser(username: string, primera: boolean) {
    this.userS.getTeamsOfUser(username, undefined, this.likeTeamname).subscribe(
      resp => {
        this.teams = resp
      },
      err => {
        this.handleErrRelog(err, "obtener equipos de usuario", primera, this.getTeams, this)
      }
    )
  }

  ngOnChanges() {
    this.getTeamsFilter()
  }   

}
