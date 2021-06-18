import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TeamService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Equipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-team-resources',
  templateUrl: './team-resources.component.html',
  styleUrls: ['./team-resources.component.css']
})
export class TeamResourcesComponent extends LoggedInController implements OnInit {
  chosen: ResourceChosen
  routeSub: Subscription
  id: string
  equipo: Equipo
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private teamService: TeamService) {
    super(session, router, data, userS);
    this.id = ""
    this.equipo = new Equipo("", "", false)
    var url = router.url
    var regexpInvitedTests = new RegExp('^.+\/'+ResourceChosen.invitedTests)
    var regexpAdminTests = new RegExp('^.+\/'+ResourceChosen.adminTests)
    if(regexpInvitedTests.test(url)){
      this.chosen = ResourceChosen.invitedTests
    }else if(regexpAdminTests.test(url)){
      this.chosen = ResourceChosen.adminTests
    }else{
      this.chosen = ResourceChosen.adminQuestions
    }
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      if (this.id != "" && this.id != undefined) {
        this.getTeam(true)
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
  }

  getTeam(primera: boolean) {
    this.teamService.getTeam(this.id).subscribe(
      resp => {
        let team = resp
        this.equipo = new Equipo(team.teamname, team.description || "", team.soloProfesores)
      },
      err => {
        this.handleErrRelog(err, "obtener el equipo", primera, this.getTeam, this)
      }
    )
  }

  isResourceInvitedTests(): boolean{
    return this.chosen == ResourceChosen.invitedTests
  }

  isResourceAdminTests(): boolean{
    return this.chosen == ResourceChosen.adminTests
  }

  isResourceAdminQuestions(): boolean{
    return this.chosen == ResourceChosen.adminQuestions
  }

}

enum ResourceChosen{
  adminTests = "adminTests",
  adminQuestions = "adminQuestions",
  invitedTests = "invitedTests"
}