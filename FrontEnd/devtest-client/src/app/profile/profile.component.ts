import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService, LoginUser, PasswordUpdate, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, SessionUser, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  sessionUser: SessionUser
  profileUser: SessionUser
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
      if (this.id != "" && this.id != undefined) {
        this.getProfileUser(this.id)
      }
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

  getProfileUser(id: string): void {
    this.userService.getUser(id).subscribe(
      resp => {
        this.profileUser = new SessionUser(resp.username, resp.email, resp.fullname, resp.type)
      },
      err => {
        this.data.cambiarMensaje(new Mensaje("No se pudo obtener el usuario: Error " + err.status, Tipo.ERROR, true))
        console.log("No se pudo obtener el usuario")
      }
    )
  }

  changePassSubmit(): void {
    if (this.id == undefined) return
    this.userService.putPassword(this.id, this.pUpdate).subscribe(
      resp => {
        this.data.cambiarMensaje(new Mensaje("Contraseña cambiada con éxito", Tipo.SUCCESS, true))
        if (this.id == this.sessionUser.username) {
          var loginUser: LoginUser = {
            loginid: this.sessionUser.username,
            pass: this.pUpdate.newpass
          }
          this.authService.login(loginUser).subscribe(
            resp => console.log("Sesion JWT recuperada con exito tras cambiar la contraseña"),
            err => console.log("Error al recuperar sesion JWT tras cambiar la contraseña")
          )
        }
      },
      err => {
        this.data.cambiarMensaje(new Mensaje("No se pudo cambiar la contraseña: Error " + err.status, Tipo.ERROR, true))
      }
    )
  }

}
