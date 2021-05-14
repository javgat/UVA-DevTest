import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PublishedTestService, TagService, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListTestsComponent } from '../list-tests.component';

@Component({
  selector: 'app-list-ptests',
  templateUrl: '../list-tests.component.html',
  styleUrls: ['../list-tests.component.css']
})
export class ListPtestsComponent extends ListTestsComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, tS: TestService, tagS: TagService, public pS: PublishedTestService) {
    super(session, router, data, userS, tS, tagS)
    this.arePublished = true
    this.getTestsFilters()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

  getTestsInclude(primera: boolean) {
    //No hay getTestsFromUser en la API que tenga filtros
    this.getTestsEdit(true)
  }

  // EN published tests testEdit es el testPublished
  getTestsEdit(primera: boolean) {
    if (this.pS == undefined) return
    this.pS.getPublicPublishedTests(this.searchTags, this.likeTitle, this.orderBy).subscribe(
      resp => this.tests = resp,
      err => this.handleErrRelog(err, "obtener tests publicados publicos", primera, this.getTestsEdit, this)
    )
  }

}