import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Test, UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-test-create',
  templateUrl: './test-create.component.html',
  styleUrls: ['./test-create.component.css']
})
export class TestCreateComponent extends LoggedInTeacherController implements OnInit {

  test: Test
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
    this.test = {
      title: "",
      description: "",
      accesoPublico: false,
      editable: true,
      username: "",
      maxMinutes: 0,
      accesoPublicoNoPublicado: true
    }
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  doHasUserAction(){
    this.test.username = this.getSessionUser().getUsername()
  }

  testSubmit(){
    this.sendTest(true)
  }

  sendTest(primera: boolean){
    this.userS.postTest(this.getSessionUser().getUsername(), this.test).subscribe(
      resp=> this.router.navigate(['/et',resp.id]),
      err=> this.handleErrRelog(err, "crear nuevo test", primera, this.sendTest, this)
    )
  }

}
