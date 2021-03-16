import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService, LoginUser } from '@javgat/devtest-api'
import { Subscription } from 'rxjs';
import { Mensaje, SessionLogin, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

// LoginComponent es el componente que permite el inicio de sesión y autenticación de un usuario

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginUserEmpty = {
    loginid: "",
    pass: ""
  }
  // Variable que se modificara en el formulario de inicio de sesión
  loginUser = this.loginUserEmpty as LoginUser
  mensaje: Mensaje
  sessionLogin : SessionLogin
  private sessionSubscription : Subscription
  private messageSubscription : Subscription

  constructor(private authService : AuthService, private datos: DataService,
    private session: SessionService, private router: Router) {
    this.mensaje = new Mensaje()
    this.sessionLogin = new SessionLogin(false)
    this.session.checkStorageSession()
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionLogin = valor
        if(this.sessionLogin.logged){
          this.router.navigate(['/'])
        }
      }
    )
    this.messageSubscription = this.datos.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  ngOnInit(): void {
    
  }


  ngOnDestroy(): void {
    this.sessionSubscription.unsubscribe();
    this.messageSubscription.unsubscribe();
    this.datos.borrarMensaje()
  }

  // Envío de petición de login a BackEnd, y manejo de la respuesta
  login(lu : LoginUser){
    this.authService.login(lu).subscribe(
      resp => {        
        this.datos.cambiarMensaje(new Mensaje("Inicio sesion con exito", Tipo.SUCCESS, true))
        this.session.cambiarSession(new SessionLogin(true, lu.loginid))
        console.log("Inicio sesion con exito")
      },
      err =>{
        let msg: string
        if(err.status >= 500)
          msg = "Error al conectar con el servidor"
        else if(err.error != undefined)
          msg = err.error.message
        else
          msg = "undefined"
        this.datos.cambiarMensaje(new Mensaje("Error al iniciar sesion: "+msg, Tipo.ERROR, true))
        console.log("Error al iniciar sesion: "+msg+ err.status)
      }
    )
  }

  // Cuando el formulario se envia, se ejecuta la funcio
  loginSubmit(){
    this.login(this.loginUser)
  }

}
