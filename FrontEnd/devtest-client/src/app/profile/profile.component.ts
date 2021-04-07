import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService, LoginUser, PasswordUpdate, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, SessionUser, Tipo, Usuario } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  sessionUser: SessionUser
  profileUser: Usuario
  mensaje: Mensaje
  private sessionUserSubscription: Subscription
  private routeSub: Subscription
  private messageSubscription: Subscription
  id: string | undefined
  pUpdate: PasswordUpdate = {
    oldpass: "",
    newpass: "",
  }

  constructor(private session: SessionService, private route: ActivatedRoute,
    private userService: UserService, private data: DataService, private authService: AuthService) {
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor => this.sessionUser = valor
    )
    this.profileUser = new SessionUser()
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.data.borrarMensaje()
      this.getProfile(true)
    });
    this.mensaje = new Mensaje()
    this.messageSubscription = this.data.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.sessionUserSubscription.unsubscribe();
    this.routeSub.unsubscribe();
    this.messageSubscription.unsubscribe();

    this.data.borrarMensaje()
  }

  getProfile(primera: boolean){
    if (this.id != "" && this.id != undefined) {
      this.getProfileUser(this.id, primera)
    } else{
      let msg = "No se pudo obtener el id del usuario"
      this.data.cambiarMensaje(new Mensaje(msg, Tipo.ERROR, true))
      console.log(msg)
    }
  }

  getProfileUser(id: string, primera: boolean): void {
    this.userService.getUser(id).subscribe(
      resp => {
        this.profileUser = new Usuario(resp.username, resp.email, resp.fullname, resp.rol)
      },
      err => {
        this.session.handleErrRelog(err, "obtener datos de perfil de usuario", primera, this.getProfile, this)
      }
    )
  }

  changePassSubmit(): void {
    this.changePass(true)
  }

  changePass(primera: boolean){
    if (this.id == undefined) return
    this.userService.putPassword(this.id, this.pUpdate).subscribe(
      resp => {
        console.log("Contraseña cambiada")
        if (this.id == this.sessionUser.username) {
          var loginUser: LoginUser = {
            loginid: this.sessionUser.username,
            pass: this.pUpdate.newpass
          }
          this.authService.login(loginUser).subscribe(
            resp => {
              console.log("Sesion JWT recuperada con exito tras cambiar la contraseña")
              this.data.cambiarMensaje(new Mensaje("Contraseña cambiada con éxito", Tipo.SUCCESS, true))
            },
            err => {
              this.data.handleShowErr(err, "recuperar sesion JWT tras cambiar la contraseña")
              this.session.logout()
            }
          )
        }
      },
      err => {
        this.session.handleErrRelog(err, "cambio de contraseña", primera, this.changePass, this)
      }
    )
  }

}
