import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';
import { tipoPrint } from '../shared/app.model'

@Component({
  selector: 'app-questions-shared-user',
  templateUrl: './questions-shared-user.component.html',
  styleUrls: ['./questions-shared-user.component.css']
})
export class QuestionsSharedUserComponent extends LoggedInTeacherController implements OnInit {

  questions: Question[]
  username: string
  routeSub: Subscription
  
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.username = ""
    this.questions = []
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
      this.getUserSharedQuestions(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.routeSub.unsubscribe()
    this.borrarMensaje()
  }

  getUserSharedQuestions(primera: boolean){
    this.userS.getSharedQuestionsOfUser(this.username).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas compartidas con el usuario", primera, this.getUserSharedQuestions, this)
    )
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined){
    return tipoPrint(tipo, eleUni)
  }

}
