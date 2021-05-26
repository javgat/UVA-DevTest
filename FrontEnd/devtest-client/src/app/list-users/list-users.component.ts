import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-users',
  templateUrl: './list-users.component.html',
  styleUrls: ['./list-users.component.css']
})
export class ListUsersComponent extends LoggedInController implements OnInit {

  @Input() update: boolean | undefined;

  users: User[]
  likeUsername: string | undefined
  editLikeUsername: string
  mensajeListaVacia: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
    this.editLikeUsername = ""
    this.mensajeListaVacia = "¡Vaya! Parece que no hay ningún usuario para mostrar"
    this.users = []
    this.getUsersFilters()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  ngOnChanges() {
    this.getUsersFilters()
  }   

  getUsersFilters(){
    this.getUsers(true)
  }

  getUsers(primera: boolean){
    this.userS.getUsers(this.likeUsername).subscribe(
      resp => {
        this.users = resp
      },
      err => this.handleErrRelog(this.handleErrRelog, "obtener la lista de usuarios", primera, this.getUsers, this)
    )
  }

  clickSearchUsername(){
    this.likeUsername = this.editLikeUsername
    this.getUsersFilters()
  }

  clickBorrarUsername(){
    this.likeUsername = undefined
    this.editLikeUsername = ""
    this.getUsersFilters()
  }

}
