import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TeamService, User, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-users',
  templateUrl: './admin-users.component.html',
  styleUrls: ['./admin-users.component.css']
})
export class AdminUsersComponent extends AdminComponent implements OnInit {

  users: User[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService) {
    super(session, router, data, userS, teamS)
    this.users = []
  }

  ngOnInit(): void {
  }
  
  getUsers(primera: boolean) {
    this.userS.getUsers().subscribe(
      resp => {
        this.users = resp
      },
      err => {
        this.handleErrRelog(err, "Obtener usuarios del panel de administración", primera, this.getUsers, this)
      }
    )
  }

  doAdminAction(){
    this.getUsers(true)
  }

}
