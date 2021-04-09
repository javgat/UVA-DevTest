import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Team, User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Equipo, SessionUser } from '../shared/app.model';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-create-team',
  templateUrl: './create-team.component.html',
  styleUrls: ['./create-team.component.css']
})
export class CreateTeamComponent implements OnInit {

  sessionUserSubscription: Subscription
  sessionUser: SessionUser
  team: Team

  constructor(private session: SessionService, private router: Router, private userS: UserService) {
    this.team = {
      teamname: "",
      description: "",
      soloProfesores: false
    }
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor => {
        this.sessionUser = valor
        if (this.sessionUser.getRol() == User.RolEnum.Estudiante) {
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.sessionUserSubscription.unsubscribe()
  }

  teamSubmit(): void {
    this.userS.postTeam(this.sessionUser.getUsername(), this.team).subscribe(

    )
  }

}
