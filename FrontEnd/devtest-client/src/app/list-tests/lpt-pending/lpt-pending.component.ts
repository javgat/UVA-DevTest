import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TagService, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ListTestsComponent } from '../list-tests.component';

@Component({
  selector: 'app-lpt-pending',
  templateUrl: '../list-tests.component.html',
  styleUrls: ['../list-tests.component.css']
})
export class LptPendingComponent extends ListTestsComponent implements OnInit {

  id: string | undefined
  routeSub: Subscription
  getIdFromSession: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, tS: TestService, tagS: TagService, private route: ActivatedRoute) {
    super(session, router, data, userS, tS, tagS)
    this.getIdFromSession=false
    this.id=""
    this.includeNonEdit = true
    this.arePublished = true
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['username']
      if(this.id==undefined){
        this.getIdFromSession=true
        this.id=this.getSessionUser().getUsername()
      }
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

  doHasUserAction(){
    if(this.getIdFromSession){
      this.id=this.getSessionUser().getUsername()
    }
    this.getTestsFilters()
    super.doHasUserAction()
  }

  getTestsInclude(primera: boolean) {
    this.getTestsEdit(true)
  }

  // EN published tests testEdit es el testPublished
  getTestsEdit(primera: boolean) {
    if(this.id==undefined || this.id==""){
      this.borrarMensaje()
      return
    } 
    this.userS.getPendingTestsFromUser(this.id, this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.saveTests(resp),
      err => this.handleErrRelog(err, "obtener tests pendientes de un usuario", primera, this.getTestsEdit, this)
    )
  }
}
