import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-ptest-u-answers',
  templateUrl: './ptest-u-answers.component.html',
  styleUrls: ['./ptest-u-answers.component.css']
})
export class PtestUAnswersComponent extends LoggedInController implements OnInit {


  testid : number
  routeSub: Subscription
  test: Test
  lookCorregidas: boolean

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.testid = 0
    this.test = new Examen()
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getPTest(true)
      }
    });
    this.lookCorregidas = true
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

  isLookingCorregidas() : boolean{
    return this.lookCorregidas
  }

  isLookingNoCorregidas(): boolean{
    return !this.lookCorregidas
  }

  lookForCorregidas(){
    this.lookCorregidas = true
  }

  lookForNoCorregidas(){
    this.lookCorregidas = false
  }
}
