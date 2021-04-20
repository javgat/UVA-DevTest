import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, Team, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ProfileComponent } from '../profile.component';

@Component({
  selector: 'app-profile-teams',
  templateUrl: './profile-teams.component.html',
  styleUrls: ['./profile-teams.component.css']
})
export class ProfileTeamsComponent extends ProfileComponent implements OnInit {

  teams: Team[]
  constructor(session: SessionService, router: Router, route: ActivatedRoute,
    userS: UserService, data: DataService, authService: AuthService) {
      super(session, router, route, userS, data, authService)
      this.teams = []
  }

  doProfileAction(): void{
    this.getTeamsOfUser(true)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

  getTeamsOfUser(primera: boolean){
    this.userS.getTeamsOfUser(this.id).subscribe(
      resp => this.teams = resp,
      err => this.handleErrRelog(err, "obtener equipos de un usuario", primera, this.getTeamsOfUser, this)
    )
  }

}
