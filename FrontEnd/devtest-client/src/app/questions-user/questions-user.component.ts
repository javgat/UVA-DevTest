import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Question, QuestionService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-questions-user',
  templateUrl: './questions-user.component.html',
  styleUrls: ['./questions-user.component.css']
})
export class QuestionsUserComponent extends LoggedInTeacherController implements OnInit {

  questions: Question[]
  username: string
  routeSub: Subscription
  includeNonEdit: boolean

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.username = ""
    this.questions = []
    this.includeNonEdit = false
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.borrarMensaje()
      this.getUserQuestions(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
    this.routeSub.unsubscribe()
    this.borrarMensaje()
  }

  getUserQuestions(primera: boolean) {
    if (this.includeNonEdit) {
      this.userS.getQuestionsOfUser(this.username).subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas de un usuario", primera, this.getUserQuestions, this)
      )
    } else {
      this.userS.getEditQuestionsOfUser(this.username).subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas editables de un usuario", primera, this.getUserQuestions, this)
      )
    }
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined) {
    return tipoPrint(tipo, eleUni)
  }

  changeFlexInclude(){
    this.includeNonEdit = !this.includeNonEdit
    this.getUserQuestions(true)
  }

}
