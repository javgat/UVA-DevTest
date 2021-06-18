import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TagService, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListTestsComponent } from '../list-tests.component';

@Component({
  selector: 'app-lt-user',
  templateUrl: '../list-tests.component.html',
  styleUrls: ['../list-tests.component.css']
})
export class LtUserComponent extends ListTestsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, tS: TestService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, tS, tagS)
    this.id=""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['username']
      this.borrarMensaje()
      this.getTestsFilters()
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
    this.routeSub.unsubscribe()
  }

  getTestsInclude(primera: boolean) {
    //No hay getTestsFromUser en la API que tenga filtros
    this.getTestsEdit(true)
  }

  getTestsEdit(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getEditTestsFromUser(this.id, this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.saveTests(resp),
      err => this.handleErrRelog(err, "obtener tests no publicados de un usuario", primera, this.getTestsEdit, this)
    )
  }
}