import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { QuestionService, TagService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListQuestionsComponent } from '../list-questions.component';

@Component({
  selector: 'app-lq-profile-public',
  templateUrl: '../list-questions.component.html',
  styleUrls: ['../list-questions.component.css']
})
export class LqProfilePublicComponent extends ListQuestionsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, qS: QuestionService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, qS, tagS)
    this.id=""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
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
    //No hay getPublicQUestionsOfUser en la API
    if(this.id==undefined) return
    this.userS.getPublicEditQuestionsOfUser(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas no publicadas publicas de un usuario", primera, this.getQuestionsInclude, this)
    )
  }

  getQuestionsEdit(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getPublicEditQuestionsOfUser(this.id, this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas no publicadas publicas de un usuario", primera, this.getQuestionsEdit, this)
    )
  }

}
