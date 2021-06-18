import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQuestionsComponent } from '../list-questions.component';
@Component({
  selector: 'app-lq-shared-user',
  templateUrl: '../list-questions.component.html',
  styleUrls: ['../list-questions.component.css']
})
export class LqSharedUserComponent  extends ListQuestionsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, qS: QuestionService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, qS, tagS)
    this.id=""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['username']
      this.borrarMensaje()
      this.getQuestionsFilters()
    });
    this.hideSwitchInclude = true
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
    this.routeSub.unsubscribe()
  }

  getQuestionsInclude(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getSharedQuestionsOfUser(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas compartidas con un usuario", primera, this.getQuestionsInclude, this)
    )
  }

  getQuestionsEdit(primera: boolean) {
    //No hay getSharedEditQuestionsOfUser en la API
    if(this.id==undefined) return
    this.userS.getSharedQuestionsOfUser(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas compartidas con un usuario", primera, this.getQuestionsEdit, this)
    )
  }

}
