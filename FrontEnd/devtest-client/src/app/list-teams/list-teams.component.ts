import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Team, TeamService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { LoggedInController } from '../shared/app.controller';

@Component({
  selector: 'app-list-teams',
  templateUrl: './list-teams.component.html',
  styleUrls: ['./list-teams.component.css']
})
export class ListTeamsComponent extends LoggedInController implements OnInit {

  teams: Team[]
  editLikeTeamname: string
  likeTeamname: string | undefined
  searchActive: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected teamS: TeamService) {
    super(session, router, data, userS)
    this.editLikeTeamname = ""
    this.searchActive = false
    this.teams = []
    this.getTeamsFilter()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
  }

  getTeamsFilter(){}

  clickSearchTeamname(){
    this.likeTeamname = this.editLikeTeamname
    this.getTeamsFilter()
  }

  clickBorrarTeamname(){
    this.likeTeamname = undefined
    this.editLikeTeamname = ""
    this.getTeamsFilter()
  }


  activateSearch(){
    this.searchActive=true
  }

  isSearchActive(): boolean{
    return this.searchActive
  }

}
