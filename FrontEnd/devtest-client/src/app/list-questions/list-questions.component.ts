import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Question, QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-questions',
  templateUrl: './list-questions.component.html',
  styleUrls: ['./list-questions.component.css']
})
export class ListQuestionsComponent extends LoggedInController implements OnInit {

  newSearchTag: string
  questions: Question[]
  searchTags: string[][]
  includeNonEdit: boolean
  orOperation: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private tagS: TagService) {
    super(session, router, data, userS)
    this.includeNonEdit = false
    this.searchTags = []
    this.questions = []
    this.newSearchTag = ""
    this.orOperation = true
    this.getQuestionsFilters()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  getQuestionsInclude(primera: boolean) {
    this.qS.getQuestions(this.searchTags).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas", primera, this.getQuestionsInclude, this)
    )
  }

  getQuestionsEdit(primera: boolean) {
    this.qS.getEditQuestions(this.searchTags).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas no publicadas", primera, this.getQuestionsEdit, this)
    )
  }

  getQuestionsFilters() {
    if (this.includeNonEdit) {
      this.getQuestionsInclude(true)
    } else {
      this.getQuestionsEdit(true)
    }
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined) {
    return tipoPrint(tipo, eleUni)
  }

  changeFlexInclude() {
    this.includeNonEdit = !this.includeNonEdit
    this.getQuestionsFilters()
  }

  addSearchTagOr() {
    var orTags = [this.newSearchTag]
    this.searchTags.push(orTags)
    this.newSearchTag = ""
  }

  addSearchTagSubmit() {
    if (this.orOperation) {
      this.addSearchTagOr()
    }
    this.getQuestionsFilters()
  }

  deleteSearchTagOr(deleteTag: string){
    this.searchTags.forEach(element => {
      element.forEach((subElement, index) => {
        if (subElement == deleteTag) element.splice(index, 1)
      })
    });
  }

  deleteSearchTagClick(deleteTag: string) {
    if(this.orOperation){
      this.deleteSearchTagOr(deleteTag)
    }
    this.getQuestionsFilters()
  }


}
