import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, Test, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-answers-user',
  templateUrl: './answers-user.component.html',
  styleUrls: ['./answers-user.component.css']
})
export class AnswersUserComponent extends LoggedInController implements OnInit {

  username : string
  routeSub: Subscription
  answers: Answer[]
  test: Test
  lookCorregidas: boolean

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.username = ""
    this.test = new Examen()
    this.answers = []
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
    });
    this.lookCorregidas = true
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.routeSub.unsubscribe()
    super.onDestroy()
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
