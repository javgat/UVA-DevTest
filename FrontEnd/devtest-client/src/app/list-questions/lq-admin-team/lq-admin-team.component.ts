import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { QuestionService, TagService, TeamService, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQuestionsComponent } from '../list-questions.component';

@Component({
  selector: 'app-lq-admin-team',
  templateUrl: '../list-questions.component.html',
  styleUrls: ['../list-questions.component.css']
})
export class LqAdminTeamComponent extends ListQuestionsComponent implements OnInit {

  @Input()
  teamname: string = "";

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, 
    qS: QuestionService, tagS: TagService, teamS: TeamService) {
    super(session, router, data, userS, qS, tagS, teamS)
    this.hideSwitchInclude=true
  }

  ngOnInit(): void {
    this.getQuestionsFilters()
  }

  getQuestionsInclude(primera: boolean) {
    this.getQuestionsEdit(primera)
  }

  getQuestionsEdit(primera: boolean) {
    if(this.teamname=="" || this.teamname==undefined || this.teamS == undefined) return
    this.teamS.getQuestionsFromTeam(this.teamname, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas de un equipo", primera, this.getQuestionsInclude, this)
    )
  }

}
