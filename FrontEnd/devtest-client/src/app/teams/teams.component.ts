import { Component, OnInit } from '@angular/core';
import { Team, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { SessionUser } from '../shared/app.model';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-teams',
  templateUrl: './teams.component.html',
  styleUrls: ['./teams.component.css']
})
export class TeamsComponent implements OnInit {

  teams : Team[]
  sessionUser : SessionUser

  private sessionUserSubscription : Subscription

  constructor(private session : SessionService, private userS : UserService) {
    this.teams = []
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor =>{
        this.sessionUser = valor
        this.getTeamsOfUser(valor.username)
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy() :void{
    this.sessionUserSubscription.unsubscribe()
  }

  getTeamsOfUser(username : string){
    this.userS.getTeamsOfUser(username).subscribe(
      resp => {
        this.teams = resp
      },
      err =>{
        console.log(err)
      }
    )
  }

}
