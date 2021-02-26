import { Component, OnInit } from '@angular/core';
import { AuthService, LoginUser } from '@javgat/devtest-api'
import { Mensaje, Session, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

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

  login(lu : LoginUser){
    this.authService.login(lu).subscribe(
      resp => {        
        this.datos.cambiarMensaje(new Mensaje("Inicio sesion con exito", Tipo.SUCCESS, true))
        this.session.cambiarSession(new Session(true, resp.token, lu.loginid))
        console.log("Inicio sesion con exito")
      },
      err =>{
        this.datos.cambiarMensaje(new Mensaje("Error al iniciar sesion", Tipo.ERROR, true))
        console.log("Error al iniciar sesion")
      }
    )
  }

  loginSubmit(){
    this.login(this.loginUser)
  }

}
