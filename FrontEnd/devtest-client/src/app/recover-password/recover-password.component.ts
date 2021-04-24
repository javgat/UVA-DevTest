import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PasswordRecovery, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-recover-password',
  templateUrl: './recover-password.component.html',
  styleUrls: ['./recover-password.component.css']
})
export class RecoverPasswordComponent implements OnInit {

  private messageSubscription : Subscription
  mensaje: Mensaje
  routeSub : Subscription
  username : string
  token: string
  pass: string
  passwordRecovery: PasswordRecovery;
  constructor(private data: DataService, private userS: UserService, private route: ActivatedRoute, private router: Router) {
    this.username = ""
    this.token = ""
    this.pass = ""
    this.passwordRecovery = {
      mailtoken: "",
      newpass: ""
    }
    this.mensaje = new Mensaje()
    this.routeSub = this.route.params.subscribe(params => {
      this.username = params['username']
      this.data.borrarMensaje()
    });
    this.route.queryParams.subscribe(params => {
      this.token = params['token']
      this.data.borrarMensaje()
    });
    this.messageSubscription = this.data.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  ngOnInit(): void {
    
  }

  ngOnDestroy(): void{
    this.routeSub.unsubscribe()
  }

  recoverPassSubmit(){
    this.passwordRecovery.mailtoken = this.token
    this.passwordRecovery.newpass = this.pass
    this.recoverPass()
  }

  recoverPass(){
    this.userS.recoverPassword(this.username, this.passwordRecovery).subscribe(
      resp => {
        this.data.cambiarMensaje(new Mensaje("Contraseña actualizada con exito", Tipo.SUCCESS, true))
        this.router.navigate(['/login'])
      },
      err => this.data.handleShowErr(err, "Actualizar contraseña perdida")
    )
  }

}
