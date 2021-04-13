import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Test, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';
@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css']
})
export class TestComponent extends LoggedInTeacherController implements OnInit {

  routeSub: Subscription
  id: number
  test: Test
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute, private testS: TestService) {
    super(session, router, data, userS)
    this.id=0
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      this.getTest(true)
    });
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
  }

  getTest(primera: boolean){
    this.testS.getTest(this.id).subscribe(
      resp => this.test = resp,
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  checkPermisosEdicion(): boolean{
    return true
  }

}
