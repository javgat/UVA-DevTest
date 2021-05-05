import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListAnswersComponent } from '../list-answers.component';

@Component({
  selector: 'app-la-uncorrected',
  templateUrl: '../list-answers.component.html',
  styleUrls: ['../list-answers.component.css']
})
export class LaUncorrectedComponent extends ListAnswersComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, ptestS: PublishedTestService, route: ActivatedRoute) {
    super(session, router, data, userS, ptestS, route)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

  getPTestAllAnswers(primera: boolean){
    this.ptestS.getUncorrectedAnswersFromPublishedTests(this.testid).subscribe(
      resp => {
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas no corregidas de test", primera, this.getPTestAllAnswers, this)
      }
    )
  }

  getPTestAnswersFromUser(primera: boolean){
    if(this.likeUsername==undefined) return
    this.userS.getUncorrectedAnswersFromUserAnsweredTest(this.likeUsername, this.testid).subscribe(
      resp =>{
        this.answers = resp
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas no corregidas de test y usuario", primera, this.getPTestAnswersFromUser, this)
      }
    )
  }

}