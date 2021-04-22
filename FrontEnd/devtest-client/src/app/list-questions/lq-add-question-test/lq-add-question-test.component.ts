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
  
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected qS: QuestionService, protected tagS: TagService) {
    super(session, router, data, userS, qS, tagS)
    this.scrollable = true
    this.selectAddQuestion = true
  }

  ngOnInit(): void {
  }

  selectQuestion(id: number | undefined){
    this.onQuestionPicked.emit(id);
  }

}
