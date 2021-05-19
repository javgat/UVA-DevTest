import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EmailUser, TeamService, User, UserService } from '@javgat/devtest-api';
import { Mensaje, Tipo } from 'src/app/shared/app.model';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-users',
  templateUrl: './admin-users.component.html',
  styleUrls: ['./admin-users.component.css']
})
export class AdminUsersComponent extends AdminComponent implements OnInit {

  addUser: boolean
  addUserEmail: string
  updateChild: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService) {
    super(session, router, data, userS, teamS)
    this.addUser = false
    this.addUserEmail = ""
    this.updateChild = false
  }

  ngOnInit(): void {
  }

  doAdminAction(){

  }

  changeNotAddUser(){
    this.addUser = false
  }

  changeAddUser(){
    this.addUser = true
  }

  crearUsuarioSubmit(){
    this.crearUsuario(true)
  }

  crearUsuario(primera: boolean){
    var eu: EmailUser
    eu = {
      email: this.addUserEmail
    }
    this.userS.postEmailUser(eu).subscribe(
      resp =>{
        this.cambiarMensaje(new Mensaje("Usuario creado con éxito", Tipo.SUCCESS, true))
        this.updateUserList()
      },
      err => this.handleErrRelog(err, "Crear usuario por correo", primera, this.crearUsuario, this)
    )
  }

  updateUserList(){
    this.updateChild = !this.updateChild
  }

}
