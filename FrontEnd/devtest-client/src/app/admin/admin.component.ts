import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TeamService, User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent extends LoggedInController implements OnInit {


  private userSub: Subscription

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected teamS: TeamService) {
    super(session, router, data, userS)

    this.userSub = this.session.sessionUser.subscribe(
      valor => {
        if (valor.getRol() == User.RolEnum.Administrador) {
          this.doAdminAction()
        } else if (!valor.isEmpty()) {
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {

  }

  ngOnDestroy(): void {
    this.userSub.unsubscribe()
    super.onDestroy()
  }

  doAdminAction(){}



}
