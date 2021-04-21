import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-questions',
  templateUrl: './questions.component.html',
  styleUrls: ['./questions.component.css']
})
export class QuestionsComponent extends LoggedInTeacherController implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private tagS: TagService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

}
