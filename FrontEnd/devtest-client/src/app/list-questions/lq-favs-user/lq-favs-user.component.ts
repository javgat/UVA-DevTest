import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQuestionsComponent } from '../list-questions.component';

@Component({
  selector: 'app-lq-favs-user',
  templateUrl: '../list-questions.component.html',
  styleUrls: ['../list-questions.component.css']
})
export class LqFavsUserComponent extends ListQuestionsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, qS: QuestionService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, qS, tagS)
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['username']
      this.borrarMensaje()
      this.getQuestionsFilters()
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
    this.routeSub.unsubscribe()
  }

  getQuestionsInclude(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getFavoriteQuestions(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas favoritas de un usuario", primera, this.getQuestionsInclude, this)
    )
  }

  getQuestionsEdit(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getFavoriteEditQuestions(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas favoritas editables de un usuario", primera, this.getQuestionsEdit, this)
    )
  }

}