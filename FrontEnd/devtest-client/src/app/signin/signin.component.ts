import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService, SigninUser } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, SessionLogin, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

// SigninComponent es el componente que permite el registro de un nuevo usuario

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css']
})
export class SigninComponent implements OnInit {

  signinUserEmpty = {
    username: "",
    email: "",
    pass: ""
  }

  private sessionSubscription: Subscription

  reppass: string

  // Variable que se modificara en el formulario de registro
  signinUser = this.signinUserEmpty as SigninUser
  sessionUser: SessionLogin
  constructor(private authService: AuthService, private datos: DataService,
    private session: SessionService, private router: Router) {
    this.sessionUser = new SessionLogin(false)
    this.reppass = ""
    this.session.checkStorageSession()
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionUser = valor
        if (this.sessionUser.isLoggedIn()) {
          this.router.navigate(['/'])
        }
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.sessionSubscription.unsubscribe();
    this.datos.borrarMensaje()
  }

  // Envío de petición de registro a BackEnd, y manejo de la respuesta
  signin(su: SigninUser) {
    this.authService.registerUser(su).subscribe(
      resp => {
        this.datos.cambiarMensaje(new Mensaje("Registro con exito", Tipo.SUCCESS, true))
        console.log("Registro con exito")
        this.router.navigate(['login'])
      },
      err => {
        if (err.status == 409) {
          this.datos.cambiarMensaje(new Mensaje("Ya existe un usuario con ese nombre de usuario o correo electrónico", Tipo.ERROR, true))
        } else {
          let msg: string
          if (err.status >= 500)
            msg = "Error al conectar con el servidor"
          else
            msg = err.error.message
          this.datos.cambiarMensaje(new Mensaje("Error al registrar nuevo usuario: " + msg, Tipo.ERROR, true))
          console.log("Error al registrar nuevo usuario: " + msg)
        }
      }
    )
  }

  // Cuando el formulario se envia, se ejecuta la funcio
  signinSubmit() {
    if(this.reppass == this.signinUser.pass)
      this.signin(this.signinUser)
    else{
      this.datos.cambiarMensaje(new Mensaje("Las contraseñas no coinciden", Tipo.ERROR, true))
    }
  }

}
