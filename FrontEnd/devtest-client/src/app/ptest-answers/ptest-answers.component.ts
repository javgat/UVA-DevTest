import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptest-answers',
  templateUrl: './ptest-answers.component.html',
  styleUrls: ['./ptest-answers.component.css']
})
export class PtestAnswersComponent extends LoggedInTeacherController implements OnInit {

  testid : number
  routeSub: Subscription
  answers: Answer[]
  test: Test

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.testid = 0
    this.test = new Examen()
    this.answers = []
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getPTest(true)
      }
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
  }

  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0) {
      this.getPTest(true)
    }
  }

  getPTest(primera: boolean) {
    this.userS.getSolvableTestFromUser(this.getSessionUser().getUsername(), this.testid).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
      },
      err => {
        this.handleErrRelog(err, "obtener test publicado", primera, this.getPTest, this)
      }
    )
  }
}
