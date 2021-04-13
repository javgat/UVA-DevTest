import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Question, QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-questions',
  templateUrl: './questions.component.html',
  styleUrls: ['./questions.component.css']
})
export class QuestionsComponent extends LoggedInTeacherController implements OnInit {

  questions: Question[]
  searchTag: string
  includeNonEdit: boolean
  tagPressed: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private tagS: TagService) {
    super(session, router, data, userS)
    this.includeNonEdit = false
    this.tagPressed = false
    this.searchTag = ""
    this.questions = []
    this.getQuestions(true)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  getQuestions(primera: boolean) {
    if (this.includeNonEdit) {
      this.qS.getQuestions().subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas", primera, this.getQuestions, this)
      )
    } else {
      this.qS.getEditQuestions().subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas editables", primera, this.getQuestions, this)
      )
    }
  }

  getQuestionsWithTagSubmit() {
    this.tagPressed=true
    this.getQuestionsWithTag(true)
  }

  getQuestionsWithTag(primera: boolean) {
    if (this.includeNonEdit) {
      this.tagS.getQuestionsFromTag(this.searchTag).subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas por etiqueta", primera, this.getQuestionsWithTag, this)
      )
    } else {
      this.tagS.getEditQuestionsFromTag(this.searchTag).subscribe(
        resp => this.questions = resp,
        err => this.handleErrRelog(err, "obtener preguntas no publicadas por etiqueta", primera, this.getQuestionsWithTag, this)
      )
    }
  }


  tipoPrint(tipo: string, eleUni: boolean | undefined) {
    return tipoPrint(tipo, eleUni)
  }

  changeFlexInclude() {
    this.includeNonEdit = !this.includeNonEdit
    if(this.tagPressed){
      this.getQuestionsWithTag(true)
    }else{
      this.getQuestions(true)
    }
  }

}
