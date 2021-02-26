import { Component, OnInit } from '@angular/core';
import { AuthService, LoginUser } from '@javgat/devtest-api'
import { Mensaje, Session, Tipo } from '../shared/app.model';
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

  constructor(private authService : AuthService, private datos: DataService,
    private session: SessionService) {
    this.mensaje = new Mensaje()
  }

  ngOnInit(): void {
    this.datos.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  // Envío de petición de login a BackEnd, y manejo de la respuesta
  login(lu : LoginUser){
    this.authService.login(lu).subscribe(
      resp => {        
        this.datos.cambiarMensaje(new Mensaje("Inicio sesion con exito", Tipo.SUCCESS, true))
        this.session.cambiarSession(new Session(true, resp.token, lu.loginid))
        console.log("Inicio sesion con exito")
        // Redireccion a pagina principal?
      },
      err =>{
        let msg: string
        if(err.status >= 500)
          msg = "Error al conectar con el servidor"
        else
          msg = err.error.message
        this.datos.cambiarMensaje(new Mensaje("Error al iniciar sesion: "+msg, Tipo.ERROR, true))
        console.log("Error al iniciar sesion: "+msg)
      }
    )
  }

  // Cuando el formulario se envia, se ejecuta la funcio
  loginSubmit(){
    this.login(this.loginUser)
  }

}
