import { Component, OnInit } from '@angular/core';
import { AuthService, SigninUser } from '@javgat/devtest-api';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css']
})
export class SigninComponent implements OnInit {

  signinUserEmpty = {
    username:"",
    email:"",
    pass:""
  }

  signinUser = this.signinUserEmpty as SigninUser
  mensaje: Mensaje
  
  constructor(private authService : AuthService, private datos : DataService) { 
    this.mensaje = new Mensaje()
  }

  ngOnInit(): void {
    this.datos.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  signin(su: SigninUser){
    this.authService.registerUser(su).subscribe(
      resp => {        
        this.datos.cambiarMensaje(new Mensaje("Registro con exito", Tipo.SUCCESS, true))
        console.log("Registro con exito")
      },
      err =>{
        this.datos.cambiarMensaje(new Mensaje("Error al registrar nuevo usuario", Tipo.ERROR, true))
        console.log("Error al registrar nuevo usuario")
      }
    )
  }

  signinSubmit(){
    this.signin(this.signinUser)
  }

}
