import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListAnswersComponent } from '../list-answers.component';

@Component({
  selector: 'app-la-u-uncorrected',
  templateUrl: '../list-answers.component.html',
  styleUrls: ['../list-answers.component.css']
})
export class LaUUncorrectedComponent extends ListAnswersComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, ptestS: PublishedTestService, route: ActivatedRoute) {
    super(session, router, data, userS, ptestS, route)
    this.buscarUsuario = false
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.ngOnDestroy()
  }

  doHasUserAction() {
    this.likeUsername = this.getSessionUser().getUsername()
    this.getPTestAllAnswers(true)
  }

  getPTestAllAnswers(primera: boolean) {
    if (this.likeUsername == undefined) return
    this.userS.getUncorrectedAnswersFromUser(this.likeUsername).subscribe(
      resp => {
        this.ptestAnswersRecieved(resp)
      },
      err => {
        this.handleErrRelog(err, "obtener respuestas no corregidas de usuario", primera, this.getPTestAllAnswers, this)
      }
    )
  }

  getPTestAnswersFromUser(primera: boolean) {
    this.getPTestAllAnswers(true)
  }

}
