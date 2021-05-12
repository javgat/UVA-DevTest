import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQuestionsComponent } from '../list-questions.component';

@Component({
  selector: 'app-lq-add-question-test',
  templateUrl: '../list-questions.component.html',
  styleUrls: ['../list-questions.component.css']
})
export class LqAddQuestionTestComponent extends ListQuestionsComponent implements OnInit {

  @Output() onQuestionPicked = new EventEmitter<number>();
  
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, qS: QuestionService, tagS: TagService) {
    super(session, router, data, userS, qS, tagS)
    this.scrollable = true
    this.selectAddQuestion = true
  }

  ngOnInit(): void {
  }

  doHasUserAction(){
    this.getQuestionsFilters()
  }

  selectQuestion(id: number | undefined){
    this.onQuestionPicked.emit(id);
  }

  getQuestionsInclude(primera: boolean) {
    if(this.getSessionUser().isEmpty()) return
    var username = this.getSessionUser().getUsername()
    this.userS.getAvailableQuestionsOfUser(username, this.searchTags, this.likeTitle).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas disponibles para un usuario", primera, this.getQuestionsInclude, this)
    )
  }

  getQuestionsEdit(primera: boolean) {
    if(this.getSessionUser().isEmpty()) return
    var username = this.getSessionUser().getUsername()
    this.userS.getAvailableEditQuestionsOfUser(username, this.searchTags, this.likeTitle).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas editables disponibles para un usuario", primera, this.getQuestionsInclude, this)
    )
  }

}
