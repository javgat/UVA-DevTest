import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TagService, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListTestsComponent } from '../list-tests.component';

@Component({
  selector: 'app-lpt-answered',
  templateUrl: '../list-tests.component.html',
  styleUrls: ['../list-tests.component.css']
})
export class LptAnsweredComponent extends ListTestsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, tS: TestService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, tS, tagS)
    this.id=""
    this.includeNonEdit = true
    this.arePublished = true
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
    this.getTestsEdit(true)
  }

  // EN published tests testEdit es el testPublished
  getTestsEdit(primera: boolean) {
    if(this.id==undefined) return
    this.userS.getAnsweredTestsFromUser(this.id, this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.tests = resp,
      err => this.handleErrRelog(err, "obtener tests disponibles de un usuario", primera, this.getTestsEdit, this)
    )
  }

}
