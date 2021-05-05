import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Answer, PublishedTestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-answers',
  templateUrl: './list-answers.component.html',
  styleUrls: ['./list-answers.component.css']
})
export class ListAnswersComponent extends LoggedInController implements OnInit {

  testid: number
  answers: Answer[]
  routeSub: Subscription
  editLikeUsername: string
  likeUsername?: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private ptestS : PublishedTestService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.testid = 0
    this.answers = []
    this.editLikeUsername = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.testid = params['testid']
      this.borrarMensaje()
      if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
        this.getPTestAnswers()
      }
    });
  }

  ngOnInit(): void {
  }


  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0) {
      this.getPTestAnswers()
    }
  }

  getPTestAnswers(){
    if(this.likeUsername == undefined)
      this.getPTestAllAnswers(true)
    else
      this.getPTestAnswersFromUser(true)
  }

  getPTestAllAnswers(primera: boolean){
    this.ptestS.getAnswersFromPublishedTests(this.testid).subscribe(
      resp => {
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas de test", primera, this.getPTestAllAnswers, this)
      }
    )
  }

  getPTestAnswersFromUser(primera: boolean){
    if(this.likeUsername==undefined) return
    this.userS.getAnswersFromUserAnsweredTest(this.likeUsername, this.testid).subscribe(
      resp =>{
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas de test y usuario", primera, this.getPTestAnswersFromUser, this)
      }
    )
  }

  clickSearchUsername(){
    this.likeUsername = this.editLikeUsername
    this.getPTestAnswers()
  }

  clickBorrarUsername(){
    this.likeUsername = undefined
    this.getPTestAnswers()
  }

}
