import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Test, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-test-ptests',
  templateUrl: './test-ptests.component.html',
  styleUrls: ['./test-ptests.component.css']
})
export class TestPTestsComponent extends LoggedInTeacherController implements OnInit {

  ptests: Test[]
  test: Test
  routeSub: Subscription
  id: number
  addTeamTeamname: string
  kickingTeamname: string
  isInAdminTeam: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private tS: TestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.test = new Examen()
    this.ptests = []
    this.id = 0
    this.isInAdminTeam = false
    this.addTeamTeamname = ""
    this.kickingTeamname = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      this.getTest(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.borrarMensaje()
    this.routeSub.unsubscribe()
  }

  doHasUserAction() {
    if (this.id!=undefined && this.id != 0)
      this.getIsInAdminTeam(true)
  }

  getTest(primera: boolean) {
    this.tS.getTest(this.id).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.getTestPTests(true)
        if (!this.getSessionUser().isEmpty())
          this.getIsInAdminTeam(true)
      },
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if(err.status!=410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  isPermisosAdministracion() : boolean{
    return (this.getSessionUser().getUsername() == this.test.username) || this.isInAdminTeam
  }

  checkPermisosAdministracion(): boolean {
    return this.test.editable && (this.getSessionUser().isAdmin() || this.isPermisosAdministracion())
  }

  getTestPTests(primera: boolean) {
    this.tS.getPublishedTestsFromTest(this.id).subscribe(
      resp => this.ptests = resp,
      err => this.handleErrRelog(err, "obtener tests publicados del test", primera, this.getTestPTests, this)
    )
  }

  printDate(d: Date | undefined): string{
    if(d==undefined) return ""
    var date = new Date(d)
    return date.toLocaleString()
  }

}
