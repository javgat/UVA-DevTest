import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, LoginUser, PasswordUpdate, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, SessionUser, Tipo, Usuario } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent extends LoggedInController implements OnInit {

  profileUser: Usuario
  private routeSub: Subscription
  id: string | undefined
  pUpdate: PasswordUpdate = {
    oldpass: "",
    newpass: "",
  }

  constructor(session: SessionService, router: Router, private route: ActivatedRoute,
    userS: UserService, data: DataService, private authService: AuthService) {
    super(session, router, data, userS)
    this.profileUser = new SessionUser()
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      this.getProfile(true)
    });
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe();
    this.borrarMensaje()
    super.onDestroy()
  }

  getProfile(primera: boolean){
    if (this.id != "" && this.id != undefined) {
      this.getProfileUser(this.id, primera)
    } else{
      let msg = "No se pudo obtener el id del usuario"
      this.cambiarMensaje(new Mensaje(msg, Tipo.ERROR, true))
      console.log(msg)
    }
  }

  getProfileUser(id: string, primera: boolean): void {
    this.userS.getUser(id).subscribe(
      resp => {
        this.profileUser = new Usuario(resp.username, resp.email, resp.fullname, resp.rol)
      },
      err => {
        this.handleErrRelog(err, "obtener datos de perfil de usuario", primera, this.getProfile, this)
      }
    )
  }

  changePassSubmit(): void {
    this.changePass(true)
  }

  changePass(primera: boolean){
    if (this.id == undefined) return
    this.userS.putPassword(this.id, this.pUpdate).subscribe(
      resp => {
        console.log("Contraseña cambiada")
        if (this.id == this.getSessionUser().username) {
          var loginUser: LoginUser = {
            loginid: this.getSessionUser().username,
            pass: this.pUpdate.newpass
          }
          this.authService.login(loginUser).subscribe(
            resp => {
              console.log("Sesion JWT recuperada con exito tras cambiar la contraseña")
              this.cambiarMensaje(new Mensaje("Contraseña cambiada con éxito", Tipo.SUCCESS, true))
            },
            err => {
              this.handleShowErr(err, "recuperar sesion JWT tras cambiar la contraseña")
              this.logout()
            }
          )
        }
      },
      err => {
        this.handleErrRelog(err, "cambio de contraseña", primera, this.changePass, this)
      }
    )
  }

}
