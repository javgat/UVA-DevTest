import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent implements OnInit {

  private messageSubscription : Subscription
  mensaje: Mensaje
  userid: string
  constructor(private datos: DataService, private router: Router, private userService: UserService) {
    this.mensaje = new Mensaje()
    this.userid = ""
    this.messageSubscription = this.datos.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  ngOnInit(): void {    
  }


  ngOnDestroy(): void {
    this.messageSubscription.unsubscribe();
  }

  forgotSubmit(){
    this.forgot()
  }

  forgot(){
    this.userService.postRecoveryToken(this.userid).subscribe(
      resp => {
        this.datos.cambiarMensaje(new Mensaje("Petición registrada. En breve recibirás un correo electrónico con la información correspondiente", Tipo.SUCCESS, true))
        this.router.navigate(['/login'])
      },
      err => this.datos.handleShowErr(err, "Registrar petición de enviar correo para recuperacion de contraseña")
    )
  }

}