import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, Tipo } from '../shared/app.model';
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
  likeEmail: string | undefined
  editLikeEmail: string
  mensajeListaVacia: string
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
    this.editLikeUsername = ""
    this.editLikeEmail = ""
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

  saveUsers(resp: User[]){
    this.users = resp
    this.borrarMensaje()
  }

  getUsersFilters(){
    this.cambiarMensaje(new Mensaje("Descargando usuarios... ", Tipo.DOWNLOADING, true))
    this.getUsers(true)
  }

  getUsers(primera: boolean){
    this.userS.getUsers(this.likeUsername, undefined, this.likeEmail).subscribe(
      resp => {
        this.saveUsers(resp)
      },
      err => this.handleErrRelog(err, "obtener la lista de usuarios", primera, this.getUsers, this)
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

  clickSearchEmail(){
    this.likeEmail = this.editLikeEmail
    this.getUsersFilters()
  }

  clickBorrarEmail(){
    this.likeEmail = undefined
    this.editLikeEmail = ""
    this.getUsersFilters()
  }

}
