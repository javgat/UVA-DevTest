import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
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

  users : User[]

  private userSub : Subscription

  constructor(session: SessionService, userS: UserService,
    router : Router, data: DataService) {
    super(session, router, data, userS)
    this.users = []

    this.userSub = this.session.sessionUser.subscribe(
      valor => {
        if (valor.getRol()==User.RolEnum.Administrador){
          this.getUsers(true)
        }else if(!valor.isEmpty()){
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {

  }

  ngOnDestroy(): void{
    this.userSub.unsubscribe()
    super.onDestroy()
  }

  getUsers(primera: boolean){
    this.userS.getUsers().subscribe(
      resp => {
        this.users = resp
      },
      err =>{
        this.handleErrRelog(err, "Obtener usuarios del panel de administración", primera, this.getUsers, this)
      }
    )
  }

}
