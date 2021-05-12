import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, TagService, TeamService, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListTestsComponent } from '../list-tests.component';

@Component({
  selector: 'app-lt-admin-team',
  templateUrl: '../list-tests.component.html',
  styleUrls: ['../list-tests.component.css']
})
export class LtAdminTeamComponent extends ListTestsComponent  implements OnInit {

  @Input()
  teamname: string = "";
  
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, tS: TestService, tagS: TagService, teamS: TeamService) {
    super(session, router, data, userS, tS, tagS, teamS)
  }

  ngOnInit(): void {
    this.getTestsFilters()
  }

  getTestsInclude(primera: boolean) {
    //No hay getTestsFromTeam en la API que tenga filtros
    this.getTestsEdit(true)
  }

  // EN published tests testEdit es el testPublished
  getTestsEdit(primera: boolean) {
    if(this.teamname=="" || this.teamname==undefined || this.teamS == undefined) return
    this.teamS.getTestsFromTeam(this.teamname, this.searchTags, this.likeTitle).subscribe(
      resp => this.tests = resp,
      err => this.handleErrRelog(err, "obtener tests que el equipo administra", primera, this.getTestsEdit, this)
    )
  }
}
