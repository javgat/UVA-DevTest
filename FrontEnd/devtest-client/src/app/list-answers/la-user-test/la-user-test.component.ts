import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListAnswersComponent } from '../list-answers.component';

@Component({
  selector: 'app-la-user-test',
  templateUrl: '../list-answers.component.html',
  styleUrls: ['../list-answers.component.css']
})
export class LaUserTestComponent extends ListAnswersComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, ptestS: PublishedTestService, route: ActivatedRoute) {
    super(session, router, data, userS, ptestS, route)
    this.buscarUsuario = false
    this.canOrderByDuracion = true
    this.canOrderByPuntuacion = true
    if (this.getSessionUser().getUsername() != undefined && this.getSessionUser().getUsername() != "") {
      this.likeUsername = this.getSessionUser().getUsername()
      this.getPTestAllAnswers(true)
    }
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.ngOnDestroy()
  }

  doHasUserAction() {
    if (this.testid != undefined && this.testid != 0) {
      this.likeUsername = this.getSessionUser().getUsername()
      this.getPTestAllAnswers(true)
    }
  }

  getPTestAllAnswers(primera: boolean) {
    if (this.likeUsername == undefined || this.testid==undefined) return
    this.userS.getAnswersFromUserAnsweredTest(this.likeUsername, this.testid, this.orderBy).subscribe(
      resp => {
        this.ptestAnswersRecieved(resp)
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas corregidas de test y usuario", primera, this.getPTestAllAnswers, this)
      }
    )
  }

  getPTestAnswersFromUser(primera: boolean) {
    this.getPTestAllAnswers(true)
  }

}